package controllers

import (
	"net/http"

	"github.com/ArnabBanik-repo/event-booking/initializers"
	"github.com/ArnabBanik-repo/event-booking/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetEvents(c *gin.Context) {

	result := initializers.DB.Find(&models.Event{})
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "err": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": result})
}

func CreateEvent(c *gin.Context) {
	var e models.Event
	err := c.BindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "err": err.Error()})
		return
	}

	e.Id, err = uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "err": err.Error()})
		return
	}
	e.UserId = "142"
	result := initializers.DB.Create(&models.Event{})
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "err": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "ok", "data": e, "dbData": result})
}
