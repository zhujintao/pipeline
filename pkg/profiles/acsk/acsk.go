package acsk

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgACSK "github.com/banzaicloud/pipeline/pkg/cluster/acsk"
	"github.com/banzaicloud/pipeline/pkg/providers"
)

type Profile struct {
	defaultNodePoolName string
	acsk                *Defaults
}

func NewProfile(defaultNodePoolName string, acsk *Defaults) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		acsk:                acsk,
	}
}

type Defaults struct {
	Location                 string    `yaml:"location"`
	RegionId                 string    `yaml:"regionId"`
	ZoneId                   string    `yaml:"zoneId"`
	MasterInstanceType       string    `yaml:"masterInstanceType"`
	MasterSystemDiskCategory string    `yaml:"masterSystemDiskCategory"`
	NodePools                NodePools `yaml:"nodePools"`
}

type NodePools struct {
	Autoscaling        bool   `yaml:"autoscaling"`
	Count              int    `yaml:"count"`
	MinCount           int    `yaml:"minCount"`
	MaxCount           int    `yaml:"maxCount"`
	Image              string `yaml:"image"`
	InstanceType       string `yaml:"instanceType"`
	SystemDiskCategory string `yaml:"systemDiskCategory"`
}

func (p *Profile) GetDefaultProfile() *pkgCluster.CreateClusterRequest {
	nodepools := make(pkgACSK.NodePools)
	nodepools[p.defaultNodePoolName] = &pkgACSK.NodePool{
		InstanceType:       p.acsk.NodePools.InstanceType,
		SystemDiskCategory: p.acsk.NodePools.SystemDiskCategory,
		//SystemDiskSize:     acsk.NodePools.SystemDiskSize,  // todo missing
		//LoginPassword:      acsk.NodePools.LoginPassword,  // todo missing
		Count: int(p.acsk.NodePools.Count),
		Image: p.acsk.NodePools.Image,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: p.acsk.Location,
		Cloud:    providers.Alibaba,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterACSK: &pkgACSK.CreateClusterACSK{
				RegionID:                 p.acsk.RegionId,
				ZoneID:                   p.acsk.ZoneId,
				MasterInstanceType:       p.acsk.MasterInstanceType,
				MasterSystemDiskCategory: p.acsk.MasterSystemDiskCategory,
				//MasterSystemDiskSize:     acsk.MasterSystemDiskSize, // todo missing
				//KeyPair:                  acsk.KeyPair, // todo missing
				NodePools: nodepools,
			},
		},
	}
}
