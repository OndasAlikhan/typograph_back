package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string `gorm:"unique;not null;size:125"`
	Name            string `gorm:"not null;size:125"`
	Password        string `gorm:"not null"`
	RoleID          uint
	Role            *Role             `gorm:"foreignKey:RoleID"`
	IsAnon          bool              `gorm:"default:false"`
	Races           []*Race           `gorm:"many2many:race_users;"`
	Lobbies         []*Lobby          `gorm:"many2many:lobby_users;"`
	UserRaceResults []*UserRaceResult `gorm:"foreignKey:UserID"`
}
