package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type pickupLocation struct {
	Street       string  `bson:"street"`
	City         string  `bson:"city"`
	State        string  `bson:"state"`
	Country      string  `bson:"country"`
	Pincode      int     `bson:"pincode"`
	GeoLocations geocode `bson:"geoLocations"`
}

type dropLocation = nodeAddress

type Shipment struct {
	Id             primitive.ObjectID `bson:"_id"`
	ShipmentId     string             `bson:"shipmentId"`
	ShipmentType   string             `bson:"shipmentType"`
	RiderId        primitive.ObjectID `bson:"riderId"`
	PickupLocation pickupLocation     `bson:"pickupLocation"`
	DropLocation   dropLocation       `bson:"dropLocation"`
	RouteId        string             `bson:"routeId"`
	Status         string             `bson:"status"`
	DropNodeId     string             `bson:"dropNodeId"`
	CreatedAt      primitive.DateTime `bson:"created_at"`
	UpdatedAt      primitive.DateTime `bson:"updated_at"`
}
