package defaults

type AmazonImages struct {
	EC2 Images `yaml:"ec2"`
	EKS Images `yaml:"eks"`
}

type Images map[string]string

func (m Images) GetDefaultAmazonImage(location string) string {
	return m[location]
}
