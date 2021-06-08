package repository

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type CredentialRepository struct {
	db *mongo.Database
}

func (c CredentialRepository) Create(ctx context.Context, credential model.CredentialDTO) (string, error) {
	panic("implement me")
}

func (c CredentialRepository) FindByLogin(ctx context.Context, login string) ([]model.CredentialDTO, error) {
	panic("implement me")
}

func (c CredentialRepository) FindByEmail(ctx context.Context, email string) ([]model.CredentialDTO, error) {
	panic("implement me")
}

func (c CredentialRepository) FindByPhone(ctx context.Context, phone string) ([]model.CredentialDTO, error) {
	panic("implement me")
}

func (c CredentialRepository) FindByName(ctx context.Context, name string) ([]model.CredentialDTO, error) {
	panic("implement me")
}

func (c CredentialRepository) FindByMiddleName(ctx context.Context, middlename string) ([]model.CredentialDTO, error) {
	panic("implement me")
}

func (c CredentialRepository) FindBySurname(ctx context.Context, surname string) ([]model.CredentialDTO, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdateLogin(ctx context.Context, newLogin string) (string, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdateEmail(ctx context.Context, newEmail string) (string, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdatePhone(ctx context.Context, newPhone string) (string, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdateName(ctx context.Context, newName string) (string, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdateMiddleName(ctx context.Context, newMiddlename string) (string, error) {
	panic("implement me")
}

func (c CredentialRepository) UpdateSurname(ctx context.Context, newSurname string) (string, error) {
	panic("implement me")
}

func (c CredentialRepository) Delete(ctx context.Context, id string) (bool, error) {
	panic("implement me")
}

func NewCredentialRepository(db *mongo.Database) *CredentialRepository {
	return &CredentialRepository{db: db}
}

