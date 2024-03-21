package repository

import (
	"github.com/gorilla/websocket"
)

type LobbyWsRepository struct {
	clients map[uint]*websocket.Conn
}

func NewLobbyWsRepository() *LobbyWsRepository {
	return &LobbyWsRepository{clients: make(map[uint]*websocket.Conn)}
}

func (lwr *LobbyWsRepository) AddClient(userId uint, conn *websocket.Conn) {
	lwr.clients[userId] = conn
}

func (lwr *LobbyWsRepository) RemoveClient(userId uint) {
	delete(lwr.clients, userId)
}
