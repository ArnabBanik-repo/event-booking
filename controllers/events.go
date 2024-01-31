package controllers

import (
	"net/http"

	"github.com/ArnabBanik-repo/event-booking/initializers"
	"github.com/ArnabBanik-repo/event-booking/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetEvents(c *gin.Context) {

	var events []models.Event
	result := initializers.DB.Find(&events)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "err": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": events, "count": result.RowsAffected})
}

func GetEvent(c *gin.Context) {
  id := c.Param("id")
  var event models.Event
  result := initializers.DB.First(&event, id)
  if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "err": result.Error.Error()})
    return 
  }

	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": event, "count": result.RowsAffected})
}

func CreateEvent(c *gin.Context) {
	var e models.Event
	err := c.BindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "err": err.Error()})
		return
	}

	e.ID = uuid.NewString()
	e.UserId = uuid.NewString()

	result := initializers.DB.Create(&e)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "err": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "ok", "data": e})
}
