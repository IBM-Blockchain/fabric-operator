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

package override

import (
	"fmt"

	current "github.com/IBM-Blockchain/fabric-operator/api/v1beta1"
	"github.com/IBM-Blockchain/fabric-operator/pkg/manager/resources"
	"github.com/IBM-Blockchain/fabric-operator/pkg/manager/resources/service"
	routev1 "github.com/openshift/api/route/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (o *Override) ProxyRoute(object v1.Object, route *routev1.Route, action resources.Action) error {
	instance := object.(*current.IBPConsole)
	switch action {
	case resources.Create:
		return o.CreateProxyRoute(instance, route)
	case resources.Update:
		return o.UpdateProxyRoute(instance, route)
	}

	return nil
}

func (o *Override) CreateProxyRoute(instance *current.IBPConsole, route *routev1.Route) error {
	route.Name = fmt.Sprintf("%s-proxy", instance.GetName())
	route.Spec.Host = instance.Namespace + "-" + instance.GetName() + "-proxy" + "." + instance.Spec.NetworkInfo.Domain
	weight := int32(100)
	route.Spec.To = routev1.RouteTargetReference{
		Kind:   "Service",
		Name:   service.GetName(instance.Name),
		Weight: &weight,
	}

	route.Spec.Port = &routev1.RoutePort{
		TargetPort: intstr.FromString("optools"),
	}

	route.Spec.TLS = &routev1.TLSConfig{
		Termination: routev1.TLSTerminationPassthrough,
	}

	return nil
}

func (o *Override) UpdateProxyRoute(instance *current.IBPConsole, route *routev1.Route) error {
	return nil
}
