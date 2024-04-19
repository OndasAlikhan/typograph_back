package service

import (
	"fmt"
	"typograph_back/src/dto"
	"typograph_back/src/model"
	"typograph_back/src/repository"
	"typograph_back/src/repository/repository_interface"
	"typograph_back/src/service/service_interface"
)

var _ service_interface.LobbyServiceInterface = (*LobbyService)(nil)

type LobbyService struct {
	repository     repository_interface.LobbyRepositoryInterface
	userService    service_interface.UserServiceInterface
	raceService    service_interface.RaceServiceInterface
	lobbyWsService LobbyWsService
}

func NewLobbyService() *LobbyService {
	return &LobbyService{repository: repository.NewLobbyRepository(), userService: NewUserService(), raceService: NewRaceService()}
}

func (ls *LobbyService) GetAll() ([]*model.Lobby, error) {
	return ls.repository.GetAll()
}

func (ls *LobbyService) GetById(id uint) (*model.Lobby, error) {
	lobby, _, err := ls.repository.GetById(id)
	return lobby, err
}

func (ls *LobbyService) Create(request *dto.LobbyCreateRequest) (*model.Lobby, error) {
	users, err := ls.userService.GetByIds(request.Users)
	if err != nil {
		return nil, err
	}

	races, err := ls.raceService.GetByIds(request.Races)

	lobby := model.Lobby{
		AdminUserID: request.AdminUserID,
		Status:      "active",
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

func (ls *LobbyService) EnterLobby(lobbyId uint, userId uint) error {
	lobby, tx, err := ls.repository.GetById(lobbyId)
	if err != nil {
		return err
	}

	newUser, err := ls.userService.GetById(userId)
	if err != nil {
		return fmt.Errorf("user with id %d not found: %s", userId, err)
	}

	lobbyUsers := append(lobby.Users, newUser)

	updateErr := ls.repository.UpdateUsers(lobbyUsers, tx)
	if updateErr != nil {
		return updateErr
	}
	return nil
}

func (ls *LobbyService) LeaveLobby(lobbyId uint, userId uint) error {
	lobby, tx, err := ls.repository.GetById(lobbyId)
	if err != nil {
		return err
	}

	updatedUsers := make([]*model.User, 0)
	for _, user := range lobby.Users {
		if user.ID != userId {
			updatedUsers = append(updatedUsers, user)
		}
	}
	updateErr := ls.repository.UpdateUsers(updatedUsers, tx)
	if updateErr != nil {
		return updateErr
	}
	return nil
}

func (ls *LobbyService) Update(id uint, request *dto.LobbyUpdateRequest) (*model.Lobby, error) {
	lobby, _, err := ls.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	lobby.Name = request.Name
	lobby.AdminUserID = request.AdminUserID
	lobby.Status = request.Status

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
