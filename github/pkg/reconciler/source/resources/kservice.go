/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resources

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/system"
	v1 "knative.dev/serving/pkg/apis/serving/v1"
)

type ServiceArgs struct {
	ServiceAccountName  string
	ReceiveAdapterImage string
	AdapterName         string
}

var (
	labels = map[string]string{
		"sources.knative.dev/role":    "adapter",
		"eventing.knative.dev/source": controllerAgentName,
	}
)

// MakeService generates, but does not create, a global Github Service
func MakeService(args *ServiceArgs) *v1.Service {
	return &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      args.AdapterName,
			Namespace: system.Namespace(),
			Labels:    labels,
		},
		Spec: v1.ServiceSpec{
			ConfigurationSpec: v1.ConfigurationSpec{
				Template: v1.RevisionTemplateSpec{
					Spec: v1.RevisionSpec{
						PodSpec: corev1.PodSpec{
							ServiceAccountName: args.ServiceAccountName,
							Containers: []corev1.Container{{
								Image: args.ReceiveAdapterImage,
								Env: []corev1.EnvVar{{
									Name:  system.NamespaceEnvKey,
									Value: system.Namespace(),
								}, {
									Name:  "METRICS_DOMAIN",
									Value: "knative.dev/sources",
								}, {
									Name:  "CONFIG_OBSERVABILITY_NAME",
									Value: "config-observability",
								}, {
									Name:  "CONFIG_LOGGING_NAME",
									Value: "config-logging",
								}, {
									Name:  "SYSTEM_NAMESPACE",
									Value: system.Namespace(),
								}},
							}},
						},
					},
				},
			},
		},
	}
}