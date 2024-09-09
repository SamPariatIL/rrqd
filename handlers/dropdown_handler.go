package handlers

import (
	"github.com/SamPariatIL/rrqd/services"
	"github.com/gin-gonic/gin"
)

type DropdownHandler interface {
	GetStates(ctx *gin.Context)
	GetCities(ctx *gin.Context)
	GetPincodes(ctx *gin.Context)
	GetFreights(ctx *gin.Context)
	GetRoles(ctx *gin.Context)
	GetIssueTypes(ctx *gin.Context)
	GetNodes(ctx *gin.Context)
	GetCityNodes(ctx *gin.Context)
	GetVendors(ctx *gin.Context)
	GetVendorsByNodeId(ctx *gin.Context)
	GetNodeSections(ctx *gin.Context)
	GetNodeIncharges(ctx *gin.Context)
	GetNodeRiders(ctx *gin.Context)
	GetPincodeDetails(ctx *gin.Context)
	GetNodeHubs(ctx *gin.Context)
}

type dropdownHandler struct {
	dropdownService services.DropdownService
}

func NewDropdownHandler(ds services.DropdownService) DropdownHandler {
	return &dropdownHandler{dropdownService: ds}
}

func (dh *dropdownHandler) GetStates(ctx *gin.Context) {}

func (dh *dropdownHandler) GetCities(ctx *gin.Context) {}

func (dh *dropdownHandler) GetPincodes(ctx *gin.Context) {}

func (dh *dropdownHandler) GetFreights(ctx *gin.Context) {}

func (dh *dropdownHandler) GetRoles(ctx *gin.Context) {}

func (dh *dropdownHandler) GetIssueTypes(ctx *gin.Context) {}

func (dh *dropdownHandler) GetNodes(ctx *gin.Context) {}

func (dh *dropdownHandler) GetCityNodes(ctx *gin.Context) {}

func (dh *dropdownHandler) GetVendors(ctx *gin.Context) {}

func (dh *dropdownHandler) GetVendorsByNodeId(ctx *gin.Context) {}

func (dh *dropdownHandler) GetNodeSections(ctx *gin.Context) {}

func (dh *dropdownHandler) GetNodeIncharges(ctx *gin.Context) {}

func (dh *dropdownHandler) GetNodeRiders(ctx *gin.Context) {}

func (dh *dropdownHandler) GetPincodeDetails(ctx *gin.Context) {}

func (dh *dropdownHandler) GetNodeHubs(ctx *gin.Context) {}
