package repository

import (
	"typograph_back/src/model"
	"typograph_back/src/repository/repository_interface"
)

var _ repository_interface.UserRaceResultRepositoryInterface = (*UserRaceResultRepository)(nil)

type UserRaceResultRepository struct {
	*BaseRepository[model.UserRaceResult]
}

func NewUserRaceResulRepository() *UserRaceResultRepository {
	return &UserRaceResultRepository{BaseRepository: NewBaseRepository[model.UserRaceResult]()}
}
