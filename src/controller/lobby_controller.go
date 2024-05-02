package controller

import (
	"net/http"
	"typograph_back/src/dto"
	"typograph_back/src/service/service_interface"

	"github.com/labstack/echo/v4"
)

type LobbyController struct {
	*BaseController
	service service_interface.LobbyServiceInterface
}

func NewLobbyController(service service_interface.LobbyServiceInterface) *LobbyController {
	return &LobbyController{
		BaseController: NewBaseController(),
		service:        service,
	}
}

// Index
// @title Index
// @description List of lobbies
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags lobby
// @success 200 {object} dto.JSONResult{data=[]dto.LobbyResponse}
// @router /lobbies [get]
func (rc *LobbyController) Index(c echo.Context) error {
	lobbies, err := rc.service.GetAll()
	if err != nil {
		return err
	}

	response := make([]*dto.LobbyResponse, len(lobbies))
	for ind, val := range lobbies {
		response[ind] = dto.NewLobbyResponse(val)
	}

	return rc.json(http.StatusOK, response, c)
}

// Show
// @title Show
// @description Get a lobby by id
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags lobby
// @success 200 {object} dto.JSONResult{data=dto.LobbyResponse}
// @router /lobbies/{id} [get]
func (rc *LobbyController) Show(c echo.Context) error {
	id, err := rc.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}
	lobby, err := rc.service.GetById(id)
	if err != nil {
		return err
	}
	return rc.json(http.StatusOK, dto.NewLobbyResponse(lobby), c)
}

// Store
// @title Create
// @description Create a new lobby
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags lobby
// @param lobbyStoreRequest body dto.LobbyCreateRequest true "Lobby store request"
// @success 200 {object} dto.JSONResult{data=dto.LobbyResponse}
// @router /lobbies [post]
func (rc *LobbyController) Store(c echo.Context) error {
	request := dto.LobbyCreateRequest{}
	if err := rc.handleRequest(&request, c); err != nil {
		return err
	}

	lobby, err := rc.service.Create(&request)
	if err != nil {
		return err
	}
	return rc.json(http.StatusOK, dto.NewLobbyResponse(lobby), c)
}

// EnterLobby
// @title EnterLobby
// @description Enter a lobby
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags lobby
// @param enterLobbyRequest body dto.EnterLobbyRequest true "Enter lobby request"
// @success 200 {object} dto.JSONResult{data=nil}
// @router /lobbies/enter [post]
func (rc *LobbyController) EnterLobby(c echo.Context) error {
	request := dto.EnterLobbyRequest{}
	if err := rc.handleRequest(&request, c); err != nil {
		return err
	}

	err := rc.service.EnterLobby(request.LobbyID, request.UserID)
	if err != nil {
		return err
	}
	return rc.json(http.StatusOK, nil, c)
}

// LeaveLobby
// @title LeaveLobby
// @description Leave a lobby
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags lobby
// @param leaveLobbyRequest body dto.LeaveLobbyRequest true "Leave lobby request"
// @success 200 {object} dto.JSONResult{data=nil}
// @router /lobbies/leave [post]
func (rc *LobbyController) LeaveLobby(c echo.Context) error {
	request := dto.LeaveLobbyRequest{}
	if err := rc.handleRequest(&request, c); err != nil {
		return err
	}

	err := rc.service.LeaveLobby(request.LobbyID, request.UserID)
	if err != nil {
		return err
	}
	return rc.json(http.StatusOK, nil, c)
}

// Update
// @title Update
// @description Update a lobby
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags lobby
// @param lobbyStoreRequest body dto.LobbyUpdateRequest true "Lobby update request"
// @param id path int true "Lobby ID"
// @success 200 {object} dto.JSONResult{data=dto.LobbyResponse}
// @router /lobbies/{id} [patch]
func (rc *LobbyController) Update(c echo.Context) error {
	id, err := rc.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	request := dto.LobbyUpdateRequest{}
	if err := rc.handleRequest(&request, c); err != nil {
		return err
	}

	lobby, err := rc.service.Update(id, &request)
	if err != nil {
		return err
	}
	return rc.json(http.StatusOK, dto.NewLobbyResponse(lobby), c)
}

// Delete
// @title Delete
// @description Delete a lobby
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags lobby
// @param id path int true "Lobby ID"
// @success 200 {object} dto.JSONResult
// @router /lobbies/{id} [delete]
func (rc *LobbyController) Delete(c echo.Context) error {
	id, err := rc.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	if err = rc.service.Delete(id); err != nil {
		return err
	}

	return rc.json(http.StatusOK, nil, c)
}
