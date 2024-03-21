package service

import (
	"typograph_back/src/repository"

	"github.com/gorilla/websocket"
)

type LobbyWsService struct {
	repository *repository.LobbyWsRepository
}

func NewLobbyWsService() *LobbyWsService {
	return &LobbyWsService{repository: repository.NewLobbyWsRepository()}
}

func (lws *LobbyWsService) AddClient(userId uint, conn *websocket.Conn) {
	lws.repository.AddClient(userId, conn)
}

func (lws *LobbyWsService) RemoveClient(userId uint) {
	lws.repository.RemoveClient(userId)
}
