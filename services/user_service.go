package services

import "github.com/SamPariatIL/rrqd/repository"

type UserService interface {
	GetUsers()
	GetUserById()
	CreateUser()
	UpdateUser()
	UpdateUserStatus()
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{userRepo: ur}
}

func (us *userService) GetUsers() {}

func (us *userService) GetUserById() {}

func (us *userService) CreateUser() {}

func (us *userService) UpdateUser() {}

func (us *userService) UpdateUserStatus() {}
