package main

import (
	"github.com/ArnabBanik-repo/event-booking/initializers"
	"github.com/ArnabBanik-repo/event-booking/models"
)

func init() {
  initializers.ConnectDB()
}

func main() {
  initializers.DB.AutoMigrate(&models.Event{})
}
