package repository

import (
	"typograph_back/src/dto"
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

func (ur *UserRaceResultRepository) GetByUserId(userId uint) ([]*model.UserRaceResult, error) {
	var values []*model.UserRaceResult
	err := ur.connection.Where("user_id = ?", userId).Order("created_at DESC").Find(&values).Error

	return values, err
}

func (ur *UserRaceResultRepository) Leaderboard() ([]*dto.LeaderboardResponse, error) {
	var values []*dto.LeaderboardResponse
	// err := ur.connection.Select("id, user_id, max(wpm) as wpm").Group("user_id").Joins("JOIN user_race_results r on ").Order("wpm DESC").Limit(10).Find(&values).Error

	err := ur.connection.Raw(`SELECT 
		u.name as user_name, r1.wpm, r2.id, r2.duration, r2.accuracy, r2.created_at
		FROM (SELECT user_id, MAX(wpm) as wpm FROM user_race_results GROUP BY user_id) r1 
		JOIN user_race_results r2 
			ON r1.user_id = r2.user_id AND r1.wpm = r2.wpm
		JOIN users u
			on r2.user_id = u.id
		ORDER BY r1.wpm DESC
		LIMIT 10`).Scan(&values).Error
	return values, err
}
