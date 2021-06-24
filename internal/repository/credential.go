package repository

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbCredentialCollectionName = "credential"

// CredentialRepository represent struct for DB operations with Credential
type CredentialRepository struct {
	db *mongo.Collection
}

// NewCredentialRepository func is a constructor for CredentialRepository
func NewCredentialRepository(db *mongo.Database) *CredentialRepository {
	collection := db.Collection(dbCredentialCollectionName)

	return &CredentialRepository{db: collection}
}

// Create func used to create new Credential and returns CredentialID
func (c CredentialRepository) Create(ctx context.Context, credential model.CredentialDTO) (string, error) {
	credentialToSave, err := credential.ConvertFromDTOToMongoModel()
	if err != nil {
		return "", err
	}

	result, err := c.db.InsertOne(ctx, credentialToSave)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// FindByLogin func used to find all Credentials with such login and returns slice of model.CredentialDTO
func (c CredentialRepository) FindByLogin(ctx context.Context, login string) ([]model.CredentialDTO, error) {

	filterQuery := bson.M{"login": login}

	result, err := c.db.Find(ctx, filterQuery)
	if err != nil {
		return nil, err
	}

	var foundedCredentials model.Credentials
	err = result.All(ctx, &foundedCredentials)
	if err != nil {
		return nil, err
	}

	return foundedCredentials.ConvertFewFromMongoModelToDTO(), nil
}

// FindByEmail func used to find all Credentials with such login and returns slice of model.CredentialDTO
func (c CredentialRepository) FindByEmail(ctx context.Context, email string) ([]model.CredentialDTO, error) {

	filerQuery := bson.M{"email": email}

	result, err := c.db.Find(ctx, filerQuery)
	if err != nil {
		return nil, err
	}

	var foundedCredentials model.Credentials
	err = result.All(ctx, &foundedCredentials)
	if err != nil {
		return nil, err
	}

	return foundedCredentials.ConvertFewFromMongoModelToDTO(), err
}

// FindByPhone func used to find all Credentials with such phone and returns slice of model.CredentialDTO
func (c CredentialRepository) FindByPhone(ctx context.Context, phone string) ([]model.CredentialDTO, error) {

	filterQuery := bson.M{"phone": phone}

	result, err := c.db.Find(ctx, filterQuery)
	if err != nil {
		return nil, err
	}

	var foundedCredentials model.Credentials
	err = result.All(ctx, &foundedCredentials)
	if err != nil {
		return nil, err
	}

	return foundedCredentials.ConvertFewFromMongoModelToDTO(), err
}

// FindByName func used to find all Credentials with such name and returns slice of model.CredentialDTO
func (c CredentialRepository) FindByName(ctx context.Context, name string) ([]model.CredentialDTO, error) {

	filterQuery := bson.M{"name": name}

	result, err := c.db.Find(ctx, filterQuery)
	if err != nil {
		return nil, err
	}

	var foundedCredentials model.Credentials
	err = result.All(ctx, &foundedCredentials)
	if err != nil {
		return nil, err
	}

	return foundedCredentials.ConvertFewFromMongoModelToDTO(), err
}

// FindByMiddleName func used to find all Credentials with such middlename and returns slice of model.CredentialDTO
func (c CredentialRepository) FindByMiddleName(ctx context.Context, middlename string) ([]model.CredentialDTO, error) {

	filterQuery := bson.M{"middlename": middlename}

	result, err := c.db.Find(ctx, filterQuery)
	if err != nil {
		return nil, err
	}

	var foundedCredentials model.Credentials
	err = result.All(ctx, &foundedCredentials)
	if err != nil {
		return nil, err
	}

	return foundedCredentials.ConvertFewFromMongoModelToDTO(), err
}

// FindBySurname func used to find all Credentials with such surname and returns slice of model.CredentialDTO
func (c CredentialRepository) FindBySurname(ctx context.Context, surname string) ([]model.CredentialDTO, error) {

	filterQuery := bson.M{"surname": surname}

	result, err := c.db.Find(ctx, filterQuery)
	if err != nil {
		return nil, err
	}

	var foundedCredentials model.Credentials
	err = result.All(ctx, &foundedCredentials)
	if err != nil {
		return nil, err
	}

	return foundedCredentials.ConvertFewFromMongoModelToDTO(), err
}

// UpdateLogin func used to update Credential's login and returns updated CredentialID
func (c CredentialRepository) UpdateLogin(ctx context.Context, id string, newLogin string) (string, error) {

	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedID}
	updateQuery := bson.D{{"$set", bson.D{{"login", newLogin}}}}

	var updatedCredential model.Credential
	err = c.db.FindOneAndUpdate(ctx, filterQuery, updateQuery).Decode(&updatedCredential)
	if err != nil {
		return "", err
	}

	return updatedCredential.ID.Hex(), nil
}

// UpdateEmail func used to update Credential's login and returns updated CredentialID
func (c CredentialRepository) UpdateEmail(ctx context.Context, id string, newEmail string) (string, error) {

	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedID}
	updateQuery := bson.D{{"$set", bson.D{{"login", newEmail}}}}

	var updatedCredential model.Credential
	err = c.db.FindOneAndUpdate(ctx, filterQuery, updateQuery).Decode(&updatedCredential)
	if err != nil {
		return "", err
	}

	return updatedCredential.ID.Hex(), nil
}

// UpdatePhone func used to update Credential's phone and returns updated CredentialID
func (c CredentialRepository) UpdatePhone(ctx context.Context, id string, newPhone string) (string, error) {

	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedID}
	updateQuery := bson.D{{"$set", bson.D{{"phone", newPhone}}}}

	var updatedCredential model.Credential
	err = c.db.FindOneAndUpdate(ctx, filterQuery, updateQuery).Decode(&updatedCredential)
	if err != nil {
		return "", err
	}

	return updatedCredential.ID.Hex(), nil
}

// UpdateName func used to update Credential's name and returns updated CredentialID
func (c CredentialRepository) UpdateName(ctx context.Context, id string, newName string) (string, error) {

	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedID}
	updateQuery := bson.D{{"$set", bson.D{{"name", newName}}}}

	var updatedCredential model.Credential
	err = c.db.FindOneAndUpdate(ctx, filterQuery, updateQuery).Decode(&updatedCredential)
	if err != nil {
		return "", err
	}

	return updatedCredential.ID.Hex(), nil
}

// UpdateMiddleName func used to update Credential's middlename and returns updated CredentialID
func (c CredentialRepository) UpdateMiddleName(ctx context.Context, id string, newMiddlename string) (string, error) {

	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedID}
	updateQuery := bson.D{{"$set", bson.D{{"middlename", newMiddlename}}}}

	var updatedCredential model.Credential

	err = c.db.FindOneAndUpdate(ctx, filterQuery, updateQuery).Decode(&updatedCredential)
	if err != nil {
		return "", err
	}

	return updatedCredential.ID.Hex(), nil
}

// UpdateSurname func used to update Credential's surname and returns updated CredentialID
func (c CredentialRepository) UpdateSurname(ctx context.Context, id string, newSurname string) (string, error) {

	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedID}
	updateQuery := bson.D{{"$set", bson.D{{"surname", newSurname}}}}

	var updatedCredential model.Credential
	err = c.db.FindOneAndUpdate(ctx, filterQuery, updateQuery).Decode(&updatedCredential)
	if err != nil {
		return "", err
	}

	return updatedCredential.ID.Hex(), nil
}

// Delete func used to delete existed Credential and returns deleted Credential.responseID
func (c CredentialRepository) Delete(ctx context.Context, id string) (string, error) {
	opts := options.FindOneAndDelete().SetProjection(bson.D{{"_id", 1}})

	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	deleteQuery := bson.M{"_id": convertedID}

	var updatedCredential model.Credential
	err = c.db.FindOneAndDelete(ctx, deleteQuery, opts).Decode(&updatedCredential)
	if err != nil {
		return "", err
	}

	return updatedCredential.ID.Hex(), nil
}
