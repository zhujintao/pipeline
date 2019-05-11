// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package istiofeature

import (
	"strconv"
	"time"

	"github.com/goph/emperror"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/banzaicloud/istio-operator/pkg/apis/istio/v1beta1"
	istiooperatorclientset "github.com/banzaicloud/istio-operator/pkg/client/clientset/versioned"
	pConfig "github.com/banzaicloud/pipeline/config"
	"github.com/banzaicloud/pipeline/internal/backoff"
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
)

func (m *MeshReconciler) ReconcileIstio(desiredState DesiredState) error {
	m.logger.Debug("reconciling Istio CR")
	defer m.logger.Debug("Istio CR reconciled")

	client, err := m.GetMasterIstioOperatorK8sClient()
	if err != nil {
		return err
	}

	if desiredState == DesiredStatePresent {
		ipRanges, err := m.Master.GetK8sIpv4Cidrs()
		if err != nil {
			return emperror.Wrap(err, "could not get ipv4 ranges for cluster")
		}

		istio, err := client.IstioV1beta1().Istios(istioOperatorNamespace).Get(m.Configuration.name, metav1.GetOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			return emperror.Wrap(err, "could not check existance Istio CR")
		}

		if k8serrors.IsNotFound(err) {
			istio = &v1beta1.Istio{
				ObjectMeta: metav1.ObjectMeta{
					Name: m.Configuration.name,
				},
			}
		}

		istio = m.configureIstioCR(istio, m.Configuration, ipRanges)

		if k8serrors.IsNotFound(err) {
			_, err = client.IstioV1beta1().Istios(istioOperatorNamespace).Create(istio)
			if err != nil {
				return emperror.Wrap(err, "could not create Istio CR")
			}
		} else if err == nil {
			_, err := client.IstioV1beta1().Istios(istioOperatorNamespace).Update(istio)
			if err != nil {
				return emperror.Wrap(err, "could not update Istio CR")
			}
		}
	} else {
		err := client.IstioV1beta1().Istios(istioOperatorNamespace).Delete(m.Configuration.name, &metav1.DeleteOptions{})
		if err != nil && !k8serrors.IsNotFound(err) {
			return emperror.Wrap(err, "could not remove Istio CR")
		}

		err = m.waitForIstioCRToBeDeleted(client)
		if err != nil {
			return emperror.Wrap(err, "timeout during waiting for Istio CR to be deleted")
		}
	}

	return nil
}

// waitForIstioCRToBeDeleted wait for Istio CR to be deleted
func (m *MeshReconciler) waitForIstioCRToBeDeleted(client *istiooperatorclientset.Clientset) error {
	m.logger.Debug("waiting for Istio CR to be deleted")

	var backoffConfig = backoff.ConstantBackoffConfig{
		Delay:      time.Duration(backoffDelaySeconds) * time.Second,
		MaxRetries: backoffMaxretries,
	}
	var backoffPolicy = backoff.NewConstantBackoffPolicy(&backoffConfig)

	err := backoff.Retry(func() error {
		_, err := client.IstioV1beta1().Istios(istioOperatorNamespace).Get(m.Configuration.name, metav1.GetOptions{})
		if k8serrors.IsNotFound(err) {
			return nil
		}

		return errors.New("Istio CR still exists")
	}, backoffPolicy)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// configureIstioCR configures istio-operator specific CR based on the given params
func (m *MeshReconciler) configureIstioCR(istio *v1beta1.Istio, config Config, ipRanges *pkgCluster.Ipv4Cidrs) *v1beta1.Istio {
	labels := istio.GetLabels()
	if labels == nil {
		labels = make(map[string]string, 0)
	}
	labels["controller-tools.k8s.io"] = "1.0"
	labels["cluster.banzaicloud.com/id"] = strconv.FormatUint(uint64(m.Master.GetID()), 10)
	labels["cluster.banzaicloud.com/cloud"] = m.Master.GetCloud()
	labels["cluster.banzaicloud.com/distribution"] = m.Master.GetDistribution()
	istio.SetLabels(labels)
	istio.Spec.MTLS = config.EnableMTLS
	istio.Spec.AutoInjectionNamespaces = config.AutoSidecarInjectNamespaces
	istio.Spec.Version = istioVersion
	istio.Spec.Gateways.IngressConfig.MaxReplicas = 1
	istio.Spec.Gateways.EgressConfig.MaxReplicas = 1
	istio.Spec.Pilot = v1beta1.PilotConfiguration{
		Image:       viper.GetString(pConfig.IstioPilotImage),
		MaxReplicas: 1,
	}
	istio.Spec.Mixer = v1beta1.MixerConfiguration{
		Image:       viper.GetString(pConfig.IstioMixerImage),
		MaxReplicas: 1,
	}

	if len(m.Remotes) > 0 {
		enabled := true
		istio.Spec.UseMCP = enabled
		istio.Spec.MTLS = enabled
		istio.Spec.MeshExpansion = &enabled
		istio.Spec.ControlPlaneSecurityEnabled = enabled
	}

	// if config.BypassEgressTraffic {
	// 	istio.Spec.IncludeIPRanges = strings.Join(ipRanges.PodIPRanges, ",") + "," + strings.Join(ipRanges.ServiceClusterIPRanges, ",")
	// }

	return istio
}
