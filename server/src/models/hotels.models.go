package models

import (
	"time"

	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	ID          uint
	Name        string
	Address     string
	Phone       string
	Email       string
	Website     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        []User
}
