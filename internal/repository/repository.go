package repository

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repositories contain other repository interfaces
type Repositories struct {
	AccountRepository Account
	CompanyRepository Company
}

// NewRepositories is a constructor for Repositories struct
func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		AccountRepository: NewAccountRepository(db),
		CompanyRepository: NewCompanyRepository(db),
	}
}

// Company represents CRUD-repo for model.CompanyDTO
type Company interface {
	Create(ctx context.Context, company model.CompanyDTO) (string, error)

	FindByName(ctx context.Context, name string) (*model.CompanyDTO, error)
	FindByURL(ctx context.Context, url string) (*model.CompanyDTO, error)

	UpdateName(ctx context.Context, id string, newName string) (string, error)
	UpdateDescription(ctx context.Context, id string, newDescription string) (string, error)
	UpdateURL(ctx context.Context, id string, newUrl string) (string, error)

	Delete(ctx context.Context, id string) (string, error)

	IsExist(ctx context.Context, name string) (bool, error)
}

// Account represents CRUD-repo for model.AccountDTO
type Account interface {
	Create(ctx context.Context, account model.AccountDTO) (string, error)

	FindByName(ctx context.Context, name string) (*model.AccountDTO, error)
	FindAccountsByUserID(ctx context.Context, id int) ([]model.AccountDTO, error)

	FindAccountsByCredentialsLogin(ctx context.Context, credentialsLogin string) ([]model.AccountDTO, error)
	FindAccountsByCredentialsEmail(ctx context.Context, credentialsEmail string) ([]model.AccountDTO, error)
	FindAccountsByCredentialsPhone(ctx context.Context, credentialsPhone string) ([]model.AccountDTO, error)
	FindAccountsByCredentialsName(ctx context.Context, credentialsName string) ([]model.AccountDTO, error)
	FindAccountsByCredentialsMiddlename(ctx context.Context, credentialsMiddlename string) ([]model.AccountDTO, error)
	FindAccountsByCredentialsSurname(ctx context.Context, credentialsSurname string) ([]model.AccountDTO, error)
	FindAccountsByCredentialsAge(ctx context.Context, credentialsAge int) ([]model.AccountDTO, error)
	FindAccountsByCredentialsCity(ctx context.Context, credentialsCity string) ([]model.AccountDTO, error)
	FindAccountsByCredentialsAddress(ctx context.Context, credentialsAddress string) ([]model.AccountDTO, error)

	Update(ctx context.Context, id string, newAccount model.AccountDTO) (string, error)

	Delete(ctx context.Context, id string) (string, error)

	IsExist(ctx context.Context, name string) (bool, error)
}
