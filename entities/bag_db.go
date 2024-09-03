package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type nextNodeAddress = nodeAddress

type Bag struct {
	Id              primitive.ObjectID `bson:"_id"`
	RouteId         string             `bson:"routeId"`
	PackageIds      []string           `bson:"packageIds"`
	CreatedBy       primitive.ObjectID `bson:"createdBy"`
	TotalPackages   int                `bson:"totalPackage"`
	BagId           string             `bson:"bagId"`
	CurrentNodeId   string             `bson:"currentNodeId"`
	NextNodeAddress nextNodeAddress    `bson:"nextNodeAddress"`
	Status          string             `bson:"status"`
	CreatedAt       primitive.DateTime `bson:"created_at"`
	UpdatedAt       primitive.DateTime `bson:"updated_at"`
}
