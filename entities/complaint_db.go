package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Complaint struct {
	Id           primitive.ObjectID `bson:"_id"`
	ShipmentId   string             `bson:"shipmentId"`
	ShipmentType string             `bson:"shipmentType"`
	RouteId      string             `bson:"routeId"`
	UserType     string             `bson:"userType"`
	Reason       string             `bson:"reason"`
	CreatedBy    primitive.ObjectID `bson:"createdBy"`
	Status       string             `bson:"status"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
	UpdatedAt    primitive.DateTime `bson:"updated_at"`
}
