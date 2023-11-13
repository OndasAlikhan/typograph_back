package model

import "gorm.io/gorm"

type Paragraph struct {
	gorm.Model
	Source string `gorm:"not null"`
	Text   string `gorm:"not null"`
	Length uint   `gorm:"not null"`
}
