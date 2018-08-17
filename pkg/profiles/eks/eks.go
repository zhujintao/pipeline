package eks

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgEC2 "github.com/banzaicloud/pipeline/pkg/cluster/ec2"
	pkgEKS "github.com/banzaicloud/pipeline/pkg/cluster/eks"
	pkgEC2Profile "github.com/banzaicloud/pipeline/pkg/profiles/ec2"
)

type Profile struct {
	defaultNodePoolName string
	*Defaults
	image string
}

type Defaults struct {
	Location  string                  `yaml:"location"`
	Version   string                  `yaml:"version"`
	NodePools pkgEC2Profile.NodePools `yaml:"nodePools"`
}

func NewProfile(defaultNodePoolName string, eks *Defaults, image string) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		Defaults:            eks,
		image:               image,
	}
}

func (p *Profile) GetDefaultProfile() *pkgCluster.CreateClusterRequest {

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

	return &pkgCluster.CreateClusterRequest{
		Location: p.Location,
		Cloud:    pkgCluster.Amazon,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterEKS: &pkgEKS.CreateClusterEKS{
				Version:   p.Version,
				NodePools: nodepools,
			},
		},
	}

}
