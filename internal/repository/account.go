package repository

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbAccountCollectionName = "account" //dbAccountCollectionName is a collection for mongo.Account struct in mongodb

type AccountRepository struct {
	db *mongo.Collection
}

func NewAccountRepository(db *mongo.Database) *AccountRepository {
	collection := db.Collection(dbAccountCollectionName)

	return &AccountRepository{db: collection}
}

// Create func used to create new Account and returns AccountID
func (a AccountRepository) Create(ctx context.Context, account model.AccountDTO) (string, error) {
	convertedAccount, err := account.ConvertFromDTOToMongoModel()
	if err != nil {
		return "", err
	}

	result, err := a.db.InsertOne(ctx, convertedAccount)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// FindByName func used to find Account by name and returns model.AccountDTO
func (a AccountRepository) FindByName(ctx context.Context, name string) (*model.AccountDTO, error) {
	filterQuery := bson.M{"name": name}

	var account model.Account
	err := a.db.FindOne(ctx, filterQuery).Decode(&account)
	if err != nil {
		return nil, nil
	}

	convertedAccount := account.ConvertFromMongoModelToDTO()

	return convertedAccount, nil
}

// FindByCredentialID func used to find Account by Credential ID and returns model.AccountDTO
func (a AccountRepository) FindByCredentialID(ctx context.Context, id string) (*model.AccountDTO, error) {
	var account model.Account

	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = a.db.FindOne(ctx, convertedID).Decode(&account)
	if err != nil {
		return nil, err
	}

	return account.ConvertFromMongoModelToDTO(), nil
}

// FindCredentialIDByAccountID func used to find Credential ID by Account ID and returns model.AccountDTO
func (a AccountRepository) FindCredentialIDByAccountID(ctx context.Context, id string) (string, error) {
	var credential model.Credential

	filterQuery := bson.M{"_id": id}

	err := a.db.FindOne(ctx, filterQuery).Decode(&credential)
	if err != nil {
		return "", err
	}

	return credential.ID.Hex(), nil
}

// FindAllByCompanyID func used to find all Accounts with such companyID and returns slice of model.AccountDTO
func (a AccountRepository) FindAllByCompanyID(ctx context.Context, id string) ([]model.AccountDTO, error) {
	var accounts model.Accounts

	query := bson.M{"company_id": id}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts.ConvertFewFromMongoModelsToDTO(), nil
}

// FindAllByUserID func used to find all Accounts with such userID and returns slice of model.AccountDTO
func (a AccountRepository) FindAllByUserID(ctx context.Context, id int) ([]model.AccountDTO, error) {
	var accounts model.Accounts

	query := bson.M{"user_id": id}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts.ConvertFewFromMongoModelsToDTO(), err
}

// UpdateName func used to update Account's name and returns updated AccountID
func (a AccountRepository) UpdateName(ctx context.Context, id string, newName string) (string, error) {
	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedID}
	updateQuery := bson.D{{"$set", bson.D{{"name", newName}}}}

	var account model.Account
	err = a.db.FindOneAndUpdate(ctx, filterQuery, updateQuery).Decode(&account)
	if err != nil {
		return "", err
	}

	return account.ID.Hex(), err
}

// UpdateDescription func used to update Account's description and returns updated AccountID
func (a AccountRepository) UpdateDescription(ctx context.Context, id string, newDescription string) (string, error) {
	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedID}
	updateQuery := bson.D{{"$set", bson.D{{"description", newDescription}}}}

	var account model.Account
	err = a.db.FindOneAndUpdate(ctx, filterQuery, updateQuery).Decode(&account)
	if err != nil {
		return "", err
	}

	return account.ID.Hex(), err
}

// UpdateCompanyID func used to update Account's companyIS and returns updated AccountID
func (a AccountRepository) UpdateCompanyID(ctx context.Context, id string, newCompanyId string) (string, error) {
	convertedAccountID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	convertedCompanyID, err := primitive.ObjectIDFromHex(newCompanyId)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedAccountID}
	updateQuery := bson.D{{"$set", bson.D{{"company_id", convertedCompanyID}}}}

	var account model.Account
	err = a.db.FindOneAndUpdate(ctx, filterQuery, updateQuery).Decode(&account)
	if err != nil {
		return "", err
	}

	return account.ID.Hex(), err
}

// Delete func used to delete existed Account and returns deleted Account.ID
func (a AccountRepository) Delete(ctx context.Context, id string) (string, error) {
	opts := options.FindOneAndDelete().SetProjection(bson.D{{"_id", id}})

	convertedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedId}

	var company model.Company
	err = a.db.FindOneAndDelete(ctx, filterQuery, opts).Decode(&company)
	if err != nil {
		return "", err
	}

	return company.ID.Hex(), err
}

// IsExist func used to check account existence and returns true if company exist or false if not
func (a AccountRepository) IsExist(ctx context.Context, name string) (bool, error) {
	company, err := a.FindByName(ctx, name)
	if err != nil {
		return false, err
	}

	if company != nil {
		return true, nil
	}

	return false, nil
}

