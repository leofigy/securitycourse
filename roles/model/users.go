package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string
	Description string
}

type User struct {
	gorm.Model
	FullName string
	Name     string
	Email    string
	Password string
	Roles    []Role `gorm:"many2many:user_roles;"`
}

type Resource struct {
	gorm.Model
	Path  string
	Roles []Role `gorm:"many2many:resource_roles;"`
}
