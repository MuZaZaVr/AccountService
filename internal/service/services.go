package service

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/internal/model/request"
	"github.com/MuZaZaVr/account-service/internal/repository"
	"github.com/MuZaZaVr/account-service/pkg/auth"
)

type Services struct {
	CompanyService Company
	AccountService Account
}

type Deps struct {
	Repos        *repository.Repositories
	TokenManager auth.TokenManager
}

func NewServices(deps Deps) *Services {
	return &Services{
		CompanyService: NewCompanyService(deps.Repos.CompanyRepository),
		AccountService: NewAccountService(deps.Repos, deps.TokenManager),
	}
}

type Company interface {
	Create(ctx context.Context, req request.CreateCompanyRequest) (string, error)

	FindByName(ctx context.Context, req request.FindCompanyByNameRequest) (*model.CompanyDTO, error)
	FindByURL(ctx context.Context, req request.FindCompanyByURLRequest) (*model.CompanyDTO, error)

	UpdateName(ctx context.Context, req request.UpdateCompanyNameRequest) (string, error)
	UpdateDescription(ctx context.Context, req request.UpdateCompanyDescriptionRequest) (string, error)
	UpdateURL(ctx context.Context, req request.UpdateCompanyURLRequest) (string, error)

	Delete(ctx context.Context, req request.DeleteCompanyRequest) (string, error)

	IsExist(ctx context.Context, req request.IsCompanyExistRequest) (bool, error)
}

type Account interface {
	Create(ctx context.Context, req request.CreateAccountRequest) (string, error)

	FindByName(ctx context.Context, req request.FindAccountByNameRequest) (*model.AccountDTO, error)
	FindAccountsByUserID(ctx context.Context, req request.FindAllAccountsByUserIDRequest) ([]model.AccountDTO, error)

	FindAccountsByCredentialsLogin(ctx context.Context, req request.FindAccountsByCredentialLoginRequest) ([]model.AccountDTO, error)
	FindAccountsByCredentialsEmail(ctx context.Context, req request.FindAccountsByCredentialEmailRequest) ([]model.AccountDTO, error)
	FindAccountsByCredentialsPhone(ctx context.Context, req request.FindAccountsByCredentialPhoneRequest) ([]model.AccountDTO, error)
	FindAccountsByCredentialsName(ctx context.Context, req request.FindAccountsByCredentialNameRequest) ([]model.AccountDTO, error)
	FindAccountsByCredentialsMiddlename(ctx context.Context, req request.FindAccountsByCredentialMiddlenameRequest) ([]model.AccountDTO, error)
	FindAccountsByCredentialsSurname(ctx context.Context, req request.FindAccountsByCredentialSurnameRequest) ([]model.AccountDTO, error)
	FindAccountsByCredentialsAge(ctx context.Context, req request.FindAccountsByCredentialAgeRequest) ([]model.AccountDTO, error)
	FindAccountsByCredentialsCity(ctx context.Context, req request.FindAccountsByCredentialCityRequest) ([]model.AccountDTO, error)
	FindAccountsByCredentialsAddress(ctx context.Context, req request.FindAccountsByCredentialAddressRequest) ([]model.AccountDTO, error)

	Update(ctx context.Context, req request.UpdateAccountRequest) (string, error)

	Delete(ctx context.Context, req request.DeleteAccountRequest) (string, error)

}
