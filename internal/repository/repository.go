package repository

import (
	"github.com/MuZaZaVr/account-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	AccountRepository    Account
	CompanyRepository    Company
	CredentialRepository Credential
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		AccountRepository:    NewAccountRepository(db),
		CompanyRepository:    NewCompanyRepository(db),
		CredentialRepository: NewCredentialRepository(db),
	}
}

// Company represents CRUD-repo for model.Company
type Company interface {
	Create(company model.Company) (int, error)

	FindByName(name string) (model.Company, error)
	FindByURL(url string) (model.Company, error)

	UpdateName(newName string) (int, error)
	UpdateDescription(newDescription string) (int, error)
	UpdateURL(newUrl string) (int, error)

	Delete(id int) (bool, error)

	IsExist(name string) (bool, error)
}

// Credential represents CRUD-repo for model.Credential
type Credential interface {
	Create(credential model.Credential) (int, error)

	FindByLogin(login string) ([]model.Credential, error)
	FindByEmail(email string) ([]model.Credential, error)
	FindByPhone(phone string) ([]model.Credential, error)
	FindByName(name string) ([]model.Credential, error)
	FindByMiddleName(middlename string) ([]model.Credential, error)
	FindBySurname(surname string) ([]model.Credential, error)

	UpdateLogin(newLogin string) (int, error)
	UpdateEmail(newEmail string) (int, error)
	UpdatePhone(newPhone string) (int, error)
	UpdateName(newName string) (int, error)
	UpdateMiddleName(newMiddlename string) (int, error)
	UpdateSurname(newSurname string) (int, error)

	Delete(in int) (bool, error)
}

// Account represents CRUD-repo for model.Account
type Account interface {
	Create( model.Account) (int, error)

	FindByName(name string) (account model.Account, err error)
	FindAllByCompanyId(id string) (int, []model.Account, error)
	FindAllByUserID(id int) ([]model.Account, error)

	UpdateName(newName string) (int, error)
	UpdateDescription(newDescription string) (int, error)
	UpdateCompanyId(newCompanyId int) (int, error)

	Delete(id int) (bool, error)

	IsExist(name string) (bool, error)
}
