package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Credential struct {
	ID primitive.ObjectID `bson:"_id"`

	Login        string `bson:"login"`
	PasswordHash string `bson:"password_hash"`
	Email        string `bson:"email"`
	Phone        string `bson:"phone"`

	Name       string `bson:"name"`
	Middlename string `bson:"middlename"`
	Surname    string `bson:"surname"`
	Age        int    `bson:"age"`

	City    string `bson:"city"`
	Address string `bson:"address"`
}
