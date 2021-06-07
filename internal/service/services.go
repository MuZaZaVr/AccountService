package service

import (
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
		AccountService:    NewAccountService(deps.Repos.AccountRepository, deps.TokenManager),
	}
}

type Company interface {
	Create(req request.CreateCompanyRequest) (string, error)

	FindByName(req request.FindCompanyByNameRequest) (model.Company, error)
	FindByURL(req request.FindCompanyByURLRequest) (model.Company, error)

	UpdateName(req request.UpdateCompanyNameRequest) (string, error)
	UpdateDescription(req request.UpdateCompanyDescriptionRequest) (string, error)
	UpdateURL(req request.UpdateCompanyURLRequest) (string, error)

	Delete(req request.DeleteCompanyRequest) (bool, error)

	IsExist(req request.IsCompanyExistRequest) (bool, error)
}

type Credential interface {
	Create(req request.CreateCredentialRequest) (int, error)

	FindByLogin(req request.FindCredentialByLoginRequest) ([]model.Credential, error)
	FindByEmail(req request.FindCredentialByEmailRequest) ([]model.Credential, error)
	FindByPhone(req request.FindCredentialByPhoneRequest) ([]model.Credential, error)
	FindByName(req request.FindCredentialByNameRequest) ([]model.Credential, error)
	FindByMiddleName(req request.FindCredentialByMiddlenameRequest) ([]model.Credential, error)
	FindBySurname(req request.FindCredentialBySurnameRequest) ([]model.Credential, error)

	UpdateLogin(req request.UpdateCredentialLoginRequest) (string, error)
	UpdateEmail(req request.UpdateCredentialEmailRequest) (string, error)
	UpdatePhone(req request.UpdateCredentialPhoneRequest) (string, error)
	UpdateName(req request.UpdateCredentialNameRequest) (string, error)
	UpdateMiddleName(req request.UpdateCredentialMiddlenameRequest) (string, error)
	UpdateSurname(req request.UpdateCredentialSurnameRequest) (string, error)

	Delete(req request.DeleteCredentialRequest) (bool, error)
}

type Account interface {
	Create(req request.CreateAccountRequest) (string, error)

	FindByCredentialLogin(req request.FindAccountsByCredentialLoginRequest) ([]repository.Account, error)
	FindByCredentialEmail(req request.FindAccountsByCredentialEmailRequest) ([]repository.Account, error)
	FindByCredentialPhone(req request.FindAccountsByCredentialPhoneRequest) ([]repository.Account, error)
	FindByCredentialName(req request.FindAccountsByCredentialNameRequest) ([]repository.Account, error)
	FindByCredentialMiddlename(req request.FindAccountsByCredentialMiddlenameRequest) ([]repository.Account, error)
	FindByCredentialSurname(req request.FindAccountsByCredentialSurnameRequest) ([]repository.Account, error)

	UpdateCredentialLogin(req request.UpdateAccountCredentialLoginRequest) (string, error)
	UpdateCredentialEmail(req request.UpdateAccountCredentialEmailRequest) (string, error)
	UpdateCredentialPhone(req request.UpdateAccountCredentialPhoneRequest) (string, error)
	UpdateCredentialName(req request.UpdateAccountCredentialNameRequest) (string, error)
	UpdateCredentialMiddlename(req request.UpdateAccountCredentialMiddlenameRequest) (string, error)
	UpdateCredentialSurname(req request.UpdateAccountCredentialSurnameRequest) (string, error)

	FindByName(req request.FindAccountByNameRequest) (model.Account, error)
	UpdateDescription(req request.UpdateCompanyDescriptionRequest)

	Delete(req request.DeleteAccountRequest) (bool, error)
}