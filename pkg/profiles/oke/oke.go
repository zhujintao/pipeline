package oke

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgOKE "github.com/banzaicloud/pipeline/pkg/providers/oracle/cluster"
)

type Profile struct {
	defaultNodePoolName string
	oke                 *Defaults // todo all the same???
}

func NewProfile(defaultNodePoolName string, oke *Defaults) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		oke:                 oke,
	}
}

type Defaults struct {
	Location  string               `yaml:"location"`
	Version   string               `yaml:"version"`
	NodePools DefaultsOKENodePools `yaml:"nodePools"`
}

type DefaultsOKENodePools struct {
	Version  string `yaml:"version"`
	Count    int    `yaml:"count"`
	MinCount int    `yaml:"minCount"`
	MaxCount int    `yaml:"maxCount"`
	Image    string `yaml:"image"`
	Shape    string `yaml:"shape"`
}

func (p *Profile) GetDefaultProfile() *pkgCluster.CreateClusterRequest {

	nodepools := make(map[string]*pkgOKE.NodePool)
	nodepools[p.defaultNodePoolName] = &pkgOKE.NodePool{
		Version: p.oke.NodePools.Version,
		Count:   uint(p.oke.NodePools.Count),
		Image:   p.oke.NodePools.Image,
		Shape:   p.oke.NodePools.Shape,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: p.oke.Location,
		Cloud:    pkgCluster.Oracle,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterOKE: &pkgOKE.Cluster{
				Version:   p.oke.Version,
				NodePools: nodepools,
			},
		},
	}
}
