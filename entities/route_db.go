package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Route struct {
	Id        primitive.ObjectID `bson:"_id"`
	RouteName string             `bson:"routeName"`
	Nodes     []string           `bson:"nodes"`
	Status    string             `bson:"status"`
	RouteId   string             `bson:"routeId"`
	CreatedBy primitive.ObjectID `bson:"createdBy"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}
