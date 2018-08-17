package defaults

import (
	"github.com/banzaicloud/pipeline/config"
	pkgCluster "github.com/banzaicloud/pipeline/pkg/cluster"
	"github.com/banzaicloud/pipeline/pkg/cluster/gke"
	"github.com/banzaicloud/pipeline/pkg/profiles"
)

// GKEProfile describes a Google cluster profile
type GKEProfile struct {
	DefaultModel
	Location      string
	NodeVersion   string
	MasterVersion string
	NodePools     []*GKENodePoolProfile `gorm:"foreignkey:Name"`
}

// GKENodePoolProfile describes a Google cluster profile's nodepools
type GKENodePoolProfile struct {
	ID               uint   `gorm:"primary_key"`
	Autoscaling      bool   `gorm:"default:false"`
	MinCount         int    `gorm:"default:1"`
	MaxCount         int    `gorm:"default:2"`
	Count            int    `gorm:"default:1"`
	NodeInstanceType string `gorm:"default:'n1-standard-1'"`
	Name             string `gorm:"unique_index:idx_model_name"`
	NodeName         string `gorm:"unique_index:idx_model_name"`
}

// TableName overrides GKEProfile's table name
func (GKEProfile) TableName() string {
	return DefaultGKEProfileTableName
}

// TableName overrides GKENodePoolProfile's table name
func (GKENodePoolProfile) TableName() string {
	return DefaultGKENodePoolProfileTableName
}

// AfterFind loads nodepools to profile
func (p *GKEProfile) AfterFind() error {
	log.Info("AfterFind gke profile... load node pools")
	return config.DB().Where(GKENodePoolProfile{Name: p.Name}).Find(&p.NodePools).Error
}

// BeforeSave clears nodepools
func (p *GKEProfile) BeforeSave() error {
	log.Info("BeforeSave gke profile...")

	var nodePools []*GKENodePoolProfile
	err := config.DB().Where(GKENodePoolProfile{
		Name: p.Name,
	}).Find(&nodePools).Delete(&nodePools).Error
	if err != nil {
		log.Errorf("Error during deleting saved nodepools: %s", err.Error())
	}

	p.addDefaults()

	return nil
}

func (p *GKEProfile) addDefaults() error {

	defaultProf, err := profiles.GetDefaultProfileManager(p.GetDistribution())
	if err != nil {
		return err
	}

	// set location
	if len(p.Location) == 0 {
		p.Location = defaultProf.GetLocation()
	}

	// set node version
	if len(p.NodeVersion) == 0 {
		p.NodeVersion = defaultProf.NodeVersion
	}

	// set master version
	if len(p.MasterVersion) == 0 {
		p.MasterVersion = gke.Master.Version
	}

	// set nodepools
	if len(p.NodePools) == 0 {
		for name, np := range gke.NodePools {
			p.NodePools = append(p.NodePools, &GKENodePoolProfile{
				Autoscaling:      np.Autoscaling,
				MinCount:         np.MinCount,
				MaxCount:         np.MaxCount,
				Count:            np.Count,
				NodeInstanceType: np.NodeInstanceType,
				Name:             p.Name,
				NodeName:         name,
			})
		}
	}

	// check nodepools
	for _, np := range p.NodePools {
		if len(np.Name) == 0 {
			gke.NodePools[defaultProf.]
		}
	}

	return nil
}

// BeforeDelete deletes all nodepools to belongs to profile
func (p *GKEProfile) BeforeDelete() error {
	log.Info("BeforeDelete gke profile... delete all nodepool")

	var nodePools []*GKENodePoolProfile
	return config.DB().Where(GKENodePoolProfile{
		Name: p.Name,
	}).Find(&nodePools).Delete(&nodePools).Error
}

// SaveInstance saves cluster profile into database
func (p *GKEProfile) SaveInstance() error {
	return save(p)
}

// IsDefinedBefore returns true if database contains en entry with profile name
func (p *GKEProfile) IsDefinedBefore() bool {
	return config.DB().First(&p).RowsAffected != int64(0)
}

// GetCloud returns profile's cloud type
func (p *GKEProfile) GetCloud() string {
	return pkgCluster.Google
}

// GetDistribution returns profile's distribution type
func (p *GKEProfile) GetDistribution() string {
	return pkgCluster.GKE
}

// GetProfile load profile from database and converts ClusterProfileResponse
func (p *GKEProfile) GetProfile() *pkgCluster.ClusterProfileResponse {

	nodePools := make(map[string]*gke.NodePool)
	if p.NodePools != nil {
		for _, np := range p.NodePools {
			nodePools[np.NodeName] = &gke.NodePool{
				Autoscaling:      np.Autoscaling,
				MinCount:         np.MinCount,
				MaxCount:         np.MaxCount,
				Count:            np.Count,
				NodeInstanceType: np.NodeInstanceType,
			}
		}
	}

	return &pkgCluster.ClusterProfileResponse{
		Name:     p.DefaultModel.Name,
		Location: p.Location,
		Cloud:    pkgCluster.Google,
		Properties: &pkgCluster.ClusterProfileProperties{
			GKE: &gke.ClusterProfileGKE{
				Master: &gke.Master{
					Version: p.MasterVersion,
				},
				NodeVersion: p.NodeVersion,
				NodePools:   nodePools,
			},
		},
	}
}

// UpdateProfile update profile's data with ClusterProfileRequest's data and if bool is true then update in the database
func (p *GKEProfile) UpdateProfile(r *pkgCluster.ClusterProfileRequest, withSave bool) error {

	if len(r.Location) != 0 {
		p.Location = r.Location
	}

	if r.Properties.GKE != nil {

		if len(r.Properties.GKE.NodeVersion) != 0 {
			p.NodeVersion = r.Properties.GKE.NodeVersion
		}

		if len(r.Properties.GKE.NodePools) != 0 {

			var nodePools []*GKENodePoolProfile
			for name, np := range r.Properties.GKE.NodePools {
				nodePools = append(nodePools, &GKENodePoolProfile{
					Autoscaling:      np.Autoscaling,
					MinCount:         np.MinCount,
					MaxCount:         np.MaxCount,
					Count:            np.Count,
					NodeInstanceType: np.NodeInstanceType,
					Name:             p.Name,
					NodeName:         name,
				})
			}

			p.NodePools = nodePools
		}

		if r.Properties.GKE.Master != nil {
			p.MasterVersion = r.Properties.GKE.Master.Version
		}
	}

	if withSave {
		return p.SaveInstance()
	}
	p.Name = r.Name
	return nil
}

// DeleteProfile deletes cluster profile from database
func (p *GKEProfile) DeleteProfile() error {
	return config.DB().Delete(&p).Error
}
