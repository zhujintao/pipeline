package ec2

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgClusterEC2 "github.com/banzaicloud/pipeline/pkg/cluster/ec2"
)

type Profile struct {
	defaultNodePoolName string
	ec2                 *Defaults
	image               string
}

func NewProfile(defaultNodePoolName string, ec2 *Defaults, image string) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		ec2:                 ec2,
		image:               image,
	}
}

type Defaults struct {
	Location           string    `yaml:"location"`
	MasterInstanceType string    `yaml:"masterInstanceType"`
	NodePools          NodePools `yaml:"nodePools"`
}

type NodePools struct {
	InstanceType string `yaml:"instanceType"`
	SpotPrice    string `yaml:"spotPrice"`
	Autoscaling  bool   `yaml:"autoscaling"`
	Count        int    `yaml:"count"`
	MinCount     int    `yaml:"minCount"`
	MaxCount     int    `yaml:"maxCount"`
}

func (p *Profile) GetDefaultProfile() *pkgCluster.CreateClusterRequest {

	nodepools := make(map[string]*pkgClusterEC2.NodePool)
	nodepools[p.defaultNodePoolName] = &pkgClusterEC2.NodePool{
		InstanceType: p.ec2.NodePools.InstanceType,
		SpotPrice:    p.ec2.NodePools.SpotPrice,
		Autoscaling:  p.ec2.NodePools.Autoscaling,
		MinCount:     p.ec2.NodePools.MinCount,
		MaxCount:     p.ec2.NodePools.MaxCount,
		Count:        p.ec2.NodePools.Count,
		Image:        p.image,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: p.ec2.Location,
		Cloud:    pkgCluster.Amazon,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterEC2: &pkgClusterEC2.CreateClusterEC2{
				NodePools: nodepools,
				Master: &pkgClusterEC2.CreateAmazonMaster{
					InstanceType: p.ec2.MasterInstanceType,
					Image:        p.image,
				},
			},
		},
	}
}
