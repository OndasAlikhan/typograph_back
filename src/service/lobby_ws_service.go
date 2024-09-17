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

type BroadcastMsg struct {
	roomId uint
	userId uint
	text   [][]dto.Letter
}

func NewLobbyWsService(repo *repository.LobbyWsRepository) *LobbyWsService {
	return &LobbyWsService{repository: repo}
}

func (lws *LobbyWsService) UpdateText(msg dto.UpdateTextMsg) {
	lws.repository.SaveUserText(msg.LobbyID, msg.UserID, msg.Text)
}

func (lws *LobbyWsService) Finish(msg dto.FinishMsg) {
	lws.repository.UserFinished(msg.LobbyID, msg.UserID)
}

func (lws *LobbyWsService) GetRoomInfo(roomId uint) dto.Room {
	return lws.repository.GetRoomInfo(roomId)
}

func (lws *LobbyWsService) StartLobby(roomId uint) {
	lws.repository.ChangeStatus(roomId, "running")
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

func (lws *LobbyWsService) Sync(rooms []repository.RoomWithId) {
	lws.repository.Sync(rooms)
}
