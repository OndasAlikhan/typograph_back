package repository

import (
	"math/rand"
	"typograph_back/src/application"
	"typograph_back/src/model"

	"gorm.io/gorm"
)

type ParagraphRepository struct {
	connection *gorm.DB
}

func NewParagraphRepository() *ParagraphRepository {
	return &ParagraphRepository{connection: application.GlobalDB}
}

func (this *ParagraphRepository) GetRandom() (*model.Paragraph, error) {
	var count int64
	this.connection.Model(&model.Paragraph{}).Count(&count)

	randomInt := rand.Intn(int(count))

	var paragraph *model.Paragraph
	err := this.connection.Limit(1).Offset(randomInt).Find(&paragraph).Error

	return paragraph, err
}

func (this *ParagraphRepository) GetById(id uint) (*model.Paragraph, error) {
	var paragraph model.Paragraph
	err := this.connection.First(&paragraph, id).Error

	return &paragraph, err
}

func (this *ParagraphRepository) Save(paragraph *model.Paragraph) (*model.Paragraph, error) {
	result := this.connection.Save(&paragraph)

	return paragraph, result.Error
}

func (this *ParagraphRepository) Delete(id uint) error {
	var paragraph model.Paragraph
	if err := this.connection.First(&paragraph, id).Error; err != nil {
		return err
	}

	return this.connection.Delete(&paragraph).Error
}
