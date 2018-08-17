package eks

import (
	pkgDefaultsEC2 "github.com/banzaicloud/pipeline/pkg/profiles/defaults/ec2"
)

type Defaults struct {
	Location  string                   `yaml:"location"`
	Version   string                   `yaml:"version"`
	NodePools pkgDefaultsEC2.NodePools `yaml:"nodePools"`
}
