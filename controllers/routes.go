package controllers

import (
	"github.com/ArnabBanik-repo/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(s *gin.Engine) {
	s.GET("/events", GetEvents)
	s.GET("/events/:id", GetEvent)

  auth := s.Group("/")
  auth.Use(middlewares.Protect)
	auth.POST("/events", CreateEvent)
	auth.PUT("/events/:id", UpdateEvent)
	auth.DELETE("/events/:id", DeleteEvent)

	s.POST("/signup", SignUp)
	s.POST("/login", Login)
}
