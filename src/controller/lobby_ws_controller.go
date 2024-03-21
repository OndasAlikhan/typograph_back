package controller

import (
	"encoding/json"
	"fmt"

	"typograph_back/src/service"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

const (
	connectionType      = "CONNECTION"
	enterLobbyType      = "ENTER_LOBBY"
	broadcastInRoomType = "BROADCAST_IN_ROOM"
)

type TypeSwitch struct {
	Type string `json:"type"`
}

type ConnectionMsg struct {
	Type   string `json:"type"`
	UserID uint   `json:"user_id"`
}
type EnterLobbyMsg struct {
	Type    string `json:"type"`
	UserID  uint   `json:"user_id"`
	LobbyID uint   `json:"lobby_id"`
}
type BroadcastInRoomMsg struct {
	Type    string `json:"type"`
	LobbyID uint   `json:"lobby_id"`
	UserID  uint   `json:"user_id"`
	Text    string `json:"text"`
}

type LobbyWSController struct {
	*BaseController
	clients        map[uint]*websocket.Conn
	lobbies        map[uint][]*websocket.Conn
	broadcast      chan interface{}
	lobbyWsService *service.LobbyWsService
}

func NewLobbyWSController() *LobbyWSController {
	lwc := &LobbyWSController{lobbyWsService: service.NewLobbyWsService()}
	// go lwc.handleBroadcast()
	return lwc
}

func (lwc LobbyWSController) Index(c echo.Context) error {
	fmt.Println("Index 1")
	// request := dto.LobbyWSCreateRequest{}
	// if err := lwc.handleRequest(&request, c); err != nil {
	// 	return err
	// }

	// ping pong

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	fmt.Println("Index 2")

	defer conn.Close()

	// lwc.clients[request.UserID] = conn

	for {
		// Write
		// err := conn.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		// if err != nil {
		// 	c.Logger().Error(err)
		// }
		// Read
		msgType, p, err := conn.ReadMessage()
		if msgType != 1 {
			return err
		}

		var typeSwitch TypeSwitch
		err = json.Unmarshal(p, &typeSwitch)

		if err != nil {
			c.Logger().Error(err)
			// delete(lwc.clients, request.UserID)
			conn.Close()

			return err
		}

		// Находим юзеров по лоббиИД
		// рассылаем этим юзерам сообщение

		switch typeSwitch.Type {
		case connectionType:
			var connectionMsg ConnectionMsg
			err := json.Unmarshal(p, &connectionMsg)

			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}

			lwc.lobbyWsService.AddClient(connectionMsg.UserID, conn)
		case enterLobbyType:
			var enterLobbyMsg EnterLobbyMsg
			err := json.Unmarshal(p, &enterLobbyMsg)

			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}
			fmt.Printf("connection!!! %v\n", enterLobbyMsg)
		case broadcastInRoomType:
			var broadcastInRoomMsg BroadcastInRoomMsg
			err := json.Unmarshal(p, &broadcastInRoomMsg)

			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}
			fmt.Printf("connection!!! %v\n", broadcastInRoomMsg)
		}

		fmt.Printf("HERE IS THE MESSAGE %s\n")
	}
}

func (lwc LobbyWSController) handleBroadcast() {
	for {
		msg := <-lwc.broadcast
		for clientID := range lwc.clients {
			err := lwc.clients[clientID].WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				lwc.clients[clientID].Close()
				delete(lwc.clients, clientID)
			}
		}
	}
}
