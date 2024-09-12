package repository

import "go.mongodb.org/mongo-driver/mongo"

type LocationRepository interface {
	GetLocations()
}

type locationRepository struct {
	db *mongo.Database
}

func NewLocationRepository(db *mongo.Database) LocationRepository {
	return &locationRepository{db: db}
}

func (lr *locationRepository) GetLocations() {}
