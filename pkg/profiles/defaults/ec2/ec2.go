package ec2

type Defaults struct {
	Location           string    `yaml:"location"`
	MasterInstanceType string    `yaml:"masterInstanceType"`
	NodePools          NodePools `yaml:"nodePools"`
}

type NodePools struct {
	InstanceType string `yaml:"instanceType"`
	SpotPrice    string `yaml:"spotPrice"`
	Autoscaling  bool   `yaml:"autoscaling"`
	Count        int    `yaml:"count"`
	MinCount     int    `yaml:"minCount"`
	MaxCount     int    `yaml:"maxCount"`
}

