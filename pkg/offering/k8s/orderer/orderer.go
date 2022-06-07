/*
 * Copyright contributors to the Hyperledger Fabric Operator project
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 * 	  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package k8sorderer

import (
	"fmt"

	current "github.com/IBM-Blockchain/fabric-operator/api/v1beta1"
	config "github.com/IBM-Blockchain/fabric-operator/operatorconfig"
	k8sclient "github.com/IBM-Blockchain/fabric-operator/pkg/k8s/controllerclient"
	baseorderer "github.com/IBM-Blockchain/fabric-operator/pkg/offering/base/orderer"
	baseordereroverride "github.com/IBM-Blockchain/fabric-operator/pkg/offering/base/orderer/override"
	"github.com/IBM-Blockchain/fabric-operator/pkg/offering/common"
	"github.com/IBM-Blockchain/fabric-operator/pkg/offering/k8s/orderer/override"
	"github.com/IBM-Blockchain/fabric-operator/version"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var log = logf.Log.WithName("k8s_orderer")

const (
	defaultOrdererNode = "./definitions/orderer/orderernode.yaml"
)

type Orderer struct {
	*baseorderer.Orderer
}

func New(client k8sclient.Client, scheme *runtime.Scheme, config *config.Config) *Orderer {
	o := &override.Override{
		Override: baseordereroverride.Override{
			Client: client,
			Config: config,
		},
	}

	orderer := &Orderer{
		Orderer: baseorderer.New(client, scheme, config, o),
	}

	return orderer
}

func (o *Orderer) Reconcile(instance *current.IBPOrderer, update baseorderer.Update) (common.Result, error) {
	var err error

	if instance.Spec.NodeNumber == nil {
		log.Info(fmt.Sprintf("Reconciling cluster instance '%s' ... update: %+v", instance.Name, update))

		versionSet, err := o.SetVersion(instance)
		if err != nil {
			return common.Result{}, errors.Wrap(err, fmt.Sprintf("failed updating CR '%s' to version '%s'", instance.Name, version.Operator))
		}
		if versionSet {
			log.Info("Instance version updated, requeuing request...")
			return common.Result{
				Result: reconcile.Result{
					Requeue: true,
				},
			}, nil
		}

		if instance.Status.Status == "" || instance.Status.Status == current.False || (instance.Status.Version != "" && version.String(instance.Status.Version).GreaterThan(version.V210)) {
			instanceUpdated, err := o.PreReconcileChecks(instance, update)
			if err != nil {
				return common.Result{}, errors.Wrap(err, "failed pre reconcile checks")
			}

			if instanceUpdated {
				log.Info("Instance updated, requeuing request...")
				return common.Result{
					Result: reconcile.Result{
						Requeue: true,
					},
					OverrideUpdateStatus: true,
				}, nil
			}
		}
	}

	// TODO: Major rehaul is needed of versioning and migration strategy. Need a way to
	// migrate as first step to get CR spec in appropriate state to avoid versioning checks
	// like below and above
	if (instance.Status.Version == "" && instance.Status.Status == current.True) || (instance.Status.Version != "" && version.String(instance.Status.Version).Equal(version.V210)) {
		if instance.Spec.NodeNumber == nil {
			number := 1
			instance.Spec.NodeNumber = &number
		}
	}

	var result common.Result
	if instance.Spec.NodeNumber == nil {
		result, err := o.ReconcileCluster(instance, update, o.AddHostPortToProfile)
		if err != nil {
			return common.Result{}, errors.Wrap(err, "failed to reconcile cluster")
		}

		return result, nil
	}

	result, err = o.ReconcileNode(instance, update)
	if err != nil {
		return common.Result{}, errors.Wrap(err, "failed to reconcile node")
	}

	return result, nil
}

func (o *Orderer) ReconcileNode(instance *current.IBPOrderer, update baseorderer.Update) (common.Result, error) {
	var err error

	hostAPI := fmt.Sprintf("%s-%s-orderer.%s", instance.Namespace, instance.Name, instance.Spec.Domain)
	hostOperations := fmt.Sprintf("%s-%s-operations.%s", instance.Namespace, instance.Name, instance.Spec.Domain)
	hostGrpc := fmt.Sprintf("%s-%s-grpcweb.%s", instance.Namespace, instance.Name, instance.Spec.Domain)
	hosts := []string{}
	currentVer := version.String(instance.Spec.FabricVersion)
	if currentVer.EqualWithoutTag(version.V2_4_1) || currentVer.GreaterThan(version.V2_4_1) {
		hostAdmin := fmt.Sprintf("%s-%s-admin.%s", instance.Namespace, instance.Name, instance.Spec.Domain)
		hosts = append(hosts, hostAPI, hostOperations, hostGrpc, hostAdmin, "127.0.0.1")
		//TODO: need to Re-enroll when orderer migrated from 1.4.x/2.2.x to 2.4.1
	} else {
		hosts = append(hosts, hostAPI, hostOperations, hostGrpc, "127.0.0.1")
	}

	o.CheckCSRHosts(instance, hosts)

	k8snode := NewNode(baseorderer.NewNode(o.Client, o.Scheme, o.Config, instance.GetName(), o.RenewCertTimers, o.RestartManager))

	log.Info(fmt.Sprintf("Reconciling Orderer node %s", instance.GetName()))
	if !instance.Spec.IsUsingChannelLess() && instance.Spec.GenesisBlock == "" && !(instance.Spec.IsPrecreateOrderer()) {
		return common.Result{}, fmt.Errorf("Genesis block not provided for orderer node: %s", instance.GetName())
	}

	result, err := k8snode.Reconcile(instance, update)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (o *Orderer) GetNodes(instance *current.IBPOrderer) []*Node {
	size := instance.Spec.ClusterSize
	nodes := []*Node{}
	for i := 1; i <= size; i++ {
		node := o.GetNode(i)
		nodes = append(nodes, node)
	}
	return nodes
}

func (o *Orderer) GetNode(nodeNumber int) *Node {
	basenode := o.NodeManager.GetNode(nodeNumber, o.RenewCertTimers, o.RestartManager)
	return NewNode(basenode)
}
