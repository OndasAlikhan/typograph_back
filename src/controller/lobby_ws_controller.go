package controller

import (
	"fmt"
	"typograph_back/src/dto"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

const (
	connectionMsg      = "CONNECTION"
	enterLobbyMsg      = "ENTER_LOBBY"
	broadcastInRoomMsg = "BROADCAST_IN_ROOM"
)

type LobbyWSController struct {
	*BaseController
	clients   map[uint]*websocket.Conn
	lobbies   map[uint][]*websocket.Conn
	broadcast chan interface{}
}

type SocketMessage struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}
type ConnectionMsgBody struct {
	UserID uint `json:"user_id"`
}
type EnterLobbyMsgBody struct {
	UserID  uint `json:"user_id"`
	LobbyID uint `json:"lobby_id"`
}
type BroadcastInRoomMsgBody struct {
	LobbyID uint   `json:"lobby_id"`
	UserID  uint   `json:"user_id"`
	Text    string `json:"text"`
}

func NewLobbyWSController() *LobbyWSController {
	lwc := &LobbyWSController{}
	lwc.handleBroadcast()
	return lwc
}

func (lwc LobbyWSController) Index(c echo.Context) error {
	request := dto.LobbyWSCreateRequest{}
	if err := lwc.handleRequest(&request, c); err != nil {
		return err
	}

	// ping pong

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	lwc.clients[request.UserID] = conn

	for {
		// Write
		// err := conn.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		// if err != nil {
		// 	c.Logger().Error(err)
		// }

		// Read
		var message SocketMessage
		err := conn.ReadJSON(&message)

		// switch message.Type {
		// case connectRoom:
		// 	dataConnect := &DataConnectRoom{}

		// 	// dataConnect = (DataConnectRoom)message.Data
		// }

		//

		if err != nil {
			c.Logger().Error(err)
			delete(lwc.clients, request.UserID)
			conn.Close()

			return err
		}

		// Находим юзеров по лоббиИД
		// рассылаем этим юзерам сообщение

		switch message.Type {
		case connectionMsg:
			data, ok := message.Body.(ConnectionMsgBody)
			if !ok {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}
			fmt.Printf("connection %v", data)
		case enterLobbyMsg:
			data, ok := message.Body.(EnterLobbyMsgBody)
			if !ok {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}
			fmt.Printf("enter_lobby %v", data)
		case broadcastInRoomMsg:
			data, ok := message.Body.(BroadcastInRoomMsgBody)
			if !ok {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}
			fmt.Printf("broadcast %v", data)
		}

		fmt.Printf("HERE IS THE MESSAGE %s\n", message)
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
