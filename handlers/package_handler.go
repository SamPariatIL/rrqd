package handlers

import (
	"github.com/SamPariatIL/rrqd/services"
	"github.com/gin-gonic/gin"
)

type PackageHandler interface {
	GetPackages(ctx *gin.Context)
	GetPackageById(ctx *gin.Context)
	GetPreviousIssue(ctx *gin.Context)
	GetSelectedPackages(ctx *gin.Context)
	GetSelectedBags(ctx *gin.Context)
	GetWarehouses(ctx *gin.Context)
	GetReturns(ctx *gin.Context)
	ScanPackage(ctx *gin.Context)
	CreatePackage(ctx *gin.Context)
	CreateMultiplePackages(ctx *gin.Context)
	CreateBulkReturns(ctx *gin.Context)
	BagPackage(ctx *gin.Context)
	ScanBag(ctx *gin.Context)
	AssignTransport(ctx *gin.Context)
	CreateIssue(ctx *gin.Context)
	AssignRider(ctx *gin.Context)
	AssignFirstMileRider(ctx *gin.Context)
	UpdatePrintStatus(ctx *gin.Context)
	InitiateReturn(ctx *gin.Context)
	ReceivePackage(ctx *gin.Context)
	MovePackage(ctx *gin.Context)
	ReceiveBag(ctx *gin.Context)
	MoveBag(ctx *gin.Context)
}

type packageHandler struct {
	packageService services.PackageService
}

func NewPackageHandler(ps services.PackageService) PackageHandler {
	return &packageHandler{packageService: ps}
}

func (ph *packageHandler) GetPackages(ctx *gin.Context) {}

func (ph *packageHandler) GetPackageById(ctx *gin.Context) {}

func (ph *packageHandler) GetPreviousIssue(ctx *gin.Context) {}

func (ph *packageHandler) GetSelectedPackages(ctx *gin.Context) {}

func (ph *packageHandler) GetSelectedBags(ctx *gin.Context) {}

func (ph *packageHandler) GetWarehouses(ctx *gin.Context) {}

func (ph *packageHandler) GetReturns(ctx *gin.Context) {}

func (ph *packageHandler) ScanPackage(ctx *gin.Context) {}

func (ph *packageHandler) CreatePackage(ctx *gin.Context) {}

func (ph *packageHandler) CreateMultiplePackages(ctx *gin.Context) {}

func (ph *packageHandler) CreateBulkReturns(ctx *gin.Context) {}

func (ph *packageHandler) BagPackage(ctx *gin.Context) {}

func (ph *packageHandler) ScanBag(ctx *gin.Context) {}

func (ph *packageHandler) AssignTransport(ctx *gin.Context) {}

func (ph *packageHandler) CreateIssue(ctx *gin.Context) {}

func (ph *packageHandler) AssignRider(ctx *gin.Context) {}

func (ph *packageHandler) AssignFirstMileRider(ctx *gin.Context) {}

func (ph *packageHandler) UpdatePrintStatus(ctx *gin.Context) {}

func (ph *packageHandler) InitiateReturn(ctx *gin.Context) {}

func (ph *packageHandler) ReceivePackage(ctx *gin.Context) {}

func (ph *packageHandler) MovePackage(ctx *gin.Context) {}

func (ph *packageHandler) ReceiveBag(ctx *gin.Context) {}

func (ph *packageHandler) MoveBag(ctx *gin.Context) {}
