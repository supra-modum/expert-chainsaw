package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;column:id"`
	Name     string `gorm:"size:255;column:name"`
	Email    string `gorm:"size:255;unique;column:email"`
	Password string `gorm:"size:255;column:password"`
}
