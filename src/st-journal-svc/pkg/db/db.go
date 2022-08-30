package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"st-journal-svc/pkg/models"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Trade{})
	db.AutoMigrate(&models.TradeDeleteLog{})

	return Handler{db}
}
