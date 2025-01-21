package main

import (
	"projetoEmail/internal/domain/campaign"
	"projetoEmail/internal/infra/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := database.NewDB()

	repository := database.CampaignRepository{Db: db}

	var campaigns []campaign.Campaign
	repository.Db.Where("status = ?", campaign.Pending).Find(&campaigns)

}
