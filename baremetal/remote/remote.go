/*
Copyright 2018 The Kubernetes Authors.
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

package remote

import (
	"github.com/pkg/errors"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha2"
	kcfg "sigs.k8s.io/cluster-api/util/kubeconfig"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewClusterClient creates a new ClusterClient.
func NewClusterClient(c client.Client, cluster *clusterv1.Cluster) (corev1.CoreV1Interface, error) {
	kubeconfig, err := kcfg.FromSecret(c, cluster)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve kubeconfig secret for Cluster %q in namespace %q",
			cluster.Name, cluster.Namespace)
	}

	restConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeconfig)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create client configuration for Cluster %q in namespace %q",
			cluster.Name, cluster.Namespace)
	}

	return corev1.NewForConfig(restConfig)
}
