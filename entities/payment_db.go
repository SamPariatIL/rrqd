package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	Id            primitive.ObjectID `bson:"_id"`
	PackageId     string             `bson:"packageId"`
	TransactionId string             `bson:"transactionId"`
	PaymentUrl    string             `bson:"paymentLink"`
	CreatedBy     string             `bson:"createdBy"`
	Status        string             `bson:"status"`
	CreatedAt     primitive.DateTime `bson:"createdAt"`
	UpdatedAt     primitive.DateTime `bson:"updatedAt"`
}
