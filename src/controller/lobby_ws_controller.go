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
	connectionType = "CONNECTION"
	enterLobbyType = "ENTER_LOBBY"
	leaveLobbyType = "LEAVE_LOBBY"
	updateTextType = "UPDATE_TEXT"
	finishType     = "FINISH"
)

type TypeSwitch struct {
	Type string `json:"type"`
}

type ConnectionMsg struct {
	Type   string `json:"type"`
	UserID uint   `json:"user_id"`
}

//	type EnterLobbyMsg struct {
//		Type    string `json:"type"`
//		UserID  uint   `json:"user_id"`
//		LobbyID uint   `json:"lobby_id"`
//	}
//
//	type LeaveLobbyMsg struct {
//		Type    string `json:"type"`
//		UserID  uint   `json:"user_id"`
//		LobbyID uint   `json:"lobby_id"`
//	}

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
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return err
	}

	defer conn.Close()

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

			// handle if client disconnects
			conn.SetCloseHandler(func(code int, text string) error {
				lwc.lobbyWsService.RemoveClient(connectionMsg.UserID)
				return nil
			})

		case updateTextType:
			var updateTextMsg dto.UpdateTextMsg
			err := json.Unmarshal(p, &updateTextMsg)

			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}
			lwc.lobbyWsService.UpdateText(updateTextMsg)

		case finishType:
			var finishMsg dto.FinishMsg
			err := json.Unmarshal(p, &finishMsg)

			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Bad request"))
			}
			fmt.Printf("finishMsg: %v\n", finishMsg)
			lwc.lobbyWsService.Finish(finishMsg)
		}
	}
}
