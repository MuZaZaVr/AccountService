package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`

	UserId      int         `bson:"user_id"`
	Company     Company     `bson:"company"`
	Credentials Credentials `bson:"credentials"`
}
