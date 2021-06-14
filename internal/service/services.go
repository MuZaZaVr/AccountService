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
	CredentialService Credential
	AccountService Account
}

type Deps struct {
	Repos        *repository.Repositories
	TokenManager auth.TokenManager
}

func NewServices(deps Deps) *Services {
	return &Services{
		CompanyService:    NewCompanyService(deps.Repos.CompanyRepository),
		CredentialService: NewCredentialService(deps.Repos.CredentialRepository),
		AccountService:    NewAccountService(deps.Repos, deps.TokenManager),
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

type Credential interface {
	Create(ctx context.Context, req request.CreateCredentialRequest) (string, error)

	FindByLogin(ctx context.Context, req request.FindCredentialByLoginRequest) ([]model.CredentialDTO, error)
	FindByEmail(ctx context.Context, req request.FindCredentialByEmailRequest) ([]model.CredentialDTO, error)
	FindByPhone(ctx context.Context, req request.FindCredentialByPhoneRequest) ([]model.CredentialDTO, error)
	FindByName(ctx context.Context, req request.FindCredentialByNameRequest) ([]model.CredentialDTO, error)
	FindByMiddleName(ctx context.Context, req request.FindCredentialByMiddlenameRequest) ([]model.CredentialDTO, error)
	FindBySurname(ctx context.Context, req request.FindCredentialBySurnameRequest) ([]model.CredentialDTO, error)

	UpdateLogin(ctx context.Context, req request.UpdateCredentialLoginRequest) (string, error)
	UpdateEmail(ctx context.Context, req request.UpdateCredentialEmailRequest) (string, error)
	UpdatePhone(ctx context.Context, req request.UpdateCredentialPhoneRequest) (string, error)
	UpdateName(ctx context.Context, req request.UpdateCredentialNameRequest) (string, error)
	UpdateMiddleName(ctx context.Context, req request.UpdateCredentialMiddlenameRequest) (string, error)
	UpdateSurname(ctx context.Context, req request.UpdateCredentialSurnameRequest) (string, error)

	Delete(ctx context.Context, req request.DeleteCredentialRequest) (string, error)
}

type Account interface {
	Create(ctx context.Context, req request.CreateAccountRequest) (string, error)

	FindByName(ctx context.Context, req request.FindAccountByNameRequest) (*model.AccountDTO, error)
	UpdateDescription(ctx context.Context, req request.UpdateCompanyDescriptionRequest) (string, error)

	Delete(ctx context.Context, req request.DeleteAccountRequest) (string, error)

	FindAllByCompanyID(ctx context.Context, req request.FindAllAccountsByCompanyIDRequest) ([]model.AccountDTO, error)
	FindAllByUserID(ctx context.Context, req request.FindAllAccountsByUserIDRequest) ([]model.AccountDTO, error)

	FindByCredentialLogin(ctx context.Context, req request.FindAccountsByCredentialLoginRequest) ([]model.AccountDTO, error)
	FindByCredentialEmail(ctx context.Context, req request.FindAccountsByCredentialEmailRequest) ([]model.AccountDTO, error)
	FindByCredentialPhone(ctx context.Context, req request.FindAccountsByCredentialPhoneRequest) ([]model.AccountDTO, error)
	FindByCredentialName(ctx context.Context, req request.FindAccountsByCredentialNameRequest) ([]model.AccountDTO, error)
	FindByCredentialMiddlename(ctx context.Context, req request.FindAccountsByCredentialMiddlenameRequest) ([]model.AccountDTO, error)
	FindByCredentialSurname(ctx context.Context, req request.FindAccountsByCredentialSurnameRequest) ([]model.AccountDTO, error)

	UpdateCredentialLogin(ctx context.Context, req request.UpdateAccountCredentialLoginRequest) (string, error)
	UpdateCredentialEmail(ctx context.Context, req request.UpdateAccountCredentialEmailRequest) (string, error)
	UpdateCredentialPhone(ctx context.Context, req request.UpdateAccountCredentialPhoneRequest) (string, error)
	UpdateCredentialName(ctx context.Context, req request.UpdateAccountCredentialNameRequest) (string, error)
	UpdateCredentialMiddlename(ctx context.Context, req request.UpdateAccountCredentialMiddlenameRequest) (string, error)
	UpdateCredentialSurname(ctx context.Context, req request.UpdateAccountCredentialSurnameRequest) (string, error)

}