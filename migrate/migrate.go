package main

import (
	"github.com/ArnabBanik-repo/event-booking/initializers"
	"github.com/ArnabBanik-repo/event-booking/models"
	"github.com/joho/godotenv"
)

func init() {
  godotenv.Load()
  initializers.ConnectDB()
}

func main() {
  initializers.DB.AutoMigrate(&models.Event{}, &models.User{})
}
