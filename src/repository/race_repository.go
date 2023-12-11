package repository

import (
	"typograph_back/src/model"
	"typograph_back/src/repository/repository_interface"

	"gorm.io/gorm"
)

var _ repository_interface.RaceRepositoryInterface = (*RaceRepository)(nil)

type RaceRepository struct {
	*BaseRepository[model.Race]
}

func NewRaceRepository() *RaceRepository {
	return &RaceRepository{BaseRepository: NewBaseRepository[model.Race]()}
}

func (rr *RaceRepository) GetByIds(ids []uint) ([]*model.Race, error) {
	var races []*model.Race
	err := rr.connection.Where("id IN ?", ids).Find(&races).Error

	return races, err
}

func (rr *RaceRepository) Save(value model.Race) (*model.Race, *gorm.DB, error) {
	result := rr.connection.Save(&value)

	return &value, result, result.Error
}

func (rr *RaceRepository) UpdateUsers(users []*model.User, tx *gorm.DB) error {
	return tx.Association("Users").Replace(users)
}
