package service

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
	"typograph_back/src/repository"
	"typograph_back/src/service/service_interface"
)

var _ service_interface.UserRaceResultServiceInterface = (*UserRaceResultService)(nil)

type UserRaceResultService struct {
	repository repository.UserRaceResultRepository
}

func NewUserRaceResultService() *UserRaceResultService {
	return &UserRaceResultService{repository: *repository.NewUserRaceResulRepository()}
}

func (us *UserRaceResultService) GetById(id uint) (*model.UserRaceResult, error) {
	return us.repository.GetById(id)
}

func (us *UserRaceResultService) GetByUserId(userId uint) ([]*model.UserRaceResult, error) {
	return us.repository.GetByUserId(userId)
}

func (us *UserRaceResultService) Leaderboard() ([]*dto.LeaderboardResponse, error) {
	return us.repository.Leaderboard()
}

func (us *UserRaceResultService) Create(request *dto.UserRaceResultCreateRequest) (*model.UserRaceResult, error) {
	value := model.UserRaceResult{
		Duration: request.Duration,
		WPM:      request.WPM,
		Accuracy: request.Accuracy,
		UserID:   request.UserID,
		RaceID:   request.RaceID,
	}

	return us.repository.Save(value)
}

func (us *UserRaceResultService) Update(id uint, request *dto.UserRaceResultUpdateRequest) (*model.UserRaceResult, error) {
	value, err := us.repository.GetById(id)

	if err != nil {
		return nil, err
	}

	value.Duration = request.Duration
	value.WPM = request.WPM
	value.Accuracy = request.Accuracy
	value.UserID = request.UserID
	value.RaceID = request.RaceID

	return us.repository.Save(*value)
}

func (us *UserRaceResultService) Delete(id uint) error {
	return us.repository.Delete(id)
}
