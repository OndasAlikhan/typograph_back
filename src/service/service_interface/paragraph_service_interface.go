package service_interface

import "typograph_back/src/model"

type ParagraphServiceInterface interface {
	GetRandom() (*model.Paragraph, error)
}
