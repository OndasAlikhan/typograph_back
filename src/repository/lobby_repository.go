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
func (lr *LobbyRepository) UpdateUsers(users []*model.User, tx *gorm.DB) error {
	return tx.Association("Users").Replace(users)
}
func (lr *LobbyRepository) UpdateRaces(races []*model.Race, tx *gorm.DB) error {
	return tx.Association("Races").Replace(races)
}
