package aks

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgAKS "github.com/banzaicloud/pipeline/pkg/cluster/aks"
)

type Profile struct {
	defaultNodePoolName string
	aks                 *Defaults
}

func NewProfile(defaultNodePoolName string, aks *Defaults) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		aks:                 aks,
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
		Autoscaling:      p.aks.NodePools.Autoscaling,
		MinCount:         p.aks.NodePools.MinCount,
		MaxCount:         p.aks.NodePools.MaxCount,
		Count:            p.aks.NodePools.Count,
		NodeInstanceType: p.aks.NodePools.InstanceType,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: p.aks.Location,
		Cloud:    pkgCluster.Azure,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterAKS: &pkgAKS.CreateClusterAKS{
				KubernetesVersion: p.aks.Version, // todo simplify p.aks and so on other ones
				NodePools:         nodepool,
			},
		},
	}
}
