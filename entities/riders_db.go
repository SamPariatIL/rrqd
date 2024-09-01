package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type riderAddress = vendorAddress

type riderDetails = vendorDetails

type drivingLicenseDetails struct {
	LicenseNumber string `bson:"licenseNumber"`
	LicenseExpiry string `bson:"licenseExpiry"`
}

type Rider struct {
	Id            primitive.ObjectID `bson:"_id"`
	VendorId      primitive.ObjectID `bson:"vendorId"`
	RiderDetails  riderDetails       `bson:"riderDetails"`
	Address       riderAddress       `bson:"address"`
	AadharDetails struct {
		AadharNumber string `bson:"aadharNumber"`
	} `bson:"aadharDetails"`
	DrivingLicenseDetails  drivingLicenseDetails `bson:"drivingLicenseDetails"`
	ServiceableNodeId      string                `bson:"serviceableNodeId"`
	CreatedBy              primitive.ObjectID    `bson:"createdBy"`
	RiderImageUrl          string                `bson:"rider_pic_url"`
	DrivingLicenseFrontUrl string                `bson:"drivingLicenseFront_pic_url"`
	DrivingLicenseBackUrl  string                `bson:"drivingLicenseBack_pic_url"`
	AadharFrontUrl         string                `bson:"aadharFront_pic_url"`
	AadharBackUrl          string                `bson:"aadharBack_pic_url"`
	Status                 string                `bson:"status"`
	CreatedAt              primitive.DateTime    `bson:"created_at"`
	UpdatedAt              primitive.DateTime    `bson:"updated_at"`
}
