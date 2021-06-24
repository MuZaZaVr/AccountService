package mongo

import (
	"github.com/MuZaZaVr/account-service/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`

	UserId      int              `bson:"user_id"`
	Company     model.Company    `bson:"company"`
	Credentials model.Credential `bson:"credentials"`
}
