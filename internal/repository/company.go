package repository

import (
	"github.com/MuZaZaVr/account-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepository struct {
	db *mongo.Database
}

func (c CompanyRepository) Create(company model.Company) (int, error) {
	panic("implement me")
}

func (c CompanyRepository) FindByName(name string) (model.Company, error) {
	panic("implement me")
}

func (c CompanyRepository) FindByURL(url string) (model.Company, error) {
	panic("implement me")
}

func (c CompanyRepository) UpdateName(newName string) (int, error) {
	panic("implement me")
}

func (c CompanyRepository) UpdateDescription(newDescription string) (int, error) {
	panic("implement me")
}

func (c CompanyRepository) UpdateURL(newUrl string) (int, error) {
	panic("implement me")
}

func (c CompanyRepository) Delete(id int) (bool, error) {
	panic("implement me")
}

func (c CompanyRepository) IsExist(name string) (bool, error) {
	panic("implement me")
}

func NewCompanyRepository(db *mongo.Database) *CompanyRepository {
	return &CompanyRepository{db: db}
}
