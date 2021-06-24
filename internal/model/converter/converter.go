package converter

import (
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/pkg/database/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertAccountFromDTOToMongo(dto model.AccountDTO, target *mongo.Account) error {
	if dto.ID != "" {
		convertedAccountID, err := primitive.ObjectIDFromHex(dto.ID)
		if err != nil {
			return err
		}
		target.ID = convertedAccountID
	}

	var convertedCompany mongo.Company
	err := ConvertCompanyFromDTOToMongo(dto.Company, &convertedCompany)
	if err != nil {
		return err
	}

	var convertedCredentials mongo.Credentials
	ConvertCredentialsFromDTOToMongo(dto.Credentials, &convertedCredentials)


	target.Name = dto.Name
	target.Description = dto.Description
	target.UserId = dto.UserId
	target.Company = convertedCompany
	target.Credentials = convertedCredentials

	return nil
}

func ConvertAccountFromMongoToDTO(object mongo.Account, target *model.AccountDTO) {
	convertedAccountID := object.ID.Hex()

	var convertedCompany model.CompanyDTO
	ConvertCompanyFromMongoToDTO(object.Company, &convertedCompany)

	var convertedCredentials model.CredentialsDTO
	ConvertCredentialsFromMongoToDTO(object.Credentials, &convertedCredentials)

	target.ID = convertedAccountID
	target.Name = object.Name
	target.Description = object.Description
	target.UserId = object.UserId
	target.Company = convertedCompany
	target.Credentials = convertedCredentials
}

func ConvertFewAccountsFromDTOToMongo(dtos []model.AccountDTO) ([]mongo.Account, error) {
	var target []mongo.Account
	var convertedAccount mongo.Account
	for _, dto := range dtos {
		err := ConvertAccountFromDTOToMongo(dto, &convertedAccount)
		if err != nil {
			return nil, err
		}

		target = append(target, convertedAccount)
	}

	return target, nil
}

func ConvertFewAccountsFromMongoToDTO(mongos []mongo.Account) ([]model.AccountDTO, error) {
	var target []model.AccountDTO
	var convertedAccount model.AccountDTO
	for _, dto := range mongos {
		ConvertAccountFromMongoToDTO(dto, &convertedAccount)
		target = append(target, convertedAccount)
	}

	return target, nil
}

func ConvertCompanyFromDTOToMongo(dto model.CompanyDTO, target *mongo.Company) error {
	if dto.ID != "" {
		convertedID, err := primitive.ObjectIDFromHex(dto.ID)
		if err != nil {
			return err
		}
		target.ID = convertedID
	}
	target.Name = dto.Name
	target.Description = dto.Description
	target.URL = dto.URL

	return nil
}

func ConvertCompanyFromMongoToDTO(object mongo.Company, target *model.CompanyDTO) {
	convertedID := object.ID.Hex()

	target.ID = convertedID
	target.Name = object.Name
	target.Description = object.Description
	target.URL = object.URL
}

func ConvertCredentialsFromDTOToMongo(dto model.CredentialsDTO, target *mongo.Credentials) {
	target.Name = dto.Name
	target.Email = dto.Email
	target.Login = dto.Login
	target.PasswordHash = dto.PasswordHash
	target.Phone = dto.Phone
	target.Name = dto.Name
	target.Surname = dto.Surname
	target.Middlename = dto.Middlename
	target.City = dto.City
	target.Address = dto.Address
}

func ConvertCredentialsFromMongoToDTO(object mongo.Credentials, target *model.CredentialsDTO) {
	target.Name = object.Name
	target.Email = object.Email
	target.Login = object.Login
	target.PasswordHash = object.PasswordHash
	target.Phone = object.Phone
	target.Name = object.Name
	target.Surname = object.Surname
	target.Middlename = object.Middlename
	target.City = object.City
	target.Address = object.Address
}
