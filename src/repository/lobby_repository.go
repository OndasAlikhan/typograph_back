package repository

import (
	"gorm.io/gorm"
	"typograph_back/src/model"
	"typograph_back/src/repository/repository_interface"
)

var _ repository_interface.LobbyRepositoryInterface = (*LobbyRepository)(nil)

type LobbyRepository struct {
	*BaseRepository[model.Lobby]
}

func NewLobbyRepository() *LobbyRepository {
	return &LobbyRepository{BaseRepository: NewBaseRepository[model.Lobby]()}
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
