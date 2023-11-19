package repository_interface

import (
	"typograph_back/src/model"
)

type UserRaceResultRepositoryInterface interface {
	GetById(id uint) (*model.UserRaceResult, error)
	GetAll() ([]*model.UserRaceResult, error)
	Save(value model.UserRaceResult) (*model.UserRaceResult, error)
	Delete(id uint) error
}
