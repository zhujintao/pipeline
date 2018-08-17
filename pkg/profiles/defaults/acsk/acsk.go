package acsk

type Defaults struct {
	Location                 string    `yaml:"location"`
	RegionId                 string    `yaml:"regionId"`
	ZoneId                   string    `yaml:"zoneId"`
	MasterInstanceType       string    `yaml:"masterInstanceType"`
	MasterSystemDiskCategory string    `yaml:"masterSystemDiskCategory"`
	NodePools                NodePools `yaml:"nodePools"`
}

type NodePools struct {
	Autoscaling        bool   `yaml:"autoscaling"`
	Count              int    `yaml:"count"`
	MinCount           int    `yaml:"minCount"`
	MaxCount           int    `yaml:"maxCount"`
	Image              string `yaml:"image"`
	InstanceType       string `yaml:"instanceType"`
	SystemDiskCategory string `yaml:"systemDiskCategory"`
}
