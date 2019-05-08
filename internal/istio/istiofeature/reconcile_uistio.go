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
	"github.com/goph/emperror"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/banzaicloud/pipeline/cluster"
	pConfig "github.com/banzaicloud/pipeline/config"
	pkgHelm "github.com/banzaicloud/pipeline/pkg/helm"
)

const uistioNamespace = "istio-system"
const uistioDeploymentName = pkgHelm.BanzaiRepository + "/" + "uistio"
const uistioReleaseName = "uistio"

func (m *MeshReconciler) ReconcileUistio(desiredState DesiredState) error {
	m.logger.Debug("reconciling Uistio")
	defer m.logger.Debug("Uistio reconciled")

	if desiredState == DesiredStatePresent {
		err := m.installUistio(m.Master, m.logger)
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
