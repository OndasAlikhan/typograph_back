package controller

import (
	"net/http"
	"typograph_back/src/dto"
	"typograph_back/src/service"
	"typograph_back/src/service/service_interface"

	"github.com/labstack/echo/v4"
)

type ParagraphController struct {
	*BaseController
	service service_interface.ParagraphServiceInterface
}

func NewParagraphController() *ParagraphController {
	return &ParagraphController{
		BaseController: NewBaseController(),
		service:        service.NewParagraphService(),
	}
}

func (this *ParagraphController) GetRandom(c echo.Context) error {
	paragraph, err := this.service.GetRandom()
	if err != nil {
		return err
	}

	response := dto.NewParagraphResponse(paragraph)

	return this.json(http.StatusOK, response, c)
}
