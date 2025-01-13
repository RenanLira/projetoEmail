package database

import (
	"os"
	"projetoEmail/internal/domain/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
