package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type pocDetails = vendorDetails
type locationAddress = pickupLocation

type Location struct {
	Id          primitive.ObjectID `bson:"_id"`
	PrincipalId primitive.ObjectID `bson:"principalId"`
	PocDetails  pocDetails         `bson:"pocDetails"`
	Address     locationAddress    `bson:"address"`
	CreatedBy   primitive.ObjectID `bson:"createdBy"`
	Status      string             `bson:"status"`
	CreatedAt   primitive.DateTime `bson:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at"`
}
