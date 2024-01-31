package controllers

import (
	"net/http"

	"github.com/ArnabBanik-repo/event-booking/initializers"
	"github.com/ArnabBanik-repo/event-booking/models"
	"github.com/ArnabBanik-repo/event-booking/utils"
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

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "error": err.Error()})
		return
	}

  c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt", token, 24 * 60 * 60, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user, "token": token})
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

func GetUser(c *gin.Context) {
}
