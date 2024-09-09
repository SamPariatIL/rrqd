package repository

import "go.mongodb.org/mongo-driver/mongo"

type NodeRepository interface {
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

type nodeRepository struct {
	db *mongo.Database
}

func NewNodeRepository(db *mongo.Database) NodeRepository {
	return &nodeRepository{db: db}
}

func (nr *nodeRepository) GetNodes() {}

func (nr *nodeRepository) GetNodeById() {}

func (nr *nodeRepository) GetPackages() {}

func (nr *nodeRepository) GetPackageById() {}

func (nr *nodeRepository) GetBags() {}

func (nr *nodeRepository) GetBagById() {}

func (nr *nodeRepository) GetNodeByInchargeId() {}

func (nr *nodeRepository) GetIssuedShipments() {}

func (nr *nodeRepository) GetNodeUsers() {}

func (nr *nodeRepository) GetComplaints() {}

func (nr *nodeRepository) GetRiderPayments() {}

func (nr *nodeRepository) GetVendorPayments() {}

func (nr *nodeRepository) GetReturns() {}

func (nr *nodeRepository) GetReturnDetails() {}

func (nr *nodeRepository) GetMultipleReturns() {}

func (nr *nodeRepository) CreateNode() {}

func (nr *nodeRepository) RaiseComplaint() {}

func (nr *nodeRepository) UpdateNode() {}

func (nr *nodeRepository) UpdateNodeStatus() {}
