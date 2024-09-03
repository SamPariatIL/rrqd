package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type ScanHistory struct {
	Id        primitive.ObjectID `bson:"_id"`
	ScanId    string             `bson:"scanId"`
	ScanTime  primitive.DateTime `bson:"scannedTime"`
	ScannedBy primitive.ObjectID `bson:"scannedBy"`
}
