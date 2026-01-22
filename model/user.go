// Package model
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   int    `gorm:"unique;not null" json:"userid"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
