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

func (this *RaceRepository) Save(value model.Race) (*model.Race, *gorm.DB, error) {
	result := this.connection.Save(&value)

	return &value, result, result.Error
}
