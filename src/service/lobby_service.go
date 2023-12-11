package service

import (
	"typograph_back/src/dto"
	"typograph_back/src/model"
	"typograph_back/src/repository"
	"typograph_back/src/repository/repository_interface"
	"typograph_back/src/service/service_interface"
)

var _ service_interface.LobbyServiceInterface = (*LobbyService)(nil)

type LobbyService struct {
	repository  repository_interface.LobbyRepositoryInterface
	userService service_interface.UserServiceInterface
	raceService service_interface.RaceServiceInterface
}

func NewLobbyService() *LobbyService {
	return &LobbyService{repository: repository.NewLobbyRepository(), userService: NewUserService(), raceService: NewRaceService()}
}

func (ls *LobbyService) GetAll() ([]*model.Lobby, error) {
	return ls.repository.GetAll()
}

func (ls *LobbyService) GetById(id uint) (*model.Lobby, error) {
	return ls.repository.GetById(id)
}

func (ls *LobbyService) Create(request *dto.LobbyCreateRequest) (*model.Lobby, error) {
	users, err := ls.userService.GetByIds(request.Users)
	if err != nil {
		return nil, err
	}

	races, err := ls.raceService.GetByIds(request.Races)

	lobby := model.Lobby{
		AdminUserID: request.AdminUserID,
		Name:        request.Name,
		Users:       users,
		Races:       races,
	}

	value, _, err := ls.repository.Save(lobby)
	if err != nil {
		return nil, err
	}

	return value, err
}

func (ls *LobbyService) Update(id uint, request *dto.LobbyUpdateRequest) (*model.Lobby, error) {
	lobby, err := ls.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	lobby.Name = request.Name
	lobby.AdminUserID = request.AdminUserID

	result, tx, err := ls.repository.Save(*lobby)
	if err != nil {
		return nil, err
	}

	if request.Users != nil {
		users, err := ls.userService.GetByIds(request.Users)
		if err != nil {
			return nil, err
		}

		updateUserErr := ls.repository.UpdateUsers(users, tx)
		if updateUserErr != nil {
			return nil, updateUserErr
		}
	}

	if request.Races != nil {
		races, err := ls.raceService.GetByIds(request.Races)
		if err != nil {
			return nil, err
		}

		updateRaceErr := ls.repository.UpdateRaces(races, tx)
		if updateRaceErr != nil {
			return nil, updateRaceErr
		}
	}

	return result, nil

}

func (ls *LobbyService) Delete(id uint) error {
	return ls.repository.Delete(id)
}
