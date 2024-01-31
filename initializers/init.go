package initializers

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dn := os.Getenv("DB_URL")
	if len(dn) == 0 {
		log.Fatal("Database url not present in env file")
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(dn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to database")
	}
}
