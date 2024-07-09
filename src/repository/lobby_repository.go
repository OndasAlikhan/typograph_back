package repository

import (
	"fmt"
	"typograph_back/src/model"
	"typograph_back/src/repository/repository_interface"

	"gorm.io/gorm"
)

var _ repository_interface.LobbyRepositoryInterface = (*LobbyRepository)(nil)

type LobbyRepository struct {
	*BaseRepository[model.Lobby]
}

func NewLobbyRepository() *LobbyRepository {
	return &LobbyRepository{BaseRepository: NewBaseRepository[model.Lobby]()}
}

func (lr *LobbyRepository) GetAll() ([]*model.Lobby, error) {
	var values []*model.Lobby
	err := lr.connection.Preload("Users").Find(&values).Error

	return values, err
}

func (lr *LobbyRepository) GetById(id uint) (*model.Lobby, *gorm.DB, error) {
	var value *model.Lobby
	result := lr.connection.Preload("Users").First(&value, id)
	err := result.Error

	fmt.Printf("model.Lobby: %v\n", value)
	return value, result, err
}
func (lr *LobbyRepository) Save(lobby model.Lobby) (*model.Lobby, *gorm.DB, error) {
	result := lr.connection.Save(&lobby)

	return &lobby, result, result.Error
}
func (lr *LobbyRepository) UpdateUsers(users []*model.User, lobby *model.Lobby) error {
	fmt.Printf("UpdateUsers users: %v\n", users)
	return lr.connection.Model(lobby).Association("Users").Replace(users)
}

func (lr *LobbyRepository) UpdateRaces(races []*model.Race, lobby *model.Lobby) error {
	return lr.connection.Model(lobby).Association("Races").Replace(races)
}
