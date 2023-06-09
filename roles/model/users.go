package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
}

type User struct {
	gorm.Model
	FullName string
	Name     string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Roles    []Role `gorm:"many2many:user_roles;"`
}

type Resource struct {
	gorm.Model
	Path  string
	Roles []Role `gorm:"many2many:resource_roles;"`
}
