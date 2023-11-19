package repository

import (
	"typograph_back/src/application"

	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	connection *gorm.DB
}

func NewBaseRepository[T any]() *BaseRepository[T] {
	return &BaseRepository[T]{connection: application.GlobalDB}
}

func (this *BaseRepository[T]) GetById(id uint) (*T, error) {
	var value *T
	err := this.connection.First(&value, id).Error

	return value, err
}

func (this *BaseRepository[T]) GetAll() ([]*T, error) {
	var values []*T
	err := this.connection.Find(&values).Error

	return values, err
}

func (this *BaseRepository[T]) Save(value T) (*T, error) {
	result := this.connection.Save(&value)

	return &value, result.Error
}

func (this *BaseRepository[T]) Delete(id uint) error {
	var value T
	if err := this.connection.First(&value, id).Error; err != nil {
		return err
	}

	return this.connection.Delete(&value).Error
}
