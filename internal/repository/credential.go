package repository

import (
	"github.com/MuZaZaVr/account-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type CredentialRepository struct {
	db *mongo.Database
}

func (c CredentialRepository) Create(credential model.Credential) (int, error) {
	panic("implement me")
}

func (c CredentialRepository) FindByLogin(login string) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialRepository) FindByEmail(email string) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialRepository) FindByPhone(phone string) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialRepository) FindByName(name string) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialRepository) FindByMiddleName(middlename string) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialRepository) FindBySurname(surname string) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdateLogin(newLogin string) (int, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdateEmail(newEmail string) (int, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdatePhone(newPhone string) (int, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdateName(newName string) (int, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdateMiddleName(newMiddlename string) (int, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdateSurname(newSurname string) (int, error) {
	panic("implement me")
}

func (c CredentialRepository) Delete(in int) (bool, error) {
	panic("implement me")
}

func NewCredentialRepository(db *mongo.Database) *CredentialRepository {
	return &CredentialRepository{db: db}
}
