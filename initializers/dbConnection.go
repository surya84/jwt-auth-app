package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {

	var err error
	dsn := os.Getenv("DB_CONN")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("error while connecting to DB %s", err)
	}
	log.Println("Connected to DataBase")

}
