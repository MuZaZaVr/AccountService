package repository

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/internal/model/converter"
	"github.com/MuZaZaVr/account-service/pkg/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbCompanyCollectionName = "company" //dbCompanyCollectionName is a collection for mongo.Company struct in mongodb

// CompanyRepository represent struct for DB operations with Company
type CompanyRepository struct {
	db *mongoDriver.Collection
}

// NewCompanyRepository func is a constructor for CompanyRepository
func NewCompanyRepository(db *mongoDriver.Database) *CompanyRepository {
	collection := db.Collection(dbCompanyCollectionName)

	return &CompanyRepository{db: collection}
}

// Create func used to create new Company and returns CompanyID
func (c CompanyRepository) Create(ctx context.Context, company model.CompanyDTO) (string, error) {
	var convertedCompany mongo.Company
	err := converter.ConvertCompanyFromDTOToMongo(company, &convertedCompany)
	if err != nil {
		return "", err
	}

	result, err := c.db.InsertOne(ctx, convertedCompany)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// FindByName func used to find Company and returns model.CompanyDTO
func (c CompanyRepository) FindByName(ctx context.Context, name string) (*model.CompanyDTO, error) {
	var mongoCompany mongo.Company

	query := bson.M{"name": name}
	err := c.db.FindOne(ctx, query).Decode(&mongoCompany)
	if err != nil {
		return nil, err
	}

	var convertedCompany model.CompanyDTO
	converter.ConvertCompanyFromMongoToDTO(mongoCompany, &convertedCompany)

	return &convertedCompany, nil
}

// FindByURL func used to find Company and returns model.CompanyDTO
func (c CompanyRepository) FindByURL(ctx context.Context, url string) (*model.CompanyDTO, error) {
	var mongoCompany mongo.Company

	query := bson.M{"URL": url}
	err := c.db.FindOne(ctx, query).Decode(&mongoCompany)
	if err != nil {
		return nil, err
	}

	var convertedCompany model.CompanyDTO
	converter.ConvertCompanyFromMongoToDTO(mongoCompany, &convertedCompany)

	return &convertedCompany, nil
}

// UpdateName func used to update Company's name and returns updated CompanyID
func (c CompanyRepository) UpdateName(ctx context.Context, id string, newName string) (string, error) {
	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filterQuery := bson.M{"_id": convertedID}
	queryQuery := bson.D{{"$set", bson.D{{"name", newName}}}}

	var updatedCompany mongo.Company
	err = c.db.FindOneAndUpdate(ctx, filterQuery, queryQuery).Decode(&updatedCompany)
	if err != nil {
		return "", err
	}

	return updatedCompany.ID.Hex(), nil
}

// UpdateDescription func used to update Company's description and returns updated CompanyID
func (c CompanyRepository) UpdateDescription(ctx context.Context, id string, newDescription string) (string, error) {
	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	filerQuery := bson.M{"_id": convertedID}
	updateQuery := bson.D{{"$set", bson.D{{"description", newDescription}}}}

	var updatedCompany mongo.Company
	err = c.db.FindOneAndUpdate(ctx, filerQuery, updateQuery).Decode(&updatedCompany)
	if err != nil {
		return "", err
	}

	return updatedCompany.ID.Hex(), nil
}

// UpdateURL func used to update Company's URL and returns updated CompanyID
func (c CompanyRepository) UpdateURL(ctx context.Context, id string, newUrl string) (string, error) {
	convertedId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		 return "", err
	}

	filterQuery := bson.M{"_id": convertedId}
	updateQuery := bson.D{{"$set", bson.D{{"URL", newUrl}}}}

	var updatedCompany mongo.Company
	err = c.db.FindOneAndUpdate(ctx, filterQuery, updateQuery).Decode(&updatedCompany)
	if err != nil {
		return "", err
	}

	return updatedCompany.ID.Hex(), nil
}

// Delete func used to delete existed Company and returns deleted Company.ID
func (c CompanyRepository) Delete(ctx context.Context, id string) (string, error) {
	opts := options.FindOneAndDelete().SetProjection(bson.D{{"_id", 1}})

	convertedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	query := bson.M{"_id": convertedID}
	var deletedDocument mongo.Company
	err = c.db.FindOneAndDelete(ctx, query, opts).Decode(&deletedDocument)
	if err != nil {
		return "", err
	}

	return deletedDocument.ID.Hex(), nil
}

// IsExist func used to check company existence and returns true if company exist or false if not
func (c CompanyRepository) IsExist(ctx context.Context, name string) (bool, error) {
	company, err := c.FindByName(ctx, name)
	if err != nil {
		return false, err
	}

	if company != nil {
		return true, nil
	}

	return false, nil
}
