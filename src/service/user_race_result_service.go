package service

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
	"typograph_back/src/repository"
)

type UserRaceResultService struct {
	repository repository.UserRaceResultRepository
}

func NewUserRaceResultService() *UserRaceResultService {
	return &UserRaceResultService{repository: *repository.NewUserRaceResulRepository()}
}

func (us *UserRaceResultService) Create(request *dto.UserRaceResultCreateRequest) (*model.UserRaceResult, error) {
	value := model.UserRaceResult{
		Duration: request.Duration,
		WPM:      request.WPM,
		Accuracy: uint8(request.Accuracy),
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
	value.Accuracy = uint8(request.Accuracy)
	value.UserID = request.UserID
	value.RaceID = request.RaceID

	return us.repository.Save(*value)
}

func (us *UserRaceResultService) Delete(id uint) error {
	return us.repository.Delete(id)
}
