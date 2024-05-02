package controller

import (
	"fmt"
	"net/http"
	"typograph_back/src/dto"
	"typograph_back/src/service/service_interface"

	"github.com/labstack/echo/v4"
)

type RaceController struct {
	*BaseController
	service               service_interface.RaceServiceInterface
	userRaceResultService service_interface.UserRaceResultServiceInterface
}

func NewRaceController(
	raceService service_interface.RaceServiceInterface,
	userRaceResultService service_interface.UserRaceResultServiceInterface,
) *RaceController {
	return &RaceController{
		BaseController:        NewBaseController(),
		service:               raceService,
		userRaceResultService: userRaceResultService,
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
	race, _ := rc.service.GetById(id)

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

	fmt.Printf("request: %v\n", request)

	result, err := rc.userRaceResultService.Create(&request)
	if err != nil {
		return err
	}

	return rc.json(http.StatusOK, dto.NewUserRaceResultResponse(result), c)
}

// GetUserRaceResult
// @title GetUserRaceResult
// @description Get a list of result by user id
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags race
// @param user_id path int true "User ID"
// @success 200 {object} dto.JSONResult{data=[]dto.UserRaceResultResponse}
// @router /races/get_user_race_result/:user_id [get]
func (rc *RaceController) GetUserRaceResults(c echo.Context) error {
	id, err := rc.parseToUint(c.Param("user_id"))
	if err != nil {
		return err
	}

	results, err := rc.userRaceResultService.GetByUserId(id)
	if err != nil {
		return err
	}

	response := make([]*dto.UserRaceResultResponse, len(results))
	for ind, val := range results {
		response[ind] = dto.NewUserRaceResultResponse(val)
	}

	return rc.json(http.StatusOK, response, c)
}

// Leaderboard
// @title Leaderboard
// @description Get a list of the best wpm results
// @accept json
// @produce json
// @security ApiKeyAuth
// @tags race
// @success 200 {object} dto.JSONResult{data=[]dto.UserRaceResultResponse}
// @router /races/leaderboard [get]
func (rc *RaceController) Leaderboard(c echo.Context) error {
	results, err := rc.userRaceResultService.Leaderboard()
	if err != nil {
		return err
	}

	response := make([]*dto.LeaderboardResponse, len(results))
	for ind, val := range results {
		response[ind] = dto.NewLeaderboardResponse(val)
	}

	return rc.json(http.StatusOK, response, c)
}
