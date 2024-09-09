package handlers

import (
	"github.com/SamPariatIL/rrqd/services"
	"github.com/gin-gonic/gin"
)

type NodeHandler interface {
	GetNodes(ctx *gin.Context)
	GetNodeById(ctx *gin.Context)
	GetPackages(ctx *gin.Context)
	GetPackageById(ctx *gin.Context)
	GetBags(ctx *gin.Context)
	GetBagById(ctx *gin.Context)
	GetNodeByInchargeId(ctx *gin.Context)
	GetIssuedShipments(ctx *gin.Context)
	GetNodeUsers(ctx *gin.Context)
	GetComplaints(ctx *gin.Context)
	GetRiderPayments(ctx *gin.Context)
	GetVendorPayments(ctx *gin.Context)
	GetReturns(ctx *gin.Context)
	GetReturnDetails(ctx *gin.Context)
	GetMultipleReturns(ctx *gin.Context)
	CreateNode(ctx *gin.Context)
	RaiseComplaint(ctx *gin.Context)
	UpdateNode(ctx *gin.Context)
	UpdateNodeStatus(ctx *gin.Context)
}

type nodeHandler struct {
	nodeService services.NodeService
}

func NewNodeHandler(ns services.NodeService) NodeHandler {
	return &nodeHandler{nodeService: ns}
}

func (nh *nodeHandler) GetNodes(ctx *gin.Context) {}

func (nh *nodeHandler) GetNodeById(ctx *gin.Context) {}

func (nh *nodeHandler) GetPackages(ctx *gin.Context) {}

func (nh *nodeHandler) GetPackageById(ctx *gin.Context) {}

func (nh *nodeHandler) GetBags(ctx *gin.Context) {}

func (nh *nodeHandler) GetBagById(ctx *gin.Context) {}

func (nh *nodeHandler) GetNodeByInchargeId(ctx *gin.Context) {}

func (nh *nodeHandler) GetIssuedShipments(ctx *gin.Context) {}

func (nh *nodeHandler) GetNodeUsers(ctx *gin.Context) {}

func (nh *nodeHandler) GetComplaints(ctx *gin.Context) {}

func (nh *nodeHandler) GetRiderPayments(ctx *gin.Context) {}

func (nh *nodeHandler) GetVendorPayments(ctx *gin.Context) {}

func (nh *nodeHandler) GetReturns(ctx *gin.Context) {}

func (nh *nodeHandler) GetReturnDetails(ctx *gin.Context) {}

func (nh *nodeHandler) GetMultipleReturns(ctx *gin.Context) {}

func (nh *nodeHandler) CreateNode(ctx *gin.Context) {}

func (nh *nodeHandler) RaiseComplaint(ctx *gin.Context) {}

func (nh *nodeHandler) UpdateNode(ctx *gin.Context) {}

func (nh *nodeHandler) UpdateNodeStatus(ctx *gin.Context) {}
