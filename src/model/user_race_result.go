package model

import "gorm.io/gorm"

type UserRaceResult struct {
	gorm.Model
	Duration uint  `gorm:"default:0"` //seconds
	WPM      uint  `gorm:"default:0"`
	Accuracy uint8 `gorm:"default:0"` //percents
	UserID   uint
	User     *User `gorm:"foreign:UserID"`
	RaceID   uint
	Race     *Race `gorm:"foreign:RaceID"`
}
