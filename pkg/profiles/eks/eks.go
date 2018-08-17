package eks

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgEC2 "github.com/banzaicloud/pipeline/pkg/cluster/ec2"
	pkgEKS "github.com/banzaicloud/pipeline/pkg/cluster/eks"
	pkgDefaultsEKS "github.com/banzaicloud/pipeline/pkg/profiles/defaults/eks"
)

type Profile struct {
	defaultNodePoolName string
	*pkgDefaultsEKS.Defaults
	image string
}

func NewProfile(defaultNodePoolName string, eks *pkgDefaultsEKS.Defaults, image string) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		Defaults:            eks,
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

	nodepools := make(map[string]*pkgEC2.NodePool)
	nodepools[p.defaultNodePoolName] = &pkgEC2.NodePool{
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
			EKS: &pkgEKS.ClusterProfileEKS{
				Version:   p.Version,
				NodePools: nodepools,
			},
		},
	}

}
