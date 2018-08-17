package defaults

import (
	"io/ioutil"

	pkgProfileEC2 "github.com/banzaicloud/pipeline/pkg/profiles/ec2"
	pkgProfileGKE "github.com/banzaicloud/pipeline/pkg/profiles/gke"
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
	//ACSK DefaultsACSK           `yaml:"acsk"`
	//AKS  DefaultsAKS            `yaml:"aks"`
	EC2 pkgProfileEC2.Defaults `yaml:"ec2"`
	//EKS  DefaultsEKS            `yaml:"eks"` // todo put back
	GKE pkgProfileGKE.Defaults `yaml:"gke"`
	//OKE  DefaultsOKE            `yaml:"oke"`
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
