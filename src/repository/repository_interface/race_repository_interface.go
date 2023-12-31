package repository_interface

import (
	"typograph_back/src/model"

	"gorm.io/gorm"
)

type RaceRepositoryInterface interface {
	GetById(id uint) (*model.Race, error)
	GetByIds(id []uint) ([]*model.Race, error)
	GetAll() ([]*model.Race, error)
	Save(race model.Race) (*model.Race, *gorm.DB, error)
	UpdateUsers(users []*model.User, tx *gorm.DB) error
	Delete(id uint) error
}
