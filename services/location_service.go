package services

import "github.com/SamPariatIL/rrqd/repository"

type LocationService interface {
	GetLocations()
}

type locationService struct {
	locationRepo repository.LocationRepository
}

func NewLocationService(lr repository.LocationRepository) LocationService {
	return &locationService{locationRepo: lr}
}

func (ls *locationService) GetLocations() {}
