package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id          primitive.ObjectID `bson:"id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`

	UserId       int                `bson:"user_id"`
	CompanyId    primitive.ObjectID `bson:"company_id"`
	CredentialId primitive.ObjectID `bson:"credential_id"`
}
