package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type userAddress struct {
	Street  string `bson:"street"`
	City    string `bson:"city"`
	State   string `bson:"state"`
	Country string `bson:"country"`
	Pincode string `bson:"pincode"`
}

type role struct {
	RoleId   string `bson:"roleId"`
	RoleName string `bson:"roleName"`
}

type node struct{}

type vendor struct{}

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Phone     string             `bson:"phone"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Status    string             `bson:"status"`
	UserId    string             `bson:"user_id"`
	ImageUrl  *string            `bson:"image_url"`
	Address   userAddress        `bson:"address"`
	Role      role               `bson:"role"`
	Node      *node              `bson:"node"`
	Vendor    *vendor            `bson:"vendor"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}
