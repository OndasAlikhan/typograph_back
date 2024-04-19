package service

import (
	"typograph_back/src/dto"
	"typograph_back/src/repository"

	"github.com/gorilla/websocket"
)

type LobbyWsService struct {
	repository *repository.LobbyWsRepository
}

func NewLobbyWsService() *LobbyWsService {
	return &LobbyWsService{repository: repository.NewLobbyWsRepository()}
}

func (lws *LobbyWsService) HandleNewText(roomId uint, userId uint, text []dto.Letter) {
	lws.repository.SaveUserText(roomId, userId, text)
	lws.repository.BroadcastToRoom(roomId)
}

func (lws *LobbyWsService) AddUserToRoom(roomId uint, userId uint) {
	lws.repository.AddUserToRoom(roomId, userId)
}

func (lws *LobbyWsService) RemoveUserFromRoom(roomId uint, userId uint) {
	lws.repository.RemoveUserFromRoom(roomId, userId)
}

func (lws *LobbyWsService) AddClient(userId uint, conn *websocket.Conn) {
	lws.repository.AddClient(userId, conn)
}

func (lws *LobbyWsService) RemoveClient(userId uint) {
	lws.repository.RemoveClient(userId)
}
