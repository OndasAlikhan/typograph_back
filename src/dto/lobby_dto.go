package dto

import "typograph_back/src/model"

type LobbyCreateRequest struct {
	AdminUserID uint   `json:"admin_user_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Users       []uint `json:"users" validate:"required"`
	Races       []uint `json:"races"`
}
type LobbyUpdateRequest struct {
	ID          uint   `json:"id" validate:"required"`
	AdminUserID uint   `json:"admin_user_id" validate:"required"`
	Status      string `json:"status"`
	Name        string `json:"name" validate:"required"`
	Users       []uint `json:"users" validate:"required"`
	Races       []uint `json:"races"`
}

type LobbyResponse struct {
	ID          uint            `json:"id"`
	AdminUserID uint            `json:"admin_user_id"`
	Status      string          `json:"status"`
	Name        string          `json:"name"`
	Users       []*UserResponse `json:"users"`
	Races       []*RaceResponse `json:"races"`
}

func NewLobbyResponse(lobby *model.Lobby) *LobbyResponse {
	var usersResponse []*UserResponse
	for _, user := range lobby.Users {
		usersResponse = append(usersResponse, NewUserResponse(user))
	}

	var racesResponse []*RaceResponse
	for _, race := range lobby.Races {
		racesResponse = append(racesResponse, NewRaceResponse(race))
	}

	return &LobbyResponse{
		ID:          lobby.ID,
		AdminUserID: lobby.AdminUserID,
		Status:      lobby.Status,
		Name:        lobby.Name,
		Users:       usersResponse,
		Races:       racesResponse,
	}
}
