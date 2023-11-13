package service

import (
	"typograph_back/src/model"
	"typograph_back/src/repository"
	"typograph_back/src/repository/repository_interface"
)

type ParagraphService struct {
	repository repository_interface.ParagraphRepositoryInterface
}

func NewParagraphService() *ParagraphService {
	return &ParagraphService{repository: repository.NewParagraphRepository()}
}

func (this *ParagraphService) GetRandom() (*model.Paragraph, error) {
	return this.repository.GetRandom()
}
