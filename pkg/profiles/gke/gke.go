package gke

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgGKE "github.com/banzaicloud/pipeline/pkg/cluster/gke"
)

type Profile struct {
	defaultNodePoolName string
	gke                 *Defaults
}

func NewProfile(defaultNodePoolName string, gke *Defaults) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		gke:                 gke,
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
		Autoscaling:      p.gke.NodePools.Autoscaling,
		MinCount:         p.gke.NodePools.MinCount,
		MaxCount:         p.gke.NodePools.MaxCount,
		Count:            p.gke.NodePools.Count,
		NodeInstanceType: p.gke.NodePools.InstanceType,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: p.gke.Location,
		Cloud:    pkgCluster.Google,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterGKE: &pkgGKE.CreateClusterGKE{
				NodeVersion: p.gke.NodeVersion,
				NodePools:   nodepools,
				Master: &pkgGKE.Master{
					Version: p.gke.MasterVersion,
				},
			},
		},
	}

}
