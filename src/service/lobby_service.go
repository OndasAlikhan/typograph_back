package service

import (
	"fmt"
	"typograph_back/src/dto"
	"typograph_back/src/model"
	"typograph_back/src/repository/repository_interface"
	"typograph_back/src/service/service_interface"
)

var _ service_interface.LobbyServiceInterface = (*LobbyService)(nil)

type LobbyService struct {
	repository     repository_interface.LobbyRepositoryInterface
	userService    service_interface.UserServiceInterface
	raceService    service_interface.RaceServiceInterface
	lobbyWsService *LobbyWsService
}

func NewLobbyService(
	repo repository_interface.LobbyRepositoryInterface,
	userService service_interface.UserServiceInterface,
	raceService service_interface.RaceServiceInterface,
	lobbyWsService *LobbyWsService,
) *LobbyService {
	return &LobbyService{
		repository:     repo,
		userService:    userService,
		raceService:    raceService,
		lobbyWsService: lobbyWsService,
	}
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
		Status:      "waiting",
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
	lobby, _, err := ls.repository.GetById(lobbyId)
	if err != nil {
		return err
	}

	newUser, err := ls.userService.GetById(userId)
	if err != nil {
		return fmt.Errorf("user with id %d not found: %s", userId, err)
	}
	lobbyUsers := append(lobby.Users, newUser)

	updateErr := ls.repository.UpdateUsers(lobbyUsers, lobby)
	if updateErr != nil {
		return updateErr
	}

	ls.lobbyWsService.AddUserToRoom(lobbyId, newUser)

	return nil
}

// todo finish UserFinished
func (ls *LobbyService) UserFinished(request *dto.UserFinishedRequest) error {
	return nil
}

func (ls *LobbyService) LeaveLobby(lobbyId uint, userId uint) error {
	lobby, _, err := ls.repository.GetById(lobbyId)
	if err != nil {
		return err
	}
	fmt.Printf("before lobby.Users: %v \n", lobby.Users)

	updatedUsers := make([]*model.User, 0)
	for _, user := range lobby.Users {
		if user.ID != userId {
			updatedUsers = append(updatedUsers, user)
		}
	}
	fmt.Printf("after lobby.Users: %v \n", updatedUsers)

	updateErr := ls.repository.UpdateUsers(updatedUsers, lobby)
	if updateErr != nil {
		return updateErr
	}

	fmt.Println("before calling wsService")
	ls.lobbyWsService.RemoveUserFromRoom(lobbyId, userId)

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

	result, _, err := ls.repository.Save(*lobby)
	if err != nil {
		return nil, err
	}

	if request.Users != nil {
		users, err := ls.userService.GetByIds(request.Users)
		if err != nil {
			return nil, err
		}

		updateUserErr := ls.repository.UpdateUsers(users, result)
		if updateUserErr != nil {
			return nil, updateUserErr
		}
	}

	if request.Races != nil {
		races, err := ls.raceService.GetByIds(request.Races)
		if err != nil {
			return nil, err
		}

		updateRaceErr := ls.repository.UpdateRaces(races, result)
		if updateRaceErr != nil {
			return nil, updateRaceErr
		}
	}

	return result, nil
}

func (ls *LobbyService) Delete(id uint) error {
	return ls.repository.Delete(id)
}

func (ls *LobbyService) StartLobby(id uint) error {
	lobby, _, err := ls.repository.GetById(id)
	if err != nil {
		return err
	}

	lobby.Status = "starting"
	_, _, saveErr := ls.repository.Save(*lobby)
	if saveErr != nil {
		return saveErr
	}

	// todo change status in lobby_ws_repository
	// then after 3 seconds change status to running

	return nil
}
