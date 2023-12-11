package service

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
	"typograph_back/src/repository"
	"typograph_back/src/repository/repository_interface"
	"typograph_back/src/service/service_interface"
)

var _ service_interface.RaceServiceInterface = (*RaceService)(nil)

type RaceService struct {
	repository  repository_interface.RaceRepositoryInterface
	userService service_interface.UserServiceInterface
}

func NewRaceService() *RaceService {
	return &RaceService{repository: repository.NewRaceRepository(), userService: NewUserService()}
}

func (rs *RaceService) GetAll() ([]*model.Race, error) {
	return rs.repository.GetAll()
}

func (rs *RaceService) GetById(id uint) (*model.Race, error) {
	return rs.repository.GetById(id)
}

func (rs *RaceService) GetByIds(ids []uint) ([]*model.Race, error) {
	return rs.repository.GetByIds(ids)
}

func (rs *RaceService) Create(request *dto.RaceCreateRequest) (*model.Race, error) {
	users, err := rs.userService.GetByIds(request.Users)
	if err != nil {
		return nil, err
	}

	race := model.Race{
		Finished:    request.Finished,
		AdminUserID: request.AdminUserID,
		ParagraphID: request.ParagraphID,
		LobbyID:     request.LobbyID,
		Users:       users,
	}
	result, _, err := rs.repository.Save(race)

	return result, err
}

func (rs *RaceService) Update(id uint, request *dto.RaceUpdateRequest) (*model.Race, error) {
	race, err := rs.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	race.Finished = request.Finished
	race.AdminUserID = request.AdminUserID
	race.ParagraphID = request.ParagraphID
	race.LobbyID = request.LobbyID

	result, tx, err := rs.repository.Save(*race)
	if request.Users != nil {
		users, err := rs.userService.GetByIds(request.Users)
		if err != nil {
			return nil, err
		}

		updateUserErr := rs.repository.UpdateUsers(users, tx)
		if updateUserErr != nil {
			return nil, updateUserErr
		}
	}
	return result, nil
}

func (rs *RaceService) Delete(id uint) error {
	return rs.repository.Delete(id)
}
