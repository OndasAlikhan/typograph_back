package repository

import (
	"fmt"
	"sync"
	"typograph_back/src/dto"

	"github.com/gorilla/websocket"
)

type Room struct {
	texts      map[uint][]dto.Letter
	users      []uint
	users_done map[uint]bool
	status     string // waiting starting running finished
}

type LobbyWsRepository struct {
	mut     sync.RWMutex
	clients map[uint]*websocket.Conn
	rooms   map[uint]Room
}

func NewLobbyWsRepository() *LobbyWsRepository {
	return &LobbyWsRepository{
		clients: make(map[uint]*websocket.Conn),
		rooms:   make(map[uint]Room),
	}
}

func (lwr *LobbyWsRepository) SaveUserText(roomId uint, userId uint, text []dto.Letter) {
	lwr.mut.RLock()
	defer lwr.mut.RUnlock()

	if entry, ok := lwr.rooms[roomId]; ok {
		entry.texts[userId] = text
		lwr.rooms[roomId] = entry
	}

	lwr.BroadcastToRoom(roomId)
}

func (lwr *LobbyWsRepository) GetUserText(roomId uint, userId uint) []dto.Letter {
	return lwr.rooms[roomId].texts[userId]
}

func (lwr *LobbyWsRepository) BroadcastToRoom(roomId uint) {
	userIds := lwr.rooms[roomId].users
	connections := make([]*websocket.Conn, 0)
	for _, userId := range userIds {
		if conn, ok := lwr.clients[userId]; ok {
			connections = append(connections, conn)
		}
	}

	for _, conn := range connections {
		conn.WriteJSON(lwr.rooms[roomId])
	}
}

func (lwr *LobbyWsRepository) AddUserToRoom(roomId uint, userId uint) {
	lwr.mut.Lock()
	defer lwr.mut.Unlock()

	lwr.createRoom(roomId)

	if entry, ok := lwr.rooms[roomId]; ok {
		entry.users = append(entry.users, userId)
		entry.users_done[userId] = false
		lwr.rooms[roomId] = entry
	}

	lwr.BroadcastToRoom(roomId)

	fmt.Printf("Added user %d to room %d\n", userId, roomId)
	fmt.Printf("Rooms: %v\n", lwr.rooms)
}

func (lwr *LobbyWsRepository) createRoom(roomId uint) {
	if _, ok := lwr.rooms[roomId]; !ok {
		lwr.rooms[roomId] = Room{
			texts:      make(map[uint][]dto.Letter),
			users:      make([]uint, 0),
			users_done: make(map[uint]bool),
			status:     "waiting",
		}
	}
}

func (lwr *LobbyWsRepository) RemoveUserFromRoom(roomId uint, userId uint) {
	lwr.mut.Lock()
	defer lwr.mut.Unlock()

	if entry, ok := lwr.rooms[roomId]; ok {
		for i, user := range entry.users {
			if user == userId {
				entry.users = append(entry.users[:i], entry.users[i+1:]...)
				delete(entry.users_done, userId)
				delete(entry.texts, userId)
				lwr.rooms[roomId] = entry
				break
			}
		}
	}

	lwr.BroadcastToRoom(roomId)
}

func (lwr *LobbyWsRepository) AddClient(userId uint, conn *websocket.Conn) {
	lwr.mut.RLock()
	lwr.clients[userId] = conn
	lwr.mut.RUnlock()

	fmt.Printf("Added client %d\n", userId)
	fmt.Printf("Clients:  %v\n", lwr.clients)
}

func (lwr *LobbyWsRepository) RemoveClient(userId uint) {
	lwr.mut.RLock()
	delete(lwr.clients, userId)
	lwr.mut.RUnlock()
	fmt.Printf("Removed client %d\n", userId)
	fmt.Printf("Clients:  %v\n", lwr.clients)
}

func (lwr *LobbyWsRepository) UserFinished(roomId uint, userId uint) error {
	lwr.mut.RLock()
	if room, ok := lwr.rooms[roomId]; ok {
		room.users_done[userId] = true
	}
	lwr.mut.RUnlock()
	lwr.BroadcastToRoom(roomId)

	return nil
}

// todo ChangeStatus
