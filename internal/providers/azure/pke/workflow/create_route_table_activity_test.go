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

package workflow

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-01-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/stretchr/testify/assert"
)

func TestGetCreateOrUpdateRouteTableParams(t *testing.T) {
	t.Run("typical input", func(t *testing.T) {
		input := CreateRouteTableActivityInput{
			OrganizationID:    1,
			SecretID:          "0123456789abcdefghijklmnopqrstuvwxyz",
			ClusterName:       "test-cluster",
			ResourceGroupName: "test-rg",
			RouteTable: RouteTable{
				Location: "test-location",
				Name:     "test-route-table",
			},
		}
		expected := network.RouteTable{
			Location: to.StringPtr("test-location"),
			Tags: map[string]*string{
				"kubernetesCluster-test-cluster": to.StringPtr("owned"),
			},
		}
		result := input.getCreateOrUpdateRouteTableParams()
		assert.Equal(t, expected, result)
	})
}