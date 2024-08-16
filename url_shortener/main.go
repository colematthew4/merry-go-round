package main

import (
	"log"
	"net/http"

	"github.com/colematthew4/merry-go-round/url_shortener/config"
	"github.com/colematthew4/merry-go-round/url_shortener/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(
		postgres.Open(config.GetDbConfig().String()),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err.Error())
	}

	err = db.AutoMigrate(&data.Url{})
	if err != nil {
		log.Fatalln("Failed to apply database migrations:", err.Error())
	}

	log.Fatalln(http.ListenAndServe(":3000", nil))
}
