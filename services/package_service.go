package services

import "github.com/SamPariatIL/rrqd/repository"

type PackageService interface{}

type packageService struct {
	packageRepo repository.PackageRepository
}

func NewPackageService(pr repository.PackageRepository) PackageService {
	return &packageService{packageRepo: pr}
}
