package ec2

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgClusterEC2 "github.com/banzaicloud/pipeline/pkg/cluster/ec2"
	pkgDefaultsEC2 "github.com/banzaicloud/pipeline/pkg/profiles/defaults/ec2"
)

type Profile struct {
	defaultNodePoolName string
	*pkgDefaultsEC2.Defaults
	image string
}

func NewProfile(defaultNodePoolName string, ec2 *pkgDefaultsEC2.Defaults, image string) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		Defaults:            ec2,
		image:               image,
	}
}

func (p *Profile) GetDefaultProfile() *pkgCluster.CreateClusterRequest {

	nodepools := make(map[string]*pkgClusterEC2.NodePool)
	nodepools[p.defaultNodePoolName] = &pkgClusterEC2.NodePool{
		InstanceType: p.NodePools.InstanceType,
		SpotPrice:    p.NodePools.SpotPrice,
		Autoscaling:  p.NodePools.Autoscaling,
		MinCount:     p.NodePools.MinCount,
		MaxCount:     p.NodePools.MaxCount,
		Count:        p.NodePools.Count,
		Image:        p.image,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: p.Location,
		Cloud:    pkgCluster.Amazon,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterEC2: &pkgClusterEC2.CreateClusterEC2{
				NodePools: nodepools,
				Master: &pkgClusterEC2.CreateAmazonMaster{
					InstanceType: p.MasterInstanceType,
					Image:        p.image,
				},
			},
		},
	}
}
