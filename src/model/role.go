package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string        `gorm:"not null;size:125"`
	Slug        string        `gorm:"not null"`
	Users       []*User       `gorm:"foreignKey:RoleID"`
	Permissions []*Permission `gorm:"many2many:roles_permissions;"`
}
