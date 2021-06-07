package repository

import (
	"github.com/MuZaZaVr/account-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepository struct {
	db *mongo.Database
}

func (a AccountRepository) Create(account model.Account) (int, error) {
	panic("implement me")
}

func (a AccountRepository) FindByName(name string) (account model.Account, err error) {
	panic("implement me")
}

func (a AccountRepository) FindAllByCompanyId(id string) (int, []model.Account, error) {
	panic("implement me")
}

func (a AccountRepository) FindAllByUserID(id int) ([]model.Account, error) {
	panic("implement me")
}

func (a AccountRepository) UpdateName(newName string) (int, error) {
	panic("implement me")
}

func (a AccountRepository) UpdateDescription(newDescription string) (int, error) {
	panic("implement me")
}

func (a AccountRepository) UpdateCompanyId(newCompanyId int) (int, error) {
	panic("implement me")
}

func (a AccountRepository) Delete(id int) (bool, error) {
	panic("implement me")
}

func (a AccountRepository) IsExist(name string) (bool, error) {
	panic("implement me")
}

func NewAccountRepository(db *mongo.Database) *AccountRepository {
	return &AccountRepository{db: db}
}