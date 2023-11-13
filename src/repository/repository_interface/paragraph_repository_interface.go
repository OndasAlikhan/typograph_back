package repository_interface

import "typograph_back/src/model"

type ParagraphRepositoryInterface interface {
	GetRandom() (*model.Paragraph, error)
	GetById(id uint) (*model.Paragraph, error)
	Save(paragraph *model.Paragraph) (*model.Paragraph, error)
	Delete(id uint) error
}
