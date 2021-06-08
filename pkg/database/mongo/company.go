package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	URL         string             `bson:"URL"`
}
