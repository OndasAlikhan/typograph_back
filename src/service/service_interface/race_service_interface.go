package service_interface

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
)

type RaceServiceInterface interface {
	GetAll() ([]*model.Race, error)
	GetById(id uint) (*model.Race, error)
	Create(request *dto.RaceCreateRequest) (*model.Race, error)
	Update(id uint, request *dto.RaceUpdateRequest) (*model.Race, error)
	Delete(id uint) error
}
