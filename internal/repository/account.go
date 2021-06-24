package repository

import (
	"context"
	"errors"
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/internal/model/converter"
	"github.com/MuZaZaVr/account-service/pkg/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbAccountCollectionName = "account" //dbAccountCollectionName is a collection for mongo.Account struct in mongodb

type AccountRepository struct {
	db *mongoDriver.Collection
}

func NewAccountRepository(db *mongoDriver.Database) *AccountRepository {
	collection := db.Collection(dbAccountCollectionName)

	return &AccountRepository{db: collection}
}

// Create func used to create new Account and returns AccountID
func (a AccountRepository) Create(ctx context.Context, account model.AccountDTO) (string, error) {
	var convertedAccount mongo.Account
	err := converter.ConvertAccountFromDTOToMongo(account, &convertedAccount)
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
	if name == "" {
		return nil, errors.New("empty account name")
	}

	filterQuery := bson.M{"name": name}

	var mongoAccount mongo.Account
	err := a.db.FindOne(ctx, filterQuery).Decode(&mongoAccount)
	if err != nil {
		return nil, err
	}

	var convertedAccount model.AccountDTO
	converter.ConvertAccountFromMongoToDTO(mongoAccount, &convertedAccount)

	return &convertedAccount, nil
}

// FindAccountsByUserID func used to find all Account's by provided id and returns []model.AccountDTO
func (a AccountRepository) FindAccountsByUserID(ctx context.Context, id int) ([]model.AccountDTO, error) {
	var accounts []mongo.Account

	query := bson.M{"user_id": id}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	convertedAccounts, err := converter.ConvertFewAccountsFromMongoToDTO(accounts)
	if err != nil {
		return nil, err
	}

	return convertedAccounts, err
}

// FindAccountsByCredentialsLogin func used to find all Account's by provided credential login and returns []model.AccountDTO
func (a AccountRepository) FindAccountsByCredentialsLogin(ctx context.Context, credentialsLogin string) ([]model.AccountDTO, error) {
	var accounts []mongo.Account

	query := bson.M{"credentials.login": credentialsLogin}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	convertedAccounts, err := converter.ConvertFewAccountsFromMongoToDTO(accounts)
	if err != nil {
		return nil, err
	}

	return convertedAccounts, nil
}

// FindAccountsByCredentialsEmail func used to find all Account's by provided credential email and returns []model.AccountDTO
func (a AccountRepository) FindAccountsByCredentialsEmail(ctx context.Context, credentialsEmail string) ([]model.AccountDTO, error) {
	var accounts []mongo.Account

	query := bson.M{"credentials.email": credentialsEmail}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	convertedAccounts, err := converter.ConvertFewAccountsFromMongoToDTO(accounts)
	if err != nil {
		return nil, err
	}

	return convertedAccounts, nil
}

// FindAccountsByCredentialsPhone func used to find all Account's by provided credential phone and returns []model.AccountDTO
func (a AccountRepository) FindAccountsByCredentialsPhone(ctx context.Context, credentialsPhone string) ([]model.AccountDTO, error) {
	var accounts []mongo.Account

	query := bson.M{"credentials.phone": credentialsPhone}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	convertedAccounts, err := converter.ConvertFewAccountsFromMongoToDTO(accounts)
	if err != nil {
		return nil, err
	}

	return convertedAccounts, nil
}

// FindAccountsByCredentialsName func used to find all Account's by provided credential name and returns []model.AccountDTO
func (a AccountRepository) FindAccountsByCredentialsName(ctx context.Context, credentialsName string) ([]model.AccountDTO, error) {
	var accounts []mongo.Account

	query := bson.M{"credentials.name": credentialsName}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	convertedAccounts, err := converter.ConvertFewAccountsFromMongoToDTO(accounts)
	if err != nil {
		return nil, err
	}

	return convertedAccounts, nil
}

// FindAccountsByCredentialsMiddlename func used to find all Account's by provided credential middlename and returns []model.AccountDTO
func (a AccountRepository) FindAccountsByCredentialsMiddlename(ctx context.Context, credentialsMiddlename string) ([]model.AccountDTO, error) {
	var accounts []mongo.Account

	query := bson.M{"credentials.middlename": credentialsMiddlename}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	convertedAccounts, err := converter.ConvertFewAccountsFromMongoToDTO(accounts)
	if err != nil {
		return nil, err
	}

	return convertedAccounts, nil
}

// FindAccountsByCredentialsSurname func used to find all Account's by provided credential surname and returns []model.AccountDTO
func (a AccountRepository) FindAccountsByCredentialsSurname(ctx context.Context, credentialsSurname string) ([]model.AccountDTO, error) {
	var accounts []mongo.Account

	query := bson.M{"credentials.surname": credentialsSurname}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	convertedAccounts, err := converter.ConvertFewAccountsFromMongoToDTO(accounts)
	if err != nil {
		return nil, err
	}

	return convertedAccounts, nil
}

// FindAccountsByCredentialsAge func used to find all Account's by provided credential age and returns []model.AccountDTO
func (a AccountRepository) FindAccountsByCredentialsAge(ctx context.Context, credentialsAge int) ([]model.AccountDTO, error) {
	var accounts []mongo.Account

	query := bson.M{"credentials.age": credentialsAge}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	convertedAccounts, err := converter.ConvertFewAccountsFromMongoToDTO(accounts)
	if err != nil {
		return nil, err
	}

	return convertedAccounts, nil
}

// FindAccountsByCredentialsCity func used to find all Account's by provided credential city and returns []model.AccountDTO
func (a AccountRepository) FindAccountsByCredentialsCity(ctx context.Context, credentialsCity string) ([]model.AccountDTO, error) {
	var accounts []mongo.Account

	query := bson.M{"credentials.city": credentialsCity}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	convertedAccounts, err := converter.ConvertFewAccountsFromMongoToDTO(accounts)
	if err != nil {
		return nil, err
	}

	return convertedAccounts, nil
}

// FindAccountsByCredentialsAddress func used to find all Account's by provided credential address and returns []model.AccountDTO
func (a AccountRepository) FindAccountsByCredentialsAddress(ctx context.Context, credentialsAddress string) ([]model.AccountDTO, error) {
	var accounts []mongo.Account

	query := bson.M{"credentials.address": credentialsAddress}

	result, err := a.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &accounts)
	if err != nil {
		return nil, err
	}

	convertedAccounts, err := converter.ConvertFewAccountsFromMongoToDTO(accounts)
	if err != nil {
		return nil, err
	}

	return convertedAccounts, nil
}

// Update func used to update Account's name and returns updated AccountID
func (a AccountRepository) Update(ctx context.Context, id string, newAccount model.AccountDTO) (string, error) {
	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedID}
	updateQuery := bson.M{"$set": newAccount}

	var account mongo.Account
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

	var company mongo.Account
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

