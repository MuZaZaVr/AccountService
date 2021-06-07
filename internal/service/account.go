package service

import (
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/internal/model/request"
	"github.com/MuZaZaVr/account-service/internal/repository"
	"github.com/MuZaZaVr/account-service/pkg/auth"
)

type AccountService struct {
	accountRepo repository.Account
	tokenManager auth.TokenManager
}

func (a AccountService) Create(req request.CreateAccountRequest) (string, error) {

	panic("implement me")
}

func (a AccountService) FindByCredentialLogin(req request.FindAccountsByCredentialLoginRequest) ([]repository.Account, error) {
	panic("implement me")
}

func (a AccountService) FindByCredentialEmail(req request.FindAccountsByCredentialEmailRequest) ([]repository.Account, error) {
	panic("implement me")
}

func (a AccountService) FindByCredentialPhone(req request.FindAccountsByCredentialPhoneRequest) ([]repository.Account, error) {
	panic("implement me")
}

func (a AccountService) FindByCredentialName(req request.FindAccountsByCredentialNameRequest) ([]repository.Account, error) {
	panic("implement me")
}

func (a AccountService) FindByCredentialMiddlename(req request.FindAccountsByCredentialMiddlenameRequest) ([]repository.Account, error) {
	panic("implement me")
}

func (a AccountService) FindByCredentialSurname(req request.FindAccountsByCredentialSurnameRequest) ([]repository.Account, error) {
	panic("implement me")
}

func (a AccountService) UpdateCredentialLogin(req request.UpdateAccountCredentialLoginRequest) (string, error) {
	panic("implement me")
}

func (a AccountService) UpdateCredentialEmail(req request.UpdateAccountCredentialEmailRequest) (string, error) {
	panic("implement me")
}

func (a AccountService) UpdateCredentialPhone(req request.UpdateAccountCredentialPhoneRequest) (string, error) {
	panic("implement me")
}

func (a AccountService) UpdateCredentialName(req request.UpdateAccountCredentialNameRequest) (string, error) {
	panic("implement me")
}

func (a AccountService) UpdateCredentialMiddlename(req request.UpdateAccountCredentialMiddlenameRequest) (string, error) {
	panic("implement me")
}

func (a AccountService) UpdateCredentialSurname(req request.UpdateAccountCredentialSurnameRequest) (string, error) {
	panic("implement me")
}

func (a AccountService) FindByName(req request.FindAccountByNameRequest) (model.Account, error) {
	panic("implement me")
}

func (a AccountService) UpdateDescription(req request.UpdateCompanyDescriptionRequest) {
	panic("implement me")
}

func (a AccountService) Delete(req request.DeleteAccountRequest) (bool, error) {
	panic("implement me")
}

func NewAccountService(accountRepo repository.Account, tknManager auth.TokenManager) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
		tokenManager: tknManager,
	}
}

