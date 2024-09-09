package repository

import "go.mongodb.org/mongo-driver/mongo"

type PackageRepository interface{}

type packageRepository struct {
	db *mongo.Database
}

func NewPackageRepository(db *mongo.Database) PackageRepository {
	return &packageRepository{db: db}
}
