package profiles

import (
	"errors"

	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	acsk2 "github.com/banzaicloud/pipeline/pkg/cluster/acsk"
	aks2 "github.com/banzaicloud/pipeline/pkg/cluster/aks"
	"github.com/banzaicloud/pipeline/pkg/profiles/defaults"
	pkgProfileEC2 "github.com/banzaicloud/pipeline/pkg/profiles/ec2"
	pkgProfileEKS "github.com/banzaicloud/pipeline/pkg/profiles/eks"
	pkgProfileGKE "github.com/banzaicloud/pipeline/pkg/profiles/gke"
	"github.com/banzaicloud/pipeline/pkg/providers"
	oke2 "github.com/banzaicloud/pipeline/pkg/providers/oracle/cluster"
)

type ProfileManager interface {
	GetDefaultProfile() *pkgCluster.CreateClusterRequest
}

func getProfileManager(distributionType string) (ProfileManager, error) {

	var manager defaults.Manager
	def, err := manager.GetDefaults()
	if err != nil {
		return nil, err
	}

	images, err := manager.GetImages()
	if err != nil {
		return nil, err
	}

	switch distributionType {
	case pkgCluster.EC2:
		return pkgProfileEC2.NewProfile(def.DefaultNodePoolName, &def.Distributions.EC2, images.EC2.GetDefaultAmazonImage(def.Distributions.EC2.Location)), nil // todo refactor!!
	case pkgCluster.EKS:
		return pkgProfileEKS.NewProfile(def.DefaultNodePoolName, &def.Distributions.EKS, images.EKS.GetDefaultAmazonImage(def.Distributions.EKS.Location)), nil // todo refactor!!
	case pkgCluster.GKE:
		return pkgProfileGKE.NewProfile(def.DefaultNodePoolName, &def.Distributions.GKE), nil
	}

	return nil, errors.New("not supported distribution type")
}

func GetDefaultProfile(distributionType string) (*pkgCluster.CreateClusterRequest, error) {

	manager, err := getProfileManager(distributionType)
	if err != nil {
		return nil, err
	}

	return manager.GetDefaultProfile(), nil

	//defaults, images, err := readFiles()
	//if err != nil {
	//	return nil, err
	//}
	//
	//switch distributionType {
	//case pkgCluster.ACSK:
	//	return createACSKRequest(&defaults.Distributions.ACSK, defaults.DefaultNodePoolName), nil
	//case pkgCluster.AKS:
	//	return createAKSRequest(&defaults.Distributions.AKS, defaults.DefaultNodePoolName), nil
	//case pkgCluster.EC2:
	//	return createEC2Request(&defaults.Distributions.EC2, defaults.DefaultNodePoolName, images), nil
	//case pkgCluster.EKS:
	//	return createEKSRequest(&defaults.Distributions.EKS, defaults.DefaultNodePoolName, images), nil
	//	case pkgCluster.GKE:
	//		return createGKERequest(&defaults.Distributions.GKE, defaults.DefaultNodePoolName), nil
	//case pkgCluster.OKE:
	//	return createOKERequest(&defaults.Distributions.OKE, defaults.DefaultNodePoolName), nil
	//
	//}
	//
	//return nil, errors.New("not supported distribution")
}

func createOKERequest(oke *DefaultsOKE, defaultNodePoolName string) *pkgCluster.CreateClusterRequest {

	nodepools := make(map[string]*oke2.NodePool)
	nodepools[defaultNodePoolName] = &oke2.NodePool{
		Version: oke.NodePools.Version,
		Count:   uint(oke.NodePools.Count),
		Image:   oke.NodePools.Image,
		Shape:   oke.NodePools.Shape,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: oke.Location,
		Cloud:    pkgCluster.Oracle,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterOKE: &oke2.Cluster{
				Version:   oke.Version,
				NodePools: nodepools,
			},
		},
	}
}

func createAKSRequest(aks *DefaultsAKS, defaultNodePoolName string) *pkgCluster.CreateClusterRequest {

	nodepool := make(map[string]*aks2.NodePoolCreate)
	nodepool[defaultNodePoolName] = &aks2.NodePoolCreate{
		Autoscaling:      aks.NodePools.Autoscaling,
		MinCount:         aks.NodePools.MinCount,
		MaxCount:         aks.NodePools.MaxCount,
		Count:            aks.NodePools.Count,
		NodeInstanceType: aks.NodePools.InstanceType,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: aks.Location,
		Cloud:    pkgCluster.Azure,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterAKS: &aks2.CreateClusterAKS{
				KubernetesVersion: aks.Version,
				NodePools:         nodepool,
			},
		},
	}
}

func createACSKRequest(acsk *DefaultsACSK, defaultNodePoolName string) *pkgCluster.CreateClusterRequest {
	nodepools := make(acsk2.NodePools)
	nodepools[defaultNodePoolName] = &acsk2.NodePool{
		InstanceType:       acsk.NodePools.InstanceType,
		SystemDiskCategory: acsk.NodePools.SystemDiskCategory,
		//SystemDiskSize:     acsk.NodePools.SystemDiskSize,  // todo missing
		//LoginPassword:      acsk.NodePools.LoginPassword,  // todo missing
		Count: int(acsk.NodePools.Count),
		Image: acsk.NodePools.Image,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: acsk.Location,
		Cloud:    providers.Alibaba,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterACSK: &acsk2.CreateClusterACSK{
				RegionID:                 acsk.RegionId,
				ZoneID:                   acsk.ZoneId,
				MasterInstanceType:       acsk.MasterInstanceType,
				MasterSystemDiskCategory: acsk.MasterSystemDiskCategory,
				//MasterSystemDiskSize:     acsk.MasterSystemDiskSize, // todo missing
				//KeyPair:                  acsk.KeyPair, // todo missing
				NodePools: nodepools,
			},
		},
	}
}

type DefaultsACSK struct {
	Location                 string                `yaml:"location"`
	RegionId                 string                `yaml:"regionId"`
	ZoneId                   string                `yaml:"zoneId"`
	MasterInstanceType       string                `yaml:"masterInstanceType"`
	MasterSystemDiskCategory string                `yaml:"masterSystemDiskCategory"`
	NodePools                DefaultsACSKNodePools `yaml:"nodePools"`
}

type DefaultsAKS struct {
	Location  string               `yaml:"location"`
	Version   string               `yaml:"version"`
	NodePools DefaultsAKSNodePools `yaml:"nodePools"`
}

type DefaultsOKE struct {
	Location  string               `yaml:"location"`
	Version   string               `yaml:"version"`
	NodePools DefaultsOKENodePools `yaml:"nodePools"`
}

type DefaultsACSKNodePools struct {
	Autoscaling        bool   `yaml:"autoscaling"`
	Count              int    `yaml:"count"`
	MinCount           int    `yaml:"minCount"`
	MaxCount           int    `yaml:"maxCount"`
	Image              string `yaml:"image"`
	InstanceType       string `yaml:"instanceType"`
	SystemDiskCategory string `yaml:"systemDiskCategory"`
}

type DefaultsAKSNodePools struct {
	Autoscaling  bool   `yaml:"autoscaling"`
	Count        int    `yaml:"count"`
	MinCount     int    `yaml:"minCount"`
	MaxCount     int    `yaml:"maxCount"`
	InstanceType string `yaml:"instanceType"`
}

type DefaultsOKENodePools struct {
	Version  string `yaml:"version"`
	Count    int    `yaml:"count"`
	MinCount int    `yaml:"minCount"`
	MaxCount int    `yaml:"maxCount"`
	Image    string `yaml:"image"`
	Shape    string `yaml:"shape"`
}
