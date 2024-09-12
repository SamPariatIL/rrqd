package handlers

import (
	"github.com/SamPariatIL/rrqd/services"
	"github.com/gin-gonic/gin"
)

type DashboardHandler interface {
	GetAdminDashboard(ctx *gin.Context)
	GetNodeDashboard(ctx *gin.Context)
	GetVendorDashboard(ctx *gin.Context)
	GetAdminPackages(ctx *gin.Context)
	GetAdminBags(ctx *gin.Context)
	GetAdminReturns(ctx *gin.Context)
	GetPackageDetails(ctx *gin.Context)
	GetBagDetails(ctx *gin.Context)
	GetReturnDetails(ctx *gin.Context)
}

type dashboardHandler struct {
	dashboardService services.DashboardService
}

func NewDashboardHandler(ds services.DashboardService) DashboardHandler {
	return &dashboardHandler{dashboardService: ds}
}

func (dh *dashboardHandler) GetAdminDashboard(ctx *gin.Context) {}

func (dh *dashboardHandler) GetNodeDashboard(ctx *gin.Context) {}

func (dh *dashboardHandler) GetVendorDashboard(ctx *gin.Context) {}

func (dh *dashboardHandler) GetAdminPackages(ctx *gin.Context) {}

func (dh *dashboardHandler) GetAdminBags(ctx *gin.Context) {}

func (dh *dashboardHandler) GetAdminReturns(ctx *gin.Context) {}

func (dh *dashboardHandler) GetPackageDetails(ctx *gin.Context) {}

func (dh *dashboardHandler) GetBagDetails(ctx *gin.Context) {}

func (dh *dashboardHandler) GetReturnDetails(ctx *gin.Context) {}
