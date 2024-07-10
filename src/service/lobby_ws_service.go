package service

import (
	"fmt"
	"typograph_back/src/dto"
	"typograph_back/src/model"
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
}

func (lws *LobbyWsService) AddUserToRoom(roomId uint, user *model.User) {
	fmt.Printf("Service AddUserToRoom roomId:%d user:%v\n", roomId, user)
	userDto := dto.NewUserResponse(user)
	lws.repository.AddUserToRoom(roomId, *userDto)
}

func (lws *LobbyWsService) RemoveUserFromRoom(roomId uint, userId uint) {
	fmt.Printf("Service RemoveUserFromRoom roomId:%d userId:%d\n", roomId, userId)
	lws.repository.RemoveUserFromRoom(roomId, userId)
}

func (lws *LobbyWsService) AddClient(userId uint, conn *websocket.Conn) {
	lws.repository.AddClient(userId, conn)
}

func (lws *LobbyWsService) RemoveClient(userId uint) {
	lws.repository.RemoveClient(userId)
}
