package model

import "gorm.io/gorm"

type Lobby struct {
	gorm.Model
	AdminUserID uint
	Status      string  // waiting starting running finished
	Name        string  `gorm:"not null;"`
	AdminUser   *User   `gorm:"foreignKey:AdminUserID"`
	Users       []*User `gorm:"many2many:lobby_users;"`
	Races       []*Race `gorm:"foreignKey:LobbyID"`
}
