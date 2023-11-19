package model

import "gorm.io/gorm"

type Race struct {
	gorm.Model
	Finished    bool `gorm:"default:false"`
	AdminUserID uint
	AdminUser   *User   `gorm:"foreignKey:AdminUserID"`
	Users       []*User `gorm:"many2many:race_users;"`
	ParagraphID uint
	Paragraph   *Paragraph `gorm:"foreignKey:ParagraphID"`
}
