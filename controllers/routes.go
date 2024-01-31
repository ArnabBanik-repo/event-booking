package controllers

import "github.com/gin-gonic/gin"

func RegisterRoutes(s *gin.Engine) {
	s.GET("/events", GetEvents)
	s.POST("/events", CreateEvent)
	s.GET("/events/:id", GetEvent)
	s.PUT("/events/:id", UpdateEvent)
  s.DELETE("/events/:id", DeleteEvent)

  s.POST("/signup", SignUp)
  s.POST("/login", Login)
}
