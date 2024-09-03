package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type vendorDetails struct {
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
	Email string `bson:"email"`
}

type vendorAddress struct {
	Street  string `bson:"street"`
	City    string `bson:"city"`
	State   string `bson:"state"`
	Country string `bson:"country"`
	Pincode int    `bson:"pincode"`
}

type Vendor struct {
	Id            primitive.ObjectID `bson:"_id"`
	VendorDetails vendorDetails      `bson:"vendorDetails"`
	Address       vendorAddress      `bson:"address"`
	CreatedBy     primitive.ObjectID `bson:"createdBy"`
	VendorUserId  primitive.ObjectID `bson:"vendorUserId"`
	AgreementUrl  string             `bson:"agreementDocumentFileUrl"`
	Status        string             `bson:"status"`
	CreatedAt     primitive.DateTime `bson:"created_at"`
	UpdatedAt     primitive.DateTime `bson:"updated_at"`
}
