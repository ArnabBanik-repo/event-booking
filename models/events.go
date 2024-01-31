package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID          string    
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
  DateTime    time.Time `binding:"required"`
	UserId      string
  CreatedAt   time.Time
	UpdatedAt   time.Time
}
