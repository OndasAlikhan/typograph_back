package dto

import (
	"typograph_back/src/model"
)

type RaceCreateRequest struct {
	Finished    bool `json:"finished" validate:"required"`
	AdminUserID uint `json:"admin_user_id" validate:"required"`
	ParagraphID uint `json:"paragraph_id" validate:"required"`
}

type RaceUpdateRequest struct {
	Finished    bool   `json:"finished" validate:"required"`
	AdminUserID uint   `json:"admin_user_id" validate:"required"`
	Users       []uint `json:"users"`
	ParagraphID uint   `json:"paragraph_id" validate:"required"`
}

type RaceResponse struct {
	ID          uint            `json:"id"`
	AdminUserID uint            `json:"admin_user_id"`
	Users       []*UserResponse `json:"users"`
	ParagraphID uint            `json:"paragraph_id"`
}

func NewRaceResponse(race *model.Race) *RaceResponse {
	var usersResponse []*UserResponse
	for _, user := range race.Users {
		usersResponse = append(usersResponse, NewUserResponse(user))
	}

	return &RaceResponse{
		ID:          race.ID,
		AdminUserID: race.AdminUserID,
		Users:       usersResponse,
		ParagraphID: race.ParagraphID,
	}
}
