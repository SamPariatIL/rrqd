package services

import "github.com/SamPariatIL/rrqd/repository"

type NodeService interface {
	GetNodes()
	GetNodeById()
	GetPackages()
	GetPackageById()
	GetBags()
	GetBagById()
	GetNodeByInchargeId()
	GetIssuedShipments()
	GetNodeUsers()
	GetComplaints()
	GetRiderPayments()
	GetVendorPayments()
	GetReturns()
	GetReturnDetails()
	GetMultipleReturns()
	CreateNode()
	RaiseComplaint()
	UpdateNode()
	UpdateNodeStatus()
}

type nodeService struct {
	nodeRepository repository.NodeRepository
}

func NewNodeService(nr repository.NodeRepository) NodeService {
	return &nodeService{nodeRepository: nr}
}

func (ns *nodeService) GetNodes() {}

func (ns *nodeService) GetNodeById() {}

func (ns *nodeService) GetPackages() {}

func (ns *nodeService) GetPackageById() {}

func (ns *nodeService) GetBags() {}

func (ns *nodeService) GetBagById() {}

func (ns *nodeService) GetNodeByInchargeId() {}

func (ns *nodeService) GetIssuedShipments() {}

func (ns *nodeService) GetNodeUsers() {}

func (ns *nodeService) GetComplaints() {}

func (ns *nodeService) GetRiderPayments() {}

func (ns *nodeService) GetVendorPayments() {}

func (ns *nodeService) GetReturns() {}

func (ns *nodeService) GetReturnDetails() {}

func (ns *nodeService) GetMultipleReturns() {}

func (ns *nodeService) CreateNode() {}

func (ns *nodeService) RaiseComplaint() {}

func (ns *nodeService) UpdateNode() {}

func (ns *nodeService) UpdateNodeStatus() {}
