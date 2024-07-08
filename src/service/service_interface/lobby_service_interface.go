package service_interface

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
)

type LobbyServiceInterface interface {
	GetAll() ([]*model.Lobby, error)
	GetById(id uint) (*model.Lobby, error)
	Create(request *dto.LobbyCreateRequest) (*model.Lobby, error)
	EnterLobby(lobbyId uint, userId uint) error
	LeaveLobby(lobbyId uint, userId uint) error
	Update(id uint, request *dto.LobbyUpdateRequest) (*model.Lobby, error)
	Delete(id uint) error
	StartLobby(id uint) error
	UserFinished(request *dto.UserFinishedRequest) error
}
