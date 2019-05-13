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
	"github.com/gofrs/uuid"
	"github.com/goph/emperror"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/banzaicloud/pipeline/internal/clustergroup/api"
)

type ServiceMeshFeatureHandler struct {
	logger       logrus.FieldLogger
	errorHandler emperror.Handler
}

const FeatureName = "servicemesh"

// NewServiceMeshFeatureHandler returns a new ServiceMeshFeatureHandler instance.
func NewServiceMeshFeatureHandler(
	logger logrus.FieldLogger,
	errorHandler emperror.Handler,
) *ServiceMeshFeatureHandler {
	return &ServiceMeshFeatureHandler{
		logger:       logger,
		errorHandler: errorHandler,
	}
}

func (h *ServiceMeshFeatureHandler) ReconcileState(featureState api.Feature) error {
	cid, err := uuid.NewV4()
	if err != nil {
		return emperror.Wrap(err, "could not generate uuid")
	}
	logger := h.logger.WithFields(logrus.Fields{
		"correlationID":    cid,
		"clusterGroupID":   featureState.ClusterGroup.Id,
		"clusterGroupName": featureState.ClusterGroup.Name,
	})

	logger.Info("reconciling service mesh feature")
	defer logger.Info("service mesh feature reconciled")

	config, err := h.GetConfigFromState(featureState)
	if err != nil {
		return errors.WithStack(err)
	}

	mesh := NewMeshReconciler(*config, logger, h.errorHandler)
	err = mesh.Reconcile()
	if err != nil {
		h.errorHandler.Handle(err)
		return emperror.Wrap(err, "could not reconcile service mesh")
	}

	return nil
}

func (h *ServiceMeshFeatureHandler) GetConfigFromState(state api.Feature) (*Config, error) {
	var config Config
	err := mapstructure.Decode(state.Properties, &config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	config.name = state.ClusterGroup.Name
	config.enabled = state.Enabled
	config.clusterGroup = state.ClusterGroup

	return &config, nil
}

func (h *ServiceMeshFeatureHandler) ValidateState(featureState api.Feature) error {
	return nil
}

func (h *ServiceMeshFeatureHandler) ValidateProperties(properties interface{}) error {
	var config Config
	err := mapstructure.Decode(properties, &config)
	if err != nil {
		return errors.WithStack(err)
	}

	if config.MasterClusterID == 0 {
		return errors.New("master cluster ID is required")
	}

	return nil
}

func (f *ServiceMeshFeatureHandler) GetMembersStatus(featureState api.Feature) (map[uint]string, error) {
	statusMap := make(map[uint]string, 0)
	for _, memberCluster := range featureState.ClusterGroup.Clusters {
		statusMap[memberCluster.GetID()] = "ready"
	}
	return statusMap, nil
}
