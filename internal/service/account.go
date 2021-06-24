package service

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/internal/model/request"
	"github.com/MuZaZaVr/account-service/internal/repository"
	"github.com/MuZaZaVr/account-service/pkg/auth"
	"github.com/pkg/errors"
)

// AccountService is a account service
type AccountService struct {
	accountRepo  repository.Account
	companyRepo  repository.Company
	tokenManager auth.TokenManager
}

// NewAccountService is a constructor for AccountService service
func NewAccountService(repositories *repository.Repositories, tknManager auth.TokenManager) *AccountService {
	return &AccountService{
		accountRepo:  repositories.AccountRepository,
		companyRepo:  repositories.CompanyRepository,
		tokenManager: tknManager,
	}
}

// Create func creates new account & returns id
func (a AccountService) Create(ctx context.Context, req request.CreateAccountRequest) (string, error) {
	accountModel := model.AccountDTO{
		Name:        req.Name,
		Description: req.Description,
		UserId:      req.UserID,
		Company:     req.Company,
		Credentials: req.Credential,
	}

	id, err := a.accountRepo.Create(ctx, accountModel)
	if err != nil {
		return "", errors.Wrap(err, "can't create new account")
	}

	return id, nil
}

// FindByName func find account by provided name & returns model.AccountDTO
func (a AccountService) FindByName(ctx context.Context, req request.FindAccountByNameRequest) (*model.AccountDTO, error) {
	account, err := a.accountRepo.FindByName(ctx, req.Name)
	if err != nil {
		return nil, errors.Wrap(err, "can't find account by name")
	}

	return account, nil
}

// FindAccountsByUserID func used to find all Account by User ID
func (a AccountService) FindAccountsByUserID(ctx context.Context, req request.FindAllAccountsByUserIDRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAccountsByUserID(ctx, req.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "can't find accounts by user ID")
	}

	return accounts, nil
}

// FindAllByUserID func used to find all Account by UserID
func (a AccountService) FindAllByUserID(ctx context.Context, req request.FindAllAccountsByUserIDRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAccountsByUserID(ctx, req.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "can't find accounts by user ID")
	}

	return accounts, nil
}

// FindAccountsByCredentialsLogin func used to find all Account by Credential login
func (a AccountService) FindAccountsByCredentialsLogin(ctx context.Context, req request.FindAccountsByCredentialLoginRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAccountsByCredentialsLogin(ctx, req.CredentialLogin)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// FindAccountsByCredentialsEmail func used to find all Account by Credential email
func (a AccountService) FindAccountsByCredentialsEmail(ctx context.Context, req request.FindAccountsByCredentialEmailRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAccountsByCredentialsEmail(ctx, req.CredentialEmail)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// FindAccountsByCredentialsPhone func used to find all Account by Credential phone
func (a AccountService) FindAccountsByCredentialsPhone(ctx context.Context, req request.FindAccountsByCredentialPhoneRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAccountsByCredentialsPhone(ctx, req.CredentialPhone)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// FindAccountsByCredentialsName func used to find all Account by Credential name
func (a AccountService) FindAccountsByCredentialsName(ctx context.Context, req request.FindAccountsByCredentialNameRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAccountsByCredentialsName(ctx, req.CredentialName)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// FindAccountsByCredentialsMiddlename func used to find all Account by Credential middleName
func (a AccountService) FindAccountsByCredentialsMiddlename(ctx context.Context, req request.FindAccountsByCredentialMiddlenameRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAccountsByCredentialsMiddlename(ctx, req.CredentialMiddlename)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// FindAccountsByCredentialsSurname func used to find all Account by Credential surname
func (a AccountService) FindAccountsByCredentialsSurname(ctx context.Context, req request.FindAccountsByCredentialSurnameRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAccountsByCredentialsSurname(ctx, req.CredentialSurname)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// FindAccountsByCredentialsAge func used to find all Account by Credential age
func (a AccountService) FindAccountsByCredentialsAge(ctx context.Context, req request.FindAccountsByCredentialAgeRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAccountsByCredentialsAge(ctx, req.CredentialAge)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// FindAccountsByCredentialsCity func used to find all Account by Credential city
func (a AccountService) FindAccountsByCredentialsCity(ctx context.Context, req request.FindAccountsByCredentialCityRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAccountsByCredentialsCity(ctx, req.CredentialCity)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// FindAccountsByCredentialsAddress func used to find all Account by Credential address
func (a AccountService) FindAccountsByCredentialsAddress(ctx context.Context, req request.FindAccountsByCredentialAddressRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAccountsByCredentialsAddress(ctx, req.CredentialAddress)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// Update func update account & returns model.Account ID
func (a AccountService) Update(ctx context.Context, req request.UpdateAccountRequest) (string, error) {
	id, err := a.accountRepo.Update(ctx, req.ID, req.UpdatedAccount)
	if err != nil {
		return "", errors.Wrap(err, "can't find account by name")
	}

	return id, nil
}

// Delete func delete Account & returns model.Account ID
func (a AccountService) Delete(ctx context.Context, req request.DeleteAccountRequest) (string, error) {
	id, err := a.accountRepo.Delete(ctx, req.ID)
	if err != nil {
		return "", errors.Wrap(err, "can't delete account")
	}

	return id, nil
}
