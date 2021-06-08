package repository

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepository struct {
	db *mongo.Database
}

func (a AccountRepository) Create(ctx context.Context, account model.AccountDTO) (string, error) {
	panic("implement me")
}

func (a AccountRepository) FindByName(ctx context.Context, name string) (account model.AccountDTO, err error) {
	panic("implement me")
}

func (a AccountRepository) FindAllByCompanyId(ctx context.Context, id string) ([]model.AccountDTO, error) {
	panic("implement me")
}

func (a AccountRepository) FindAllByUserID(ctx context.Context, id int) ([]model.AccountDTO, error) {
	panic("implement me")
}

func (a AccountRepository) UpdateName(ctx context.Context, newName string) (string, error) {
	panic("implement me")
}

func (a AccountRepository) UpdateDescription(ctx context.Context, newDescription string) (string, error) {
	panic("implement me")
}

func (a AccountRepository) UpdateCompanyId(ctx context.Context, newCompanyId int) (string, error) {
	panic("implement me")
}

func (a AccountRepository) Delete(ctx context.Context, id string) (bool, error) {
	panic("implement me")
}

func (a AccountRepository) IsExist(ctx context.Context, name string) (bool, error) {
	panic("implement me")
}

func NewAccountRepository(db *mongo.Database) *AccountRepository {
	return &AccountRepository{db: db}
}

