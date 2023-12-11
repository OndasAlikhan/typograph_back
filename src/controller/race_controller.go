package controller

import (
	"net/http"
	"typograph_back/src/dto"
	"typograph_back/src/service"
	"typograph_back/src/service/service_interface"

	"github.com/labstack/echo/v4"
)

type RaceController struct {
	*BaseController
	service               service_interface.RaceServiceInterface
	userRaceResultService service_interface.UserRaceResultServiceInterface
}

func NewRaceController() *RaceController {
	return &RaceController{
		BaseController:        NewBaseController(),
		service:               service.NewRaceService(),
		userRaceResultService: service.NewUserRaceResultService(),
	}
}

// Index
// @title Index
// @description List of races
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags race
// @success 200 {object} dto.JSONResult{data=[]dto.RaceResponse}
// @router /races [get]
func (rc *RaceController) Index(c echo.Context) error {
	races, err := rc.service.GetAll()
	if err != nil {
		return err
	}

	response := make([]*dto.RaceResponse, len(races))
	for ind, val := range races {
		response[ind] = dto.NewRaceResponse(val)
	}

	return rc.json(http.StatusOK, response, c)
}

// Show
// @title Show
// @description Get a race by id
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags race
// @success 200 {object} dto.JSONResult{data=dto.RaceResponse}
// @router /races/{id} [get]
func (rc *RaceController) Show(c echo.Context) error {
	id, err := rc.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}
	race, err := rc.service.GetById(id)

	return rc.json(http.StatusOK, dto.NewRaceResponse(race), c)
}

// Store
// @title Create
// @description Create a new race
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags race
// @param raceStoreRequest body dto.RaceCreateRequest true "Race store request"
// @success 200 {object} dto.JSONResult{data=dto.RaceResponse}
// @router /races [post]
func (rc *RaceController) Store(c echo.Context) error {
	request := dto.RaceCreateRequest{}
	if err := rc.handleRequest(&request, c); err != nil {
		return err
	}

	race, err := rc.service.Create(&request)
	if err != nil {
		return err
	}
	return rc.json(http.StatusOK, dto.NewRaceResponse(race), c)
}

// Update
// @title Update
// @description Update a race
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags race
// @param raceStoreRequest body dto.RaceUpdateRequest true "Race update request"
// @param id path int true "Race ID"
// @success 200 {object} dto.JSONResult{data=dto.RaceResponse}
// @router /races/{id} [patch]
func (rc *RaceController) Update(c echo.Context) error {
	id, err := rc.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	request := dto.RaceUpdateRequest{}
	if err := rc.handleRequest(&request, c); err != nil {
		return err
	}

	race, err := rc.service.Update(id, &request)
	if err != nil {
		return err
	}
	return rc.json(http.StatusOK, dto.NewRaceResponse(race), c)
}

// Delete
// @title Delete
// @description Delete a race
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags race
// @param id path int true "Race ID"
// @success 200 {object} dto.JSONResult
// @router /races/{id} [delete]
func (rc *RaceController) Delete(c echo.Context) error {
	id, err := rc.parseToUint(c.Param("id"))
	if err != nil {
		return err
	}

	if err = rc.service.Delete(id); err != nil {
		return err
	}

	return rc.json(http.StatusOK, nil, c)
}

// AddUserRaceResult
// @title AddUserRaceResult
// @description Add a race result for each user
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags race
// @param userRaceResultCreateRequest body dto.UserRaceResultCreateRequest true "Result update request"
// @success 200 {object} dto.JSONResult{data=dto.UserRaceResultResponse}
// @router /races/add_user_race_result [post]
func (rc *RaceController) AddUserRaceResult(c echo.Context) error {
	request := dto.UserRaceResultCreateRequest{}
	if err := rc.handleRequest(&request, c); err != nil {
		return err
	}

	result, err := rc.userRaceResultService.Create(&request)
	if err != nil {
		return err
	}

	return rc.json(http.StatusOK, dto.NewUserRaceResultResponse(result), c)
}
