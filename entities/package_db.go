package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type packageDropLocation = userAddress

type warehouseArea struct {
	Section string `bson:"section"`
	Area    string `bson:"name"`
}

type Package struct {
	Id             primitive.ObjectID  `bson:"_id"`
	DropLocation   packageDropLocation `bson:"dropLocation"`
	OrderDetails   orderDetails        `bson:"orderDetails"`
	OrderId        string              `bson:"orderId"`
	CreatedBy      primitive.ObjectID  `bson:"createdBy"`
	PickupLocation pickupLocation      `bson:"pickupLocation"`
	RouteId        string              `bson:"routeId"`
	PackageId      string              `bson:"packageId"`
	Status         string              `bson:"status"`
	LastNodeId     string              `bson:"lastNodeId"`
	WarehouseArea  warehouseArea       `bson:"warehouseArea"`
	CreatedAt      primitive.DateTime  `bson:"createdAt"`
	UpdatedAt      primitive.DateTime  `bson:"updatedAt"`
}
