package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pincode struct {
	Id        primitive.ObjectID `bson:"_id"`
	PincodeId int                `bson:"pincode_id"`
	Pincode   string             `bson:"pincode"`
	CityId    int                `bson:"city_id"`
}

type State struct {
	Id      primitive.ObjectID `bson:"_id"`
	StateId int                `bson:"id"`
	State   string             `bson:"name"`
}

type City struct {
	Id      primitive.ObjectID `bson:"_id"`
	CityId  int                `bson:"city_id"`
	City    string             `bson:"name"`
	StateId int                `bson:"state_id"`
}

type Freight struct {
	Id        primitive.ObjectID `bson:"_id"`
	FreightId int                `bson:"id"`
	Freight   string             `bson:"type"`
}

type Issue struct {
	Id      primitive.ObjectID `bson:"_id"`
	IssueId int                `bson:"id"`
	Issue   string             `bson:"type"`
	Purpose string             `bson:"purpose"`
}

type Role struct {
	Id          primitive.ObjectID `bson:"_id"`
	RoleId      string             `bson:"id"`
	DisplayName string             `bson:"role_name"`
	RoleName    string             `bson:"name"`
}
