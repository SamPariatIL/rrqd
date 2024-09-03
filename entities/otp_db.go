package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type OTP struct {
	Id         primitive.ObjectID `bson:"_id"`
	Phone      string             `bson:"phone"`
	OTP        string             `bson:"otp"`
	CreatedAt  primitive.DateTime `bson:"createdAt"`
	ExpiryAt   primitive.DateTime `bson:"expiryAt"`
	Purpose    string             `bson:"purpose"`
	Verified   bool               `bson:"isVerified"`
	VerifiedAt primitive.DateTime `bson:"verifiedAt"`
	MessageId  string             `bson:"messageId"`
}
