package gke

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgGKE "github.com/banzaicloud/pipeline/pkg/cluster/gke"
)

type Profile struct {
	defaultNodePoolName string
	*Defaults
}

func NewProfile(defaultNodePoolName string, gke *Defaults) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		Defaults:            gke,
	}
}

type Defaults struct {
	Location      string               `yaml:"location"`
	MasterVersion string               `yaml:"masterVersion"`
	NodeVersion   string               `yaml:"nodeVersion"`
	NodePools     DefaultsGKENodePools `yaml:"nodePools"`
}

type DefaultsGKENodePools struct {
	Autoscaling  bool   `yaml:"autoscaling"`
	Count        int    `yaml:"count"`
	MinCount     int    `yaml:"minCount"`
	MaxCount     int    `yaml:"maxCount"`
	InstanceType string `yaml:"instanceType"`
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
