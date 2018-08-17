package acsk

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgACSK "github.com/banzaicloud/pipeline/pkg/cluster/acsk"
	pkgDefaultsACSK "github.com/banzaicloud/pipeline/pkg/profiles/defaults/acsk"
	"github.com/banzaicloud/pipeline/pkg/providers"
)

type Profile struct {
	defaultNodePoolName string
	acsk                *pkgDefaultsACSK.Defaults
}

func NewProfile(defaultNodePoolName string, acsk *pkgDefaultsACSK.Defaults) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		acsk:                acsk,
	}
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
