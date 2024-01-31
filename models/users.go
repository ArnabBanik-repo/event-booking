package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
  ID       string `gorm:"primaryKey"`
  Email    string `gorm:"unique" binding:"required"`
	Password string `binding:"required"`
	Events   []Event
}
