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

func (p *Profile) GetDefaultNodePoolName() string {
	return p.defaultNodePoolName
}

func (p *Profile) GetLocation() string {
	return p.Location
}

func (p *Profile) GetDefaultProfile() *pkgCluster.ClusterProfileResponse {

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

	return &pkgCluster.ClusterProfileResponse{
		Name:     "default", // todo const
		Location: p.Location,
		Cloud:    pkgCluster.Amazon,
		Properties: &pkgCluster.ClusterProfileProperties{
			EC2: &pkgClusterEC2.ClusterProfileEC2{
				NodePools: nodepools,
				Master: &pkgClusterEC2.ProfileMaster{
					InstanceType: p.MasterInstanceType,
					Image:        p.image,
				},
			},
		},
	}
}
