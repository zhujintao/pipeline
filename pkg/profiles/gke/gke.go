package gke

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgGKE "github.com/banzaicloud/pipeline/pkg/cluster/gke"
	pkgDefaultsGKE "github.com/banzaicloud/pipeline/pkg/profiles/defaults/gke"
)

type Profile struct {
	defaultNodePoolName string
	*pkgDefaultsGKE.Defaults
}

func NewProfile(defaultNodePoolName string, gke *pkgDefaultsGKE.Defaults) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		Defaults:            gke,
	}
}

func (p *Profile) GetDefaultProfile() *pkgCluster.CreateClusterRequest {

	nodepools := make(map[string]*pkgGKE.NodePool)
	nodepools[p.defaultNodePoolName] = &pkgGKE.NodePool{
		Autoscaling:      p.NodePools.Autoscaling,
		MinCount:         p.NodePools.MinCount,
		MaxCount:         p.NodePools.MaxCount,
		Count:            p.NodePools.Count,
		NodeInstanceType: p.NodePools.InstanceType,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: p.Location,
		Cloud:    pkgCluster.Google,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterGKE: &pkgGKE.CreateClusterGKE{
				NodeVersion: p.NodeVersion,
				NodePools:   nodepools,
				Master: &pkgGKE.Master{
					Version: p.MasterVersion,
				},
			},
		},
	}

}
