package model

import (
	"github.com/MuZaZaVr/account-service/pkg/database/mongo"
)

// Credential represents mongo.Credential model for mongo
type Credential mongo.Credential

// Credentials represents a slice of Credential models for mongo
type Credentials []Credential

// CredentialsDTO represents a slice of DTO objects for mongo.Credential
type CredentialsDTO []CredentialDTO

// CredentialDTO represents DTO structure fot mongo.Credential
type CredentialDTO struct {
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`

	Name       string `json:"name"`
	Middlename string `json:"middlename"`
	Surname    string `json:"surname"`
	Age        int    `json:"age"`

	City    string `json:"city"`
	Address string `json:"address"`
}

// ConvertFromDTOToMongoModel func convert CredentialDTO object into mongo.Credential model
func (c CredentialDTO) ConvertFromDTOToMongoModel() *Credential {
	credential := Credential{
		Login:        c.Login,
		PasswordHash: c.PasswordHash,
		Email:        c.Email,
		Phone:        c.Phone,

		Name:       c.Name,
		Middlename: c.Middlename,
		Surname:    c.Surname,
		Age:        c.Age,

		City:    c.City,
		Address: c.Address,
	}

	return &credential
}

// ConvertFromMongoModelToDTO func convert mongo.Credential model into CredentialDTO object
func (c Credential) ConvertFromMongoModelToDTO() *CredentialDTO {
	credentialDTO := CredentialDTO{
		Login:        c.Login,
		PasswordHash: c.PasswordHash,
		Email:        c.Email,
		Phone:        c.Phone,
		Name:         c.Name,
		Middlename:   c.Middlename,
		Surname:      c.Surname,
		Age:          c.Age,
		City:         c.City,
		Address:      c.Address,
	}

	return &credentialDTO
}

// ConvertFewFromDTOToMongoModel func convert slice of CredentialDTO object into slice of mongo.Credential models
func (c CredentialsDTO) ConvertFewFromDTOToMongoModel() (Credentials, error) {
	var credentials Credentials

	for _, credential := range c {
		convertedCredential := credential.ConvertFromDTOToMongoModel()
		credentials = append(credentials, *convertedCredential)
	}

	return credentials, nil
}

// ConvertFewFromMongoModelToDTO func convert slice of mongo.Credential models into slice if CredentialDTO objects
func (c Credentials) ConvertFewFromMongoModelToDTO() CredentialsDTO {
	var credentialsDTO CredentialsDTO

	for _, credential := range c {
		credentialsDTO = append(credentialsDTO, *credential.ConvertFromMongoModelToDTO())
	}

	return credentialsDTO
}
