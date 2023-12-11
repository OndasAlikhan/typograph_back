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

// GetRandom
// @title GetRandom
// @description Get random quote
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags paragraph
// @success 200	{object} dto.JSONResult{data=dto.ParagraphResponse}
// @router /random_paragraph [get]
func (pc *ParagraphController) GetRandom(c echo.Context) error {
	paragraph, err := pc.service.GetRandom()
	if err != nil {
		return err
	}

	response := dto.NewParagraphResponse(paragraph)

	return pc.json(http.StatusOK, response, c)
}
