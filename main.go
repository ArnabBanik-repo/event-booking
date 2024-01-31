package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ArnabBanik-repo/event-booking/controllers"
	"github.com/ArnabBanik-repo/event-booking/initializers"
	"github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)


func init() {
  initializers.ConnectDB()

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

}

func main() {
	server := gin.Default()

	server.GET("/events", controllers.GetEvents)
	server.POST("/events", controllers.CreateEvent)

	err := server.Run()
	if err != nil {
		fmt.Println("Error in running the server")
		os.Exit(1)
	}

}
