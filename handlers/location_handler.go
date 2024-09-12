package handlers

import (
	"github.com/SamPariatIL/rrqd/services"
	"github.com/gin-gonic/gin"
)

type LocationHandler interface {
	GetLocations(ctx *gin.Context)
}

type locationHandler struct {
	locationService services.LocationService
}

func NewLocationHandler(ls services.LocationService) LocationHandler {
	return &locationHandler{locationService: ls}
}

func (lh *locationHandler) GetLocations(ctx *gin.Context) {}
