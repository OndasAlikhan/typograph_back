package repository_interface

import (
	"typograph_back/src/model"

	"gorm.io/gorm"
)

type LobbyRepositoryInterface interface {
	GetById(id uint) (*model.Lobby, *gorm.DB, error)
	GetAll() ([]*model.Lobby, error)
	UpdateUsers(users []*model.User, tx *gorm.DB) error
	UpdateRaces(users []*model.Race, tx *gorm.DB) error
	Save(lobby model.Lobby) (*model.Lobby, *gorm.DB, error)
	Delete(id uint) error
}
