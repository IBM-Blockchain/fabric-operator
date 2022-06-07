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
	current "github.com/IBM-Blockchain/fabric-operator/api/v1beta1"
	"github.com/IBM-Blockchain/fabric-operator/pkg/manager/resources"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (o *Override) ServiceAccount(object v1.Object, sa *corev1.ServiceAccount, action resources.Action) error {
	instance := object.(*current.IBPConsole)
	switch action {
	case resources.Create:
		return o.CreateServiceAccount(instance, sa)
	case resources.Update:
		return o.UpdateServiceAccount(instance, sa)
	}

	return nil
}

func (o *Override) CreateServiceAccount(instance *current.IBPConsole, sa *corev1.ServiceAccount) error {
	return o.commonServiceAccount(instance, sa)
}

func (o *Override) UpdateServiceAccount(instance *current.IBPConsole, sa *corev1.ServiceAccount) error {
	return o.commonServiceAccount(instance, sa)
}

func (o *Override) commonServiceAccount(instance *current.IBPConsole, sa *corev1.ServiceAccount) error {
	for _, pullSecret := range instance.Spec.ImagePullSecrets {
		imagePullSecret := corev1.LocalObjectReference{
			Name: pullSecret,
		}

		sa.ImagePullSecrets = append(sa.ImagePullSecrets, imagePullSecret)
	}

	return nil
}
