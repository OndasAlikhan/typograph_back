package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name  string  `gorm:"not null;size:125"`
	Slug  string  `gorm:"not null"`
	Roles []*Role `gorm:"many2many:roles_permissions;"`
}
