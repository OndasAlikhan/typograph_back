package dto

import "typograph_back/src/model"

type UserRaceResultCreateRequest struct {
	Duration float32 `json:"duration" validate:"required"`
	WPM      float32 `json:"wpm" validate:"required"`
	Accuracy float32 `json:"accuracy" validate:"required,max=100"`
	UserID   uint    `json:"user_id" validate:"required"`
	RaceID   uint    `json:"race_id"`
}

type UserRaceResultUpdateRequest struct {
	ID       uint    `json:"id" validate:"required"`
	Duration float32 `json:"duration" validate:"required"`
	WPM      float32 `json:"wpm" validate:"required"`
	Accuracy float32 `json:"accuracy" validate:"required,max=100"`
	UserID   uint    `json:"user_id" validate:"required"`
	RaceID   uint    `json:"race_id"`
}

type UserRaceResultResponse struct {
	ID        uint    `json:"id"`
	Duration  float32 `json:"duration"`
	WPM       float32 `json:"wpm"`
	Accuracy  float32 `json:"accuracy"`
	UserID    uint    `json:"user_id"`
	RaceID    uint    `json:"race_id"`
	CreatedAt string  `json:"created_at"`
}

type LeaderboardResponse struct {
	ID        uint    `json:"id"`
	Duration  float32 `json:"duration"`
	WPM       float32 `json:"wpm"`
	Accuracy  float32 `json:"accuracy"`
	UserName  string  `json:"user_name"`
	RaceID    uint    `json:"race_id"`
	CreatedAt string  `json:"created_at"`
}

func NewUserRaceResultResponse(userRaceResult *model.UserRaceResult) *UserRaceResultResponse {
	return &UserRaceResultResponse{
		ID:        userRaceResult.ID,
		Duration:  userRaceResult.Duration,
		WPM:       userRaceResult.WPM,
		Accuracy:  userRaceResult.Accuracy,
		UserID:    userRaceResult.UserID,
		RaceID:    userRaceResult.RaceID,
		CreatedAt: userRaceResult.CreatedAt.Format("2006-01-02 15:04"),
	}
}

func NewLeaderboardResponse(userRaceResult *LeaderboardResponse) *LeaderboardResponse {
	return &LeaderboardResponse{
		ID:        userRaceResult.ID,
		Duration:  userRaceResult.Duration,
		WPM:       userRaceResult.WPM,
		Accuracy:  userRaceResult.Accuracy,
		UserName:  userRaceResult.UserName,
		RaceID:    userRaceResult.RaceID,
		CreatedAt: userRaceResult.CreatedAt,
	}
}
