package repository

import (
	"encoding/json"
	"sync"
	"typograph_back/src/dto"

	"github.com/gorilla/websocket"
)

type RoomWithId struct {
	ID uint
	dto.Room
}

type LobbyWsRepository struct {
	mut     sync.RWMutex
	clients map[uint]*websocket.Conn
	rooms   map[uint]dto.Room
}

func NewLobbyWsRepository() *LobbyWsRepository {
	return &LobbyWsRepository{
		clients: make(map[uint]*websocket.Conn),
		rooms:   make(map[uint]dto.Room),
	}
}

func (lwr *LobbyWsRepository) Sync(rooms []RoomWithId) {
	for _, room := range rooms {
		addRoom := dto.Room{
			Texts:     make(map[uint][][]dto.Letter),
			Users:     room.Users,
			UsersDone: make(map[uint]bool),
			Status:    room.Status,
		}
		lwr.rooms[room.ID] = addRoom
	}
}

func (lwr *LobbyWsRepository) SaveUserText(roomId uint, userId uint, text [][]dto.Letter) {
	lwr.mut.RLock()
	defer lwr.mut.RUnlock()

	if entry, ok := lwr.rooms[roomId]; ok {
		entry.Texts[userId] = text
		lwr.rooms[roomId] = entry
	}

	lwr.BroadcastToRoom(roomId, "update_text")
}

func (lwr *LobbyWsRepository) GetRoomInfo(roomId uint) dto.Room {
	return lwr.rooms[roomId]
}

func (lwr *LobbyWsRepository) BroadcastToRoom(roomId uint, messageType string) {
	users := lwr.rooms[roomId].Users
	connections := make([]*websocket.Conn, 0)
	for _, user := range users {
		if conn, ok := lwr.clients[user.ID]; ok {
			connections = append(connections, conn)
		}
	}

	for _, conn := range connections {
		wsMsg := dto.RoomWSMessage{
			Type: messageType,
			Data: lwr.rooms[roomId],
		}
		json, _ := json.Marshal(wsMsg)
		// stringJson := string(json)
		conn.WriteMessage(websocket.TextMessage, json)
	}
}

func (lwr *LobbyWsRepository) AddUserToRoom(roomId uint, user dto.UserResponse) {
	lwr.mut.Lock()
	defer lwr.mut.Unlock()

	lwr.createRoom(roomId)

	if entry, ok := lwr.rooms[roomId]; ok {
		entry.Users = append(entry.Users, user)
		entry.UsersDone[user.ID] = false
		lwr.rooms[roomId] = entry
	}

	lwr.BroadcastToRoom(roomId, "update_users")
}

func (lwr *LobbyWsRepository) createRoom(roomId uint) {
	if _, ok := lwr.rooms[roomId]; !ok {
		lwr.rooms[roomId] = dto.Room{
			Texts:     make(map[uint][][]dto.Letter),
			Users:     make([]dto.UserResponse, 0),
			UsersDone: make(map[uint]bool),
			Status:    "waiting",
		}
	}
}

func (lwr *LobbyWsRepository) RemoveUserFromRoom(roomId uint, userId uint) {
	lwr.mut.Lock()
	defer lwr.mut.Unlock()

	if entry, ok := lwr.rooms[roomId]; ok {
		for i, user := range entry.Users {
			if user.ID == userId {
				entry.Users = append(entry.Users[:i], entry.Users[i+1:]...)
				delete(entry.UsersDone, userId)
				delete(entry.Texts, userId)
				lwr.rooms[roomId] = entry
				break
			}
		}
	}

	lwr.BroadcastToRoom(roomId, "update_users")
}

func (lwr *LobbyWsRepository) AddClient(userId uint, conn *websocket.Conn) {
	lwr.mut.RLock()
	lwr.clients[userId] = conn
	lwr.mut.RUnlock()
}

func (lwr *LobbyWsRepository) RemoveClient(userId uint) {
	lwr.mut.RLock()
	delete(lwr.clients, userId)
	lwr.mut.RUnlock()
}

func (lwr *LobbyWsRepository) UserFinished(roomId uint, userId uint) error {
	lwr.mut.RLock()
	if room, ok := lwr.rooms[roomId]; ok {
		room.UsersDone[userId] = true
	}
	lwr.mut.RUnlock()
	lwr.BroadcastToRoom(roomId, "update_users_done")

	return nil
}

// todo ChangeStatus
func (lwr *LobbyWsRepository) ChangeStatus(roomId uint, status string) error {
	lwr.mut.RLock()
	if room, ok := lwr.rooms[roomId]; ok {
		room.Status = status
	}
	lwr.mut.RUnlock()
	lwr.BroadcastToRoom(roomId, "update_status")

	return nil
}
