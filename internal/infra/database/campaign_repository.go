package database

import (
	"errors"
	"projetoEmail/internal/domain/campaign"
	internalerrors "projetoEmail/internal/internal_errors"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	tx := c.Db.Create(campaign)

	return tx.Error
}

func (c *CampaignRepository) All() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := c.Db.Model(&campaign.Campaign{}).Find(&campaigns)

	return campaigns, tx.Error
}

func (c *CampaignRepository) Get(id string) (campaign.Campaign, error) {
	var campaignData campaign.Campaign
	tx := c.Db.Model(&campaign.Campaign{}).Preload("Contacts").First(&campaignData, "id = ?", id)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return campaignData, internalerrors.NewErrEntityNotFound("campaign")
		}

		return campaignData, tx.Error
	}

	return campaignData, nil
}

func (c *CampaignRepository) Update(id string, values *campaign.Campaign) error {
	tx := c.Db.Model(&campaign.Campaign{}).Where("id = ?", id).UpdateColumns(values)

	return tx.Error
}

func (c *CampaignRepository) Delete(id string) error {
	tx := c.Db.Delete(&campaign.Campaign{}, "id = ?", id)

	return tx.Error
}
