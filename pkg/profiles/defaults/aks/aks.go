package aks

type Defaults struct {
	Location  string    `yaml:"location"`
	Version   string    `yaml:"version"`
	NodePools NodePools `yaml:"nodePools"`
}

type NodePools struct {
	Autoscaling  bool   `yaml:"autoscaling"`
	Count        int    `yaml:"count"`
	MinCount     int    `yaml:"minCount"`
	MaxCount     int    `yaml:"maxCount"`
	InstanceType string `yaml:"instanceType"`
}
