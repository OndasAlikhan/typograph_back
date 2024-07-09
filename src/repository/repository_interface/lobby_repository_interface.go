package repository_interface

import (
	"typograph_back/src/model"

	"gorm.io/gorm"
)

type LobbyRepositoryInterface interface {
	GetById(id uint) (*model.Lobby, *gorm.DB, error)
	GetAll() ([]*model.Lobby, error)
	UpdateUsers(users []*model.User, lobby *model.Lobby) error
	UpdateRaces(users []*model.Race, lobby *model.Lobby) error
	Save(lobby model.Lobby) (*model.Lobby, *gorm.DB, error)
	Delete(id uint) error
}
