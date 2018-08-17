package profiles

import (
	"errors"
	"io/ioutil"

	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	acsk2 "github.com/banzaicloud/pipeline/pkg/cluster/acsk"
	aks2 "github.com/banzaicloud/pipeline/pkg/cluster/aks"
	ec22 "github.com/banzaicloud/pipeline/pkg/cluster/ec2"
	eks2 "github.com/banzaicloud/pipeline/pkg/cluster/eks"
	pkgProfileGKE "github.com/banzaicloud/pipeline/pkg/profiles/gke"
	"github.com/banzaicloud/pipeline/pkg/providers"
	oke2 "github.com/banzaicloud/pipeline/pkg/providers/oracle/cluster"
	"gopkg.in/yaml.v2"
)

type ProfileManager interface {
	GetDefaultProfile() *pkgCluster.CreateClusterRequest
}

func getProfileManager(distributionType string) (ProfileManager, error) {
	defaults, _, err := readFiles()
	if err != nil {
		return nil, err
	}

	switch distributionType {
	case pkgCluster.GKE:
		return pkgProfileGKE.NewProfile(defaults.DefaultNodePoolName, &defaults.Distributions.GKE), nil
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

func createEKSRequest(eks *DefaultsEKS, defaultNodePoolName string, images DefaultAmazonImages) *pkgCluster.CreateClusterRequest {

	image := getAmazonImage(images.EKS, eks.Location)

	nodepools := make(map[string]*ec22.NodePool)
	nodepools[defaultNodePoolName] = &ec22.NodePool{
		InstanceType: eks.NodePools.InstanceType,
		SpotPrice:    eks.NodePools.SpotPrice,
		Autoscaling:  eks.NodePools.Autoscaling,
		MinCount:     eks.NodePools.MinCount,
		MaxCount:     eks.NodePools.MaxCount,
		Count:        eks.NodePools.Count,
		Image:        image,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: eks.Location,
		Cloud:    pkgCluster.Amazon,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterEKS: &eks2.CreateClusterEKS{
				Version:   eks.Version,
				NodePools: nodepools,
			},
		},
	}

}

func createEC2Request(ec2 *DefaultsEC2, defaultNodePoolName string, images DefaultAmazonImages) *pkgCluster.CreateClusterRequest {

	image := getAmazonImage(images.EC2, ec2.Location)

	nodepools := make(map[string]*ec22.NodePool)
	nodepools[defaultNodePoolName] = &ec22.NodePool{
		InstanceType: ec2.NodePools.InstanceType,
		SpotPrice:    ec2.NodePools.SpotPrice,
		Autoscaling:  ec2.NodePools.Autoscaling,
		MinCount:     ec2.NodePools.MinCount,
		MaxCount:     ec2.NodePools.MaxCount,
		Count:        ec2.NodePools.Count,
		Image:        image,
	}

	return &pkgCluster.CreateClusterRequest{
		Location: ec2.Location,
		Cloud:    pkgCluster.Amazon,
		Properties: &pkgCluster.CreateClusterProperties{
			CreateClusterEC2: &ec22.CreateClusterEC2{
				NodePools: nodepools,
				Master: &ec22.CreateAmazonMaster{
					InstanceType: ec2.MasterInstanceType,
					Image:        image,
				},
			},
		},
	}
}

func getAmazonImage(images AmazonImages, location string) string {
	return images[location]
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

func readFiles() (defaults Defaults, images DefaultAmazonImages, err error) {

	if err = readYaml("defaults/defaults.yaml", &defaults); err != nil {
		return
	}

	err = readYaml("defaults/defaults-amazon-images.yaml", &images)

	return
}

func readYaml(filePath string, out interface{}) error {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, out)
	if err != nil {
		return err
	}

	return nil
}

type DefaultAmazonImages struct {
	EC2 AmazonImages `yaml:"ec2"`
	EKS AmazonImages `yaml:"eks"`
}

type AmazonImages map[string]string

type Defaults struct {
	DefaultNodePoolName string               `yaml:"defaultNodePoolName"`
	Distributions       DefaultsDistribution `yaml:"distributions"`
}

type DefaultsDistribution struct {
	ACSK DefaultsACSK           `yaml:"acsk"`
	AKS  DefaultsAKS            `yaml:"aks"`
	EC2  DefaultsEC2            `yaml:"ec2"`
	EKS  DefaultsEKS            `yaml:"eks"`
	GKE  pkgProfileGKE.Defaults `yaml:"gke"`
	OKE  DefaultsOKE            `yaml:"oke"`
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

type DefaultsEC2 struct {
	Location           string                  `yaml:"location"`
	MasterInstanceType string                  `yaml:"masterInstanceType"`
	NodePools          DefaultsAmazonNodePools `yaml:"nodePools"`
}

type DefaultsEKS struct {
	Location  string                  `yaml:"location"`
	Version   string                  `yaml:"version"`
	NodePools DefaultsAmazonNodePools `yaml:"nodePools"`
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

type DefaultsAmazonNodePools struct {
	InstanceType string `yaml:"instanceType"`
	SpotPrice    string `yaml:"spotPrice"`
	Autoscaling  bool   `yaml:"autoscaling"`
	Count        int    `yaml:"count"`
	MinCount     int    `yaml:"minCount"`
	MaxCount     int    `yaml:"maxCount"`
}

type DefaultsOKENodePools struct {
	Version  string `yaml:"version"`
	Count    int    `yaml:"count"`
	MinCount int    `yaml:"minCount"`
	MaxCount int    `yaml:"maxCount"`
	Image    string `yaml:"image"`
	Shape    string `yaml:"shape"`
}
