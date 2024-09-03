package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type geocode struct {
	Latitude  string `bson:"lat"`
	Longitude string `bson:"long"`
}

type nodeDetails struct {
	Name   string `bson:"name"`
	NodeId int    `bson:"nodeId"`
}

type nodeAddress struct {
	Street  string  `bson:"street"`
	City    string  `bson:"city"`
	State   string  `bson:"state"`
	Country string  `bson:"country"`
	Pincode string  `bson:"pincode"`
	Geocode geocode `bson:"geocode"`
}

type serviceableArea struct {
	State string `bson:"state"`
	City  struct {
		Id   string `bson:"_id"`
		City string `bson:"city"`
	} `bson:"city"`
	Pincodes []string `bson:"pincodes"`
}

type nodeArea struct {
	AreaName string   `bson:"areaName"`
	Sections []string `bson:"sections"`
}

type Node struct {
	Id           primitive.ObjectID `bson:"_id"`
	NodeDetails  nodeDetails        `bson:"nodeDetails"`
	Address      nodeAddress        `bson:"address"`
	NodeType     string             `bson:"nodeType"`
	NodeIncharge struct {
		NodeUserId primitive.ObjectID `bson:"nodeUserId"`
	} `bson:"nodeIncharge"`
	ServiceableArea []serviceableArea  `bson:"serviceableArea"`
	NodeAreas       []nodeArea         `bson:"nodeAreas"`
	CreatedBy       primitive.ObjectID `bson:"createdBy"`
	Status          string             `bson:"status"`
	CreatedAt       primitive.DateTime `bson:"created_at"`
	UpdatedAt       primitive.DateTime `bson:"updated_at"`
}
