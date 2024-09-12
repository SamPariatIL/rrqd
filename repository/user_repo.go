package repository

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository interface {
	GetUsers()
	GetUserById()
	CreateUser()
	UpdateUser()
	UpdateUserStatus()
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetUsers() {}

func (ur *userRepository) GetUserById() {}

func (ur *userRepository) CreateUser() {}

func (ur *userRepository) UpdateUser() {}

func (ur *userRepository) UpdateUserStatus() {}
