package service

import (
	"fmt"
	"typograph_back/src/dto"
	"typograph_back/src/repository"

	"github.com/gorilla/websocket"
)

type LobbyWsService struct {
	repository *repository.LobbyWsRepository
}

func NewLobbyWsService(repo *repository.LobbyWsRepository) *LobbyWsService {
	return &LobbyWsService{repository: repo}
}

func (lws *LobbyWsService) HandleNewText(roomId uint, userId uint, text []dto.Letter) {
	lws.repository.SaveUserText(roomId, userId, text)
	lws.repository.BroadcastToRoom(roomId)
}

func (lws *LobbyWsService) AddUserToRoom(roomId uint, userId uint) {
	fmt.Printf("AddUserToRoom roomId:%d userId:%d\n", roomId, userId)
	lws.repository.AddUserToRoom(roomId, userId)
}

func (lws *LobbyWsService) RemoveUserFromRoom(roomId uint, userId uint) {
	fmt.Printf("RemoveUserFromRoom roomId:%d userId:%d\n", roomId, userId)
	lws.repository.RemoveUserFromRoom(roomId, userId)
}

func (lws *LobbyWsService) AddClient(userId uint, conn *websocket.Conn) {
	fmt.Printf("AddClient userId:%d conn:%s\n", userId, conn)
	lws.repository.AddClient(userId, conn)
}

func (lws *LobbyWsService) RemoveClient(userId uint) {
	fmt.Printf("RemoveClient userId:%d\n", userId)
	lws.repository.RemoveClient(userId)
}
