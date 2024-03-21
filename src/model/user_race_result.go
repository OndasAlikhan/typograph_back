package model

import "gorm.io/gorm"

type UserRaceResult struct {
	gorm.Model
	Duration float32 `gorm:"default:0"` //seconds
	WPM      float32 `gorm:"default:0"`
	Accuracy float32 `gorm:"default:0"` //percents
	UserID   uint
	User     *User `gorm:"foreign:UserID"`
	RaceID   uint  `gorm:"default:null"`
	Race     *Race `gorm:"foreign:RaceID"`
}
