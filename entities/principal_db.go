package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type principalDetails struct {
	Name                  string `bson:"name"`
	Phone                 string `bson:"phone"`
	Email                 string `bson:"email"`
	CustomerSupportNumber string `bson:"customerSupportNumber"`
}

type principalAddress = vendorAddress

type Principal struct {
	Id               primitive.ObjectID `bson:"_id"`
	PrincipalDetails principalDetails   `bson:"principalDetails"`
	Address          principalAddress   `bson:"address"`
	CreatedBy        primitive.ObjectID `bson:"createdBy"`
	ContractUrl      string             `bson:"contract_pdf_url"`
	PrincipalImage   string             `bson:"principal_image_url"`
	Status           string             `bson:"status"`
	Role             role               `bson:"role"`
	Password         string             `bson:"password"`
	CreatedAt        primitive.DateTime `bson:"created_at"`
	UpdatedAt        primitive.DateTime `bson:"updated_at"`
}
