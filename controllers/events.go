package controllers

import (
	"net/http"
	"time"

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

  result := initializers.DB.First(&event, "id = ?", id)
  if result.Error != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "err": "No records with that ID"})
    return
  }

	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": event})
}

func CreateEvent(c *gin.Context) {
	var e models.Event
  err := c.ShouldBindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "err": err.Error()})
		return
	}

	e.ID = uuid.NewString()
  ui, _  := c.Get("user")
  u := ui.(models.User)
  e.UserId = u.ID

	result := initializers.DB.Create(&e)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "err": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "ok", "data": e})
}

func UpdateEvent(c *gin.Context) {

  id := c.Param("id")
  var event models.Event

  result := initializers.DB.First(&event, "id = ?", id)
  if result.Error != nil {
    c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "err": result.Error.Error()})
    return
  }

  type e struct {
    Name string
    Description string
    Location string
    DateTime time.Time
  }

  var updatedEvent e

  err := c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "err": err.Error()})
		return
	}

  result = initializers.DB.Model(&event).Updates(updatedEvent)
  if result.Error != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "err": result.Error.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"status": "success", "data": event})
}

func DeleteEvent(c *gin.Context){
  id := c.Param("id")
  event := models.Event{ID: id}

  result := initializers.DB.Delete(&event)
  if result.Error != nil || result.RowsAffected == 0 {
    c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "err": "That record could not be deleted"})
    return
  }

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
