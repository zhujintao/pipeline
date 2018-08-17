package aks

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgAKS "github.com/banzaicloud/pipeline/pkg/cluster/aks"
	pkgDefaultsAKS "github.com/banzaicloud/pipeline/pkg/profiles/defaults/aks"
)

type Profile struct {
	defaultNodePoolName string
	*pkgDefaultsAKS.Defaults
}

func NewProfile(defaultNodePoolName string, aks *pkgDefaultsAKS.Defaults) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		Defaults:            aks,
	}
}

func (p *Profile) GetDefaultNodePoolName() string {
	return p.defaultNodePoolName
}

func (p *Profile) GetLocation() string {
	return p.Location
}

func (p *Profile) GetDefaultProfile() *pkgCluster.ClusterProfileResponse {

	nodepool := make(map[string]*pkgAKS.NodePoolCreate)
	nodepool[p.defaultNodePoolName] = &pkgAKS.NodePoolCreate{
		Autoscaling:      p.NodePools.Autoscaling,
		MinCount:         p.NodePools.MinCount,
		MaxCount:         p.NodePools.MaxCount,
		Count:            p.NodePools.Count,
		NodeInstanceType: p.NodePools.InstanceType,
	}

	return &pkgCluster.ClusterProfileResponse{
		Name:     "default", // todo const
		Location: p.Location,
		Cloud:    pkgCluster.Azure,
		Properties: &pkgCluster.ClusterProfileProperties{
			AKS: &pkgAKS.ClusterProfileAKS{
				KubernetesVersion: p.Version,
				NodePools:         nodepool,
			},
		},
	}
}
