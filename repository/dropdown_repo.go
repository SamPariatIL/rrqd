package repository

import (
	"github.com/SamPariatIL/rrqd/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type DropdownRepository interface {
	GetStates(searchTerm *string) ([]entities.State, error)
	GetCities(stateId *int, searchTerm *string) ([]entities.City, error)
	GetPincodes(cityId *int, searchTerm *string) ([]entities.Pincode, error)
	GetFreights() ([]entities.Freight, error)
	GetRoles() ([]entities.Role, error)
	GetIssueTypes() ([]entities.Issue, error)
	GetNodes()
	GetCityNodes()
	GetVendors()
	GetVendorsByNodeId()
	GetNodeSections()
	GetNodeIncharge()
	GetNodeRiders()
	GetPincodeDetails()
	GetNodeHubs()
}

type dropdownRepository struct {
	db *mongo.Database
}

func NewDropdownRepository(db *mongo.Database) DropdownRepository {
	return &dropdownRepository{db: db}
}

func (dr *dropdownRepository) GetStates(searchTerm *string) ([]entities.State, error) {
	return nil, nil
}

func (dr *dropdownRepository) GetCities(stateId *int, searchTerm *string) ([]entities.City, error) {
	return nil, nil
}

func (dr *dropdownRepository) GetPincodes(cityId *int, searchTerm *string) ([]entities.Pincode, error) {
	return nil, nil
}

func (dr *dropdownRepository) GetFreights() ([]entities.Freight, error) {
	return nil, nil
}

func (dr *dropdownRepository) GetRoles() ([]entities.Role, error) {
	return nil, nil
}

func (dr *dropdownRepository) GetIssueTypes() ([]entities.Issue, error) {
	return nil, nil
}

func (dr *dropdownRepository) GetNodes() {}

func (dr *dropdownRepository) GetCityNodes() {}

func (dr *dropdownRepository) GetVendors() {}

func (dr *dropdownRepository) GetVendorsByNodeId() {}

func (dr *dropdownRepository) GetNodeSections() {}

func (dr *dropdownRepository) GetNodeIncharge() {}

func (dr *dropdownRepository) GetNodeRiders() {}

func (dr *dropdownRepository) GetPincodeDetails() {}

func (dr *dropdownRepository) GetNodeHubs() {}
