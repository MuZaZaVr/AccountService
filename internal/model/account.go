package model

import (
	"github.com/MuZaZaVr/account-service/pkg/database/mongo"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Account represents mongo.Account model for mongo
type Account mongo.Account

// Accounts represents a slice of Account models for mongo
type Accounts []Account

// AccountsDTO represents a slice of DTO objects for mongo.Account
type AccountsDTO []AccountDTO

// AccountDTO represents DTO structure for mongo.Account
type AccountDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	UserId       int    `json:"user_id"`
	CompanyID    string `json:"company_id"`
	CredentialID string `json:"credential_id"`
}

// ConvertFromDTOToMongoModel func convert AccountDTO object to mongo.Account model
func (a *AccountDTO) ConvertFromDTOToMongoModel() (*Account, error) {
	var err error
	account := Account{
		Name:        a.Name,
		Description: a.Description,

		UserId: a.UserId,
	}

	if a.ID != "" {
		account.ID, err = primitive.ObjectIDFromHex(a.ID)
		if err != nil {
			return nil, errors.Wrap(err, "invalid account id")
		}
	}

	if a.CompanyID != "" {
		account.CompanyID, err = primitive.ObjectIDFromHex(a.CompanyID)
		if err != nil {
			return nil, errors.Wrap(err, "invalid company id")
		}
	}

	if a.CredentialID != "" {
		account.CredentialID, err = primitive.ObjectIDFromHex(a.CredentialID)
		if err != nil {
			return nil, errors.Wrap(err, "invalid credential id")
		}
	}

	return &account, err
}

// ConvertFromMongoModelToDTO func convert mongo.Account model to CompanyDTO object
func (a *Account) ConvertFromMongoModelToDTO() *AccountDTO {
	account := AccountDTO{
		ID:           a.ID.Hex(),
		Name:         a.Name,
		Description:  a.Description,
		UserId:       a.UserId,
		CompanyID:    a.CompanyID.Hex(),
		CredentialID: a.CredentialID.Hex(),
	}

	return &account
}

// ConvertFewFromDTOToMongoModels func convert a slice of AccountDTO objects into a slice of mongo.Account models
func (a AccountsDTO) ConvertFewFromDTOToMongoModels() (*Accounts, error) {
	var mongoAccounts Accounts
	for _, accountDTO := range a {
		mongoAccount, err := accountDTO.ConvertFromDTOToMongoModel()
		if err != nil {
			return nil, errors.Wrap(err, "can't convert from DTO objects to Mongo models")
		}
		mongoAccounts = append(mongoAccounts, *mongoAccount)
	}

	return &mongoAccounts, nil
}

// ConvertFewFromMongoModelsToDTO func convert a slice of mongo.Account models into a slice of AccountDTO objects
func (a Accounts) ConvertFewFromMongoModelsToDTO() *AccountsDTO {
	var accountsDTO AccountsDTO
	for _, mongoAccount := range a{
		accountDTO := mongoAccount.ConvertFromMongoModelToDTO()
		accountsDTO = append(accountsDTO, *accountDTO)
	}

	return &accountsDTO
}