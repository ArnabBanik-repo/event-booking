package middlewares

import (
	"net/http"

	"github.com/ArnabBanik-repo/event-booking/initializers"
	"github.com/ArnabBanik-repo/event-booking/models"
	"github.com/ArnabBanik-repo/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func Protect(c *gin.Context) {
	ck, e := c.Request.Cookie("jwt")
	if e != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "error": e.Error()})
		return
	}

	id, e := utils.VerifyToken(ck.Value)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "error": e.Error()})
		return
	}

	var u models.User
  result := initializers.DB.First(&u, "id = ?", id)
  if result.Error != nil {
    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "error": result.Error.Error()})
    return
  }

	c.Set("user", u)
	c.Next()
}
