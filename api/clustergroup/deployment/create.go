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

package deployment

import (
	"context"
	"net/http"

	pkgDep "github.com/banzaicloud/pipeline/internal/clustergroup/deployment"
	"github.com/gin-gonic/gin"

	"github.com/banzaicloud/pipeline/auth"
	ginutils "github.com/banzaicloud/pipeline/internal/platform/gin/utils"
	pkgCommon "github.com/banzaicloud/pipeline/pkg/common"
)

func (n *API) Create(c *gin.Context) {
	ctx := ginutils.Context(context.Background(), c)

	clusterGroupID, ok := ginutils.UintParam(c, "id")
	if !ok {
		return
	}

	orgID := auth.GetCurrentOrganization(c.Request).ID
	clusterGroup, err := n.clusterGroupManager.GetClusterGroupByID(ctx, clusterGroupID, orgID)
	if err != nil {
		n.errorHandler.Handle(c, err)
		return
	}

	organization, err := auth.GetOrganizationById(clusterGroup.OrganizationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkgCommon.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Error  getting organization",
			Error:   err.Error(),
		})
		return
	}
	var deployment *pkgDep.ClusterGroupDeployment
	if err := c.ShouldBindJSON(&deployment); err != nil {
		n.errorHandler.Handle(c, c.Error(err).SetType(gin.ErrorTypeBind))
		return
	}

	// TODO currently we don't check release name uniqueness, but upgrade any release found with the same name
	// 	may be we can make this optional
	if len(deployment.ReleaseName) == 0 {
		deployment.ReleaseName = n.deploymentManager.GenerateReleaseName(clusterGroup)
	}

	targetClusterStatus, err := n.deploymentManager.CreateDeployment(clusterGroup, organization.Name, deployment)
	if err != nil {
		n.errorHandler.Handle(c, err)
		return
	}

	n.logger.Debug("Release name: ", deployment.ReleaseName)
	response := pkgDep.CreateUpdateDeploymentResponse{
		ReleaseName:    deployment.ReleaseName,
		TargetClusters: targetClusterStatus,
	}
	c.JSON(http.StatusCreated, response)
	return
}
