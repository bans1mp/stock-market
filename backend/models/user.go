package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

type UserBalance struct {
	gorm.Model 
	UserID     uint    `gorm:"not null;index"` 
	Balance    float64 `gorm:"not null;default:0"`
	ReservedBalance float64 `gorm:"not null;default:0"`
}