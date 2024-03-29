package service_interface

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
)

type UserRaceResultServiceInterface interface {
	GetById(id uint) (*model.UserRaceResult, error)
	GetByUserId(userId uint) ([]*model.UserRaceResult, error)
	Leaderboard() ([]*dto.LeaderboardResponse, error)
	Create(request *dto.UserRaceResultCreateRequest) (*model.UserRaceResult, error)
	Update(id uint, request *dto.UserRaceResultUpdateRequest) (*model.UserRaceResult, error)
	Delete(id uint) error
}
