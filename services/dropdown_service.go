package services

import "github.com/SamPariatIL/rrqd/repository"

type DropdownService interface {
	GetStates()
	GetCities()
	GetPincodes()
	GetFreights()
	GetRoles()
	GetIssueTypes()
	GetNodes()
	GetCityNodes()
	GetVendors()
	GetVendorsByNodeId()
	GetNodeSections()
	GetNodeIncharges()
	GetNodeRiders()
	GetPincodeDetails()
	GetNodeHubs()
}

type dropdownService struct {
	dropdownRepo repository.DropdownRepository
}

func NewDropdownService(dr repository.DropdownRepository) DropdownService {
	return &dropdownService{
		dropdownRepo: dr,
	}
}

func (ds *dropdownService) GetStates() {}

func (ds *dropdownService) GetCities() {}

func (ds *dropdownService) GetPincodes() {}

func (ds *dropdownService) GetFreights() {}

func (ds *dropdownService) GetRoles() {}

func (ds *dropdownService) GetIssueTypes() {}

func (ds *dropdownService) GetNodes() {}

func (ds *dropdownService) GetCityNodes() {}

func (ds *dropdownService) GetVendors() {}

func (ds *dropdownService) GetVendorsByNodeId() {}

func (ds *dropdownService) GetNodeSections() {}

func (ds *dropdownService) GetNodeIncharges() {}

func (ds *dropdownService) GetNodeRiders() {}

func (ds *dropdownService) GetPincodeDetails() {}

func (ds *dropdownService) GetNodeHubs() {}
