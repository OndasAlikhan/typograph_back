package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"typograph_back/src/dto"
	"typograph_back/src/service"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	connectionType      = "CONNECTION"
	enterLobbyType      = "ENTER_LOBBY"
	leaveLobbyType      = "LEAVE_LOBBY"
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
type LeaveLobbyMsg struct {
	Type    string `json:"type"`
	UserID  uint   `json:"user_id"`
	LobbyID uint   `json:"lobby_id"`
}
type BroadcastInRoomMsg struct {
	Type    string       `json:"type"`
	LobbyID uint         `json:"lobby_id"`
	UserID  uint         `json:"user_id"`
	Text    []dto.Letter `json:"text"`
}

type LobbyWSController struct {
	*BaseController
	lobbyWsService *service.LobbyWsService
}

func NewLobbyWSController(lws *service.LobbyWsService) *LobbyWSController {
	lwc := &LobbyWSController{lobbyWsService: lws}
	// go lwc.handleBroadcast()
	return lwc
}

func (lwc LobbyWSController) Index(c echo.Context) error {
	fmt.Printf("----connecting...\n")
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return err
	}

	defer conn.Close()

	fmt.Printf("-----BEFORE FOR-----\n")
	for {
		msgType, p, err := conn.ReadMessage()
		if msgType != 1 {
			return err
		}

		var typeSwitch TypeSwitch
		err = json.Unmarshal(p, &typeSwitch)

		if err != nil {
			c.Logger().Error(err)
			conn.Close()

			return err
		}

		switch typeSwitch.Type {
		case connectionType:
			var connectionMsg ConnectionMsg
			err := json.Unmarshal(p, &connectionMsg)

			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}

			lwc.lobbyWsService.AddClient(connectionMsg.UserID, conn)
			conn.SetCloseHandler(func(code int, text string) error {
				lwc.lobbyWsService.RemoveClient(connectionMsg.UserID)
				return nil
			})

		case enterLobbyType:
			var enterLobbyMsg EnterLobbyMsg
			err := json.Unmarshal(p, &enterLobbyMsg)

			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}
			lwc.lobbyWsService.AddUserToRoom(enterLobbyMsg.LobbyID, enterLobbyMsg.UserID)

		case leaveLobbyType:
			var leaveLobbyMsg LeaveLobbyMsg
			err := json.Unmarshal(p, &leaveLobbyMsg)

			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}
			lwc.lobbyWsService.RemoveUserFromRoom(leaveLobbyMsg.LobbyID, leaveLobbyMsg.UserID)

		case broadcastInRoomType:
			var broadcastInRoomMsg BroadcastInRoomMsg
			err := json.Unmarshal(p, &broadcastInRoomMsg)

			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}
			fmt.Printf("broadcastMessage: %v\n", broadcastInRoomMsg)
			lwc.lobbyWsService.HandleNewText(broadcastInRoomMsg.LobbyID, broadcastInRoomMsg.UserID, broadcastInRoomMsg.Text)
		}

		fmt.Printf("HERE IS THE MESSAGE \n")
	}
}

// func (lwc LobbyWSController) handleBroadcast() {
// 	for {
// 		msg := <-lwc.broadcast
// 		for clientID := range lwc.clients {
// 			err := lwc.clients[clientID].WriteJSON(msg)
// 			if err != nil {
// 				fmt.Println(err)
// 				lwc.clients[clientID].Close()
// 				delete(lwc.clients, clientID)
// 			}
// 		}
// 	}
// }
