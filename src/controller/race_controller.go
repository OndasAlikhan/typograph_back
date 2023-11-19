package controller

import "typograph_back/src/service/service_interface"

type RaceController struct {
	*BaseController
	service service_interface.RaceServiceInterface
}
