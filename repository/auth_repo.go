package repository

import "go.mongodb.org/mongo-driver/mongo"

type AuthRepository interface{}

type authRepository struct {
	db *mongo.Database
}

func NewAuthRepository(db *mongo.Database) AuthRepository {
	return &authRepository{db: db}
}
