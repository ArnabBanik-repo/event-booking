package controllers

import (
	"net/http"

	"github.com/ArnabBanik-repo/event-booking/initializers"
	"github.com/ArnabBanik-repo/event-booking/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var u models.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "email = ?", u.Email)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": "Invalid email or password"})
		return
	}

  c.JSON(http.StatusOK, gin.H{"status": "success", "data": user, "token": "jwt"})
}

func SignUp(c *gin.Context) {
	var u models.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	h, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	i, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": err.Error()})
		return
	}
	u.ID = i.String()
	u.Password = string(h)

	result := initializers.DB.Create(&u)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": u}})
}
