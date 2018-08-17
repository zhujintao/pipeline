package acsk

import (
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	pkgACSK "github.com/banzaicloud/pipeline/pkg/cluster/acsk"
	pkgDefaultsACSK "github.com/banzaicloud/pipeline/pkg/profiles/defaults/acsk"
	"github.com/banzaicloud/pipeline/pkg/providers"
)

type Profile struct {
	defaultNodePoolName string
	*pkgDefaultsACSK.Defaults
}

func NewProfile(defaultNodePoolName string, acsk *pkgDefaultsACSK.Defaults) *Profile {
	return &Profile{
		defaultNodePoolName: defaultNodePoolName,
		Defaults:            acsk,
	}
}

func (p *Profile) GetDefaultNodePoolName() string {
	return p.defaultNodePoolName
}

func (p *Profile) GetLocation() string {
	return p.Location
}

func (p *Profile) GetDefaultProfile() *pkgCluster.ClusterProfileResponse {
	nodepools := make(pkgACSK.NodePools)
	nodepools[p.defaultNodePoolName] = &pkgACSK.NodePool{
		InstanceType:       p.NodePools.InstanceType,
		SystemDiskCategory: p.NodePools.SystemDiskCategory,
		//SystemDiskSize:     acsk.NodePools.SystemDiskSize,  // todo missing
		//LoginPassword:      acsk.NodePools.LoginPassword,  // todo missing
		Count: int(p.NodePools.Count),
		Image: p.NodePools.Image,
	}

	return &pkgCluster.ClusterProfileResponse{
		Name:     "default", // todo const
		Location: p.Location,
		Cloud:    providers.Alibaba,
		Properties: &pkgCluster.ClusterProfileProperties{
			ACSK: &pkgACSK.ClusterProfileACSK{
				RegionID: p.RegionId,
				ZoneID:   p.ZoneId,
				//MasterInstanceType:       p.acsk.MasterInstanceType, // todo missing
				//MasterSystemDiskCategory: p.acsk.MasterSystemDiskCategory, // todo missing
				//MasterSystemDiskSize:     acsk.MasterSystemDiskSize, // todo missing
				//KeyPair:                  acsk.KeyPair, // todo missing
				NodePools: nodepools,
			},
		},
	}
}
