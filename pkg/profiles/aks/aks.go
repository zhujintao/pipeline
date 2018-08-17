package aks

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgAKS "github.com/banzaicloud/pipeline/pkg/cluster/aks"
)

type Profile struct {
	defaultNodePoolName string
	*Defaults
}

func NewProfile(defaultNodePoolName string, aks *Defaults) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		Defaults:            aks,
	}
}

type Defaults struct {
	Location  string    `yaml:"location"`
	Version   string    `yaml:"version"`
	NodePools NodePools `yaml:"nodePools"`
}

type NodePools struct {
	Autoscaling  bool   `yaml:"autoscaling"`
	Count        int    `yaml:"count"`
	MinCount     int    `yaml:"minCount"`
	MaxCount     int    `yaml:"maxCount"`
	InstanceType string `yaml:"instanceType"`
}

func (p *Profile) GetDefaultProfile() *pkgCluster.CreateClusterRequest {

	nodepool := make(map[string]*pkgAKS.NodePoolCreate)
	nodepool[p.defaultNodePoolName] = &pkgAKS.NodePoolCreate{
		Autoscaling:      p.NodePools.Autoscaling,
		MinCount:         p.NodePools.MinCount,
		MaxCount:         p.NodePools.MaxCount,
		Count:            p.NodePools.Count,
		NodeInstanceType: p.NodePools.InstanceType,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: p.Location,
		Cloud:    pkgCluster.Azure,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterAKS: &pkgAKS.CreateClusterAKS{
				KubernetesVersion: p.Version,
				NodePools:         nodepool,
			},
		},
	}
}
