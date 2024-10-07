package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint
	Username   string
	Password   string
	Name       string
	Contact_no string
	Email      string
	Role       string
	Hotel_ID   uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
