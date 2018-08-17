package gke

type Defaults struct {
	Location      string               `yaml:"location"`
	MasterVersion string               `yaml:"masterVersion"`
	NodeVersion   string               `yaml:"nodeVersion"`
	NodePools     DefaultsGKENodePools `yaml:"nodePools"`
}

type DefaultsGKENodePools struct {
	Autoscaling  bool   `yaml:"autoscaling"`
	Count        int    `yaml:"count"`
	MinCount     int    `yaml:"minCount"`
	MaxCount     int    `yaml:"maxCount"`
	InstanceType string `yaml:"instanceType"`
}
