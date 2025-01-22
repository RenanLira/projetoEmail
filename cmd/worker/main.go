package main

import (
	"fmt"
	"projetoEmail/internal/domain/campaign"
	"projetoEmail/internal/infra/database"
	"projetoEmail/internal/infra/mail"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB
	repository database.CampaignRepository
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {

	db = database.NewDB()
	repository = database.CampaignRepository{Db: db}

	service := &campaign.Service{Repository: &repository, SendMail: mail.SendMail}

	campaigns := getCampaignsToBeSent()

	for _, c := range campaigns {
		fmt.Println("Sending campaign: ", c.Name)
		
		go func(c campaign.Campaign) {
			err := service.SendMailAndUpdateStatus(&c)
			if err != nil {
				fmt.Println("Error sending campaign: ", c.Name)
			}
		}(c)
	}

}

func getCampaignsToBeSent() []campaign.Campaign {
	var campaigns []campaign.Campaign
	tx := repository.Db.Preload("Contacts").Where("status = ?", campaign.Started).Where("date_part('minute', now()::timestamp - update_at::timestamp) > 3").Find(&campaigns)

	if tx.Error != nil {
		fmt.Println("Error fetching campaigns: ", tx.Error)
	}

	return campaigns
}
