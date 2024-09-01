package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type returnDropLocation = pickupLocation
type returnPickupLocation = vendorAddress

type orderDetails struct {
	PaymentMode     string `bson:"paymentMode"`
	Amount          int    `bson:"amount"`
	CustomerDetails struct {
		Name    string `bson:"name"`
		OrderId string `bson:"orderId"`
		Phone   string `bson:"phone"`
	} `bson:"customerDetails"`
}

type Return struct {
	Id             primitive.ObjectID   `bson:"_id"`
	ReturnId       string               `bson:"returnId"`
	OrderId        string               `bson:"orderId"`
	DropLocation   returnDropLocation   `bson:"dropLocation"`
	PickupLocation returnPickupLocation `bson:"pickupLocation"`
	OrderDetails   orderDetails         `bson:"orderDetails"`
	RouteId        string               `bson:"routeId"`
	CreatedBy      primitive.ObjectID   `bson:"createdBy"`
	PackageId      string               `bson:"packageId"`
	ImageUrls      []string             `bson:"imageUrls"`
	Status         string               `bson:"status"`
	PickupBy       string               `bson:"pickupBy"`
	Origin         string               `bson:"origin"`
	LastNodeId     string               `bson:"lastNodeId"`
	CreatedAt      primitive.DateTime   `bson:"created_at"`
	UpdatedAt      primitive.DateTime   `bson:"updated_at"`
}
