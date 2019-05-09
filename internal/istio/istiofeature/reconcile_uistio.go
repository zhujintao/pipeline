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
	"time"

	"github.com/goph/emperror"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/banzaicloud/pipeline/cluster"
	pConfig "github.com/banzaicloud/pipeline/config"
	"github.com/banzaicloud/pipeline/internal/backoff"
	pkgHelm "github.com/banzaicloud/pipeline/pkg/helm"
)

const uistioNamespace = "istio-system"
const uistioDeploymentName = pkgHelm.BanzaiRepository + "/" + "uistio"
const uistioReleaseName = "uistio"

func (m *MeshReconciler) ReconcileUistio(desiredState DesiredState) error {
	m.logger.Debug("reconciling Uistio")
	defer m.logger.Debug("Uistio reconciled")

	if desiredState == DesiredStatePresent {
		c, _ := m.GetApiExtensionK8sClient(m.Master)
		err := m.waitForMetricCRD("metrics.config.istio.io", c)
		if err != nil {
			return emperror.Wrap(err, "could not install Uistio")
		}

		err = m.installUistio(m.Master, m.logger)
		if err != nil {
			return emperror.Wrap(err, "could not install Uistio")
		}
	} else {
		err := m.uninstallUistio(m.Master, m.logger)
		if err != nil {
			return emperror.Wrap(err, "could not remove Uistio")
		}
	}

	return nil
}

// waitForMetricCRD
func (m *MeshReconciler) waitForMetricCRD(name string, client *apiextensionsclient.Clientset) error {
	m.logger.WithField("name", name).Debug("waiting for metric CRD")

	var backoffConfig = backoff.ConstantBackoffConfig{
		Delay:      time.Duration(backoffDelaySeconds) * time.Second,
		MaxRetries: backoffMaxretries,
	}
	var backoffPolicy = backoff.NewConstantBackoffPolicy(&backoffConfig)

	err := backoff.Retry(func() error {
		c, err := m.GetApiExtensionK8sClient(m.Master)
		if err != nil {
			return err
		}

		_, err = c.ApiextensionsV1beta1().CustomResourceDefinitions().Get(name, metav1.GetOptions{})

		return err
	}, backoffPolicy)

	return err
}

// uninstallIstioOperator removes istio-operator from a cluster
func (m *MeshReconciler) uninstallUistio(c cluster.CommonCluster, logger logrus.FieldLogger) error {
	logger.Debug("removing Uistio")

	err := deleteDeployment(c, uistioReleaseName)
	if err != nil {
		return emperror.Wrap(err, "could not remove Uistio")
	}

	return nil
}

// installIstioOperator installs istio-operator on a cluster
func (m *MeshReconciler) installUistio(c cluster.CommonCluster, logger logrus.FieldLogger) error {
	logger.Debug("installing Uistio")

	err := installDeployment(
		c,
		uistioNamespace,
		uistioDeploymentName,
		uistioReleaseName,
		[]byte{},
		viper.GetString(pConfig.UistioChartVersion),
		true,
		m.logger,
	)
	if err != nil {
		return emperror.Wrap(err, "could not install Uistio")
	}

	return nil
}
