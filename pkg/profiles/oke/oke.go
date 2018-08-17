package oke

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgDefaultsOKE "github.com/banzaicloud/pipeline/pkg/profiles/defaults/oke"
	pkgOKE "github.com/banzaicloud/pipeline/pkg/providers/oracle/cluster"
)

type Profile struct {
	defaultNodePoolName string
	*pkgDefaultsOKE.Defaults // todo all the same???
}

func NewProfile(defaultNodePoolName string, oke *pkgDefaultsOKE.Defaults) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		Defaults:            oke,
	}
}

func (p *Profile) GetDefaultProfile() *pkgCluster.CreateClusterRequest {

	nodepools := make(map[string]*pkgOKE.NodePool)
	nodepools[p.defaultNodePoolName] = &pkgOKE.NodePool{
		Version: p.NodePools.Version,
		Count:   uint(p.NodePools.Count),
		Image:   p.NodePools.Image,
		Shape:   p.NodePools.Shape,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: p.Location,
		Cloud:    pkgCluster.Oracle,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterOKE: &pkgOKE.Cluster{
				Version:   p.Version,
				NodePools: nodepools,
			},
		},
	}
}
