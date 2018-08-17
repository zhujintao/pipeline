package defaults

import (
	"io/ioutil"

	pkgDefaultsACSK "github.com/banzaicloud/pipeline/pkg/profiles/defaults/acsk"
	pkgDefaultsAKS "github.com/banzaicloud/pipeline/pkg/profiles/defaults/aks"
	pkgDefaultsEC2 "github.com/banzaicloud/pipeline/pkg/profiles/defaults/ec2"
	pkgDefaultsEKS "github.com/banzaicloud/pipeline/pkg/profiles/defaults/eks"
	pkgDefaultsGKE "github.com/banzaicloud/pipeline/pkg/profiles/defaults/gke"
	pkgDefaultsOKE "github.com/banzaicloud/pipeline/pkg/profiles/defaults/oke"
	"gopkg.in/yaml.v2"
)

type Manager struct {
	defaults *Defaults
	images   *AmazonImages
}

func (m *Manager) GetImages() (*AmazonImages, error) { // todo init??

	if m.images == nil || m.defaults == nil {
		var err error
		m.defaults, m.images, err = loadDefaults()
		if err != nil {
			return nil, err
		}
	}

	return m.images, nil
}

func (m *Manager) GetDefaults() (*Defaults, error) { // todo init??

	if m.defaults == nil || m.images == nil {
		var err error
		m.defaults, m.images, err = loadDefaults()
		if err != nil {
			return nil, err
		}
	}

	return m.defaults, nil
}

type Defaults struct {
	DefaultNodePoolName string                 `yaml:"defaultNodePoolName"`
	Distributions       DistributionProperties `yaml:"distributions"`
}

type DistributionProperties struct {
	ACSK pkgDefaultsACSK.Defaults `yaml:"acsk"`
	AKS  pkgDefaultsAKS.Defaults  `yaml:"aks"`
	EC2  pkgDefaultsEC2.Defaults  `yaml:"ec2"`
	EKS  pkgDefaultsEKS.Defaults  `yaml:"eks"`
	GKE  pkgDefaultsGKE.Defaults  `yaml:"gke"`
	OKE  pkgDefaultsOKE.Defaults  `yaml:"oke"`
}

func loadDefaults() (defaults *Defaults, images *AmazonImages, err error) {

	if err = readYaml("defaults/defaults.yaml", &defaults); err != nil { // todo move to const
		return
	}

	err = readYaml("defaults/defaults-amazon-images.yaml", &images) // todo move to const

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
