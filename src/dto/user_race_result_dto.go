package dto

import "typograph_back/src/model"

type UserRaceResultCreateRequest struct {
	Duration uint `json:"duration" validate:"required"`
	WPM      uint `json:"wpm" validate:"required"`
	Accuracy uint `json:"accuracy" validate:"required,max=100"`
	UserID   uint `json:"user_id" validate:"required"`
	RaceID   uint `json:"race_id" validate:"required"`
}

type UserRaceResultUpdateRequest struct {
	ID       uint `json:"id" validate:"required"`
	Duration uint `json:"duration" validate:"required"`
	WPM      uint `json:"wpm" validate:"required"`
	Accuracy uint `json:"accuracy" validate:"required,max=100"`
	UserID   uint `json:"user_id" validate:"required"`
	RaceID   uint `json:"race_id" validate:"required"`
}

type UserRaceResultResponse struct {
	ID       uint  `json:"id"`
	Duration uint  `json:"duration"`
	WPM      uint  `json:"wpm"`
	Accuracy uint8 `json:"accuracy"`
	UserID   uint  `json:"user_id"`
	RaceID   uint  `json:"race_id"`
}

func NewUserRaceResultResponse(userRaceResult *model.UserRaceResult) *UserRaceResultResponse {
	return &UserRaceResultResponse{
		ID:       userRaceResult.ID,
		Duration: userRaceResult.Duration,
		WPM:      userRaceResult.WPM,
		Accuracy: userRaceResult.Accuracy,
		UserID:   userRaceResult.UserID,
		RaceID:   userRaceResult.RaceID,
	}
}
