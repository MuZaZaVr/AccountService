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
	accountRepo    repository.Account
	companyRepo    repository.Company
	credentialRepo repository.Credential
	tokenManager   auth.TokenManager
}

// NewAccountService is a constructor for AccountService service
func NewAccountService(repositories *repository.Repositories, tknManager auth.TokenManager) *AccountService {
	return &AccountService{
		accountRepo:    repositories.AccountRepository,
		companyRepo:    repositories.CompanyRepository,
		credentialRepo: repositories.CredentialRepository,
		tokenManager:   tknManager,
	}
}
// Create func creates new account & returns id
func (a AccountService) Create(ctx context.Context, req request.CreateAccountRequest) (string, error) {
	accountModel := model.AccountDTO{
		Name:         req.Name,
		Description:  req.Description,
		UserId:       req.UserID,
		CompanyID:    req.CompanyID,
		CredentialID: req.CredentialID,
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

// UpdateDescription func update account's description & returns model.Account ID
func (a AccountService) UpdateDescription(ctx context.Context, req request.UpdateCompanyDescriptionRequest) (string, error) {
	id, err := a.accountRepo.UpdateDescription(ctx, req.ID, req.NewDescription)
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

// FindAllByCompanyID func used to find all Account by Company ID
func (a AccountService) FindAllByCompanyID(ctx context.Context, req request.FindAllAccountsByCompanyIDRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAllByCompanyID(ctx, req.CompanyID)
	if err != nil {
		return nil, errors.Wrap(err, "can't find accounts by company ID")
	}

	return accounts, nil
}

// FindAllByUserID func used to find all Account by UserID
func (a AccountService) FindAllByUserID(ctx context.Context, req request.FindAllAccountsByUserIDRequest) ([]model.AccountDTO, error) {
	accounts, err := a.accountRepo.FindAllByUserID(ctx, req.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "can't find accounts by user ID")
	}

	return accounts, nil
}

// FindByCredentialLogin func used to find all Account by Credential login
func (a AccountService) FindByCredentialLogin(ctx context.Context, req request.FindAccountsByCredentialLoginRequest) ([]model.AccountDTO, error) {
	var accounts []model.AccountDTO

	credentials, err := a.credentialRepo.FindByLogin(ctx, req.CredentialLogin)
	if err != nil {
		return nil, err
	}

	for _, credential := range credentials {
		foundedByLoginAccount, err := a.accountRepo.FindByCredentialID(ctx, credential.ID)
		if err != nil {
			return nil, errors.Wrap(err, "can't find by credentialLogin")
		}

		accounts = append(accounts, *foundedByLoginAccount)
	}

	return accounts, nil
}

// FindByCredentialEmail func used to find all Account by Credential email
func (a AccountService) FindByCredentialEmail(ctx context.Context, req request.FindAccountsByCredentialEmailRequest) ([]model.AccountDTO, error) {
	var accounts []model.AccountDTO

	credentials, err := a.credentialRepo.FindByEmail(ctx, req.CredentialEmail)
	if err != nil {
		return nil, err
	}

	for _, credential := range credentials {
		foundedByEmailAccount, err := a.accountRepo.FindByCredentialID(ctx, credential.ID)
		if err != nil {
			return nil, errors.Wrap(err, "can't find by credentialEmail")
		}

		accounts = append(accounts, *foundedByEmailAccount)
	}

	return accounts, nil
}

// FindByCredentialPhone func used to find all Account by Credential phone
func (a AccountService) FindByCredentialPhone(ctx context.Context, req request.FindAccountsByCredentialPhoneRequest) ([]model.AccountDTO, error) {
	var accounts []model.AccountDTO

	credentials, err := a.credentialRepo.FindByLogin(ctx, req.CredentialPhone)
	if err != nil {
		return nil, err
	}

	for _, credential := range credentials {
		foundedByPhoneAccount, err := a.accountRepo.FindByCredentialID(ctx, credential.ID)
		if err != nil {
			return nil, errors.Wrap(err, "can't find by credentialPhone")
		}

		accounts = append(accounts, *foundedByPhoneAccount)
	}

	return accounts, nil
}

// FindByCredentialName func used to find all Account by Credential name
func (a AccountService) FindByCredentialName(ctx context.Context, req request.FindAccountsByCredentialNameRequest) ([]model.AccountDTO, error) {
	var accounts []model.AccountDTO

	credentials, err := a.credentialRepo.FindByLogin(ctx, req.CredentialName)
	if err != nil {
		return nil, err
	}

	for _, credential := range credentials {
		foundedByNameAccount, err := a.accountRepo.FindByCredentialID(ctx, credential.ID)
		if err != nil {
			return nil, errors.Wrap(err, "can't find by credentialName")
		}

		accounts = append(accounts, *foundedByNameAccount)
	}

	return accounts, nil
}

// FindByCredentialMiddlename func used to find all Account by Credential middleName
func (a AccountService) FindByCredentialMiddlename(ctx context.Context, req request.FindAccountsByCredentialMiddlenameRequest) ([]model.AccountDTO, error) {
	var accounts []model.AccountDTO

	credentials, err := a.credentialRepo.FindByLogin(ctx, req.CredentialMiddlename)
	if err != nil {
		return nil, err
	}

	for _, credential := range credentials {
		foundedByMiddleNameAccount, err := a.accountRepo.FindByCredentialID(ctx, credential.ID)
		if err != nil {
			return nil, errors.Wrap(err, "can't find by credentialMiddleName")
		}

		accounts = append(accounts, *foundedByMiddleNameAccount)
	}

	return accounts, nil
}

// FindByCredentialSurname func used to find all Account by Credential surname
func (a AccountService) FindByCredentialSurname(ctx context.Context, req request.FindAccountsByCredentialSurnameRequest) ([]model.AccountDTO, error) {
	var accounts []model.AccountDTO

	credentials, err := a.credentialRepo.FindByLogin(ctx, req.CredentialSurname)
	if err != nil {
		return nil, err
	}

	for _, credential := range credentials {
		foundedBySurnameAccount, err := a.accountRepo.FindByCredentialID(ctx, credential.ID)
		if err != nil {
			return nil, errors.Wrap(err, "can't find by credentialSurname")
		}

		accounts = append(accounts, *foundedBySurnameAccount)
	}

	return accounts, nil
}

// UpdateCredentialLogin func used to update Credential login by ID & returned ID of updated Credential
func (a AccountService) UpdateCredentialLogin(ctx context.Context, req request.UpdateAccountCredentialLoginRequest) (string, error) {
	credentialID, err := a.accountRepo.FindCredentialIDByAccountID(ctx, req.ID)
	if err != nil {
		return "", err
	}

	updatedCredentialID, err := a.credentialRepo.UpdateLogin(ctx, credentialID, req.NewCredentialLogin)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential login")
	}

	return updatedCredentialID, nil
}

// UpdateCredentialEmail func used to update Credential email by ID & returned ID of updated Credential
func (a AccountService) UpdateCredentialEmail(ctx context.Context, req request.UpdateAccountCredentialEmailRequest) (string, error) {
	credentialID, err := a.accountRepo.FindCredentialIDByAccountID(ctx, req.ID)
	if err != nil {
		return "", err
	}

	updatedCredentialID, err := a.credentialRepo.UpdateEmail(ctx, credentialID, req.NewCredentialEmail)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential email")
	}

	return updatedCredentialID, nil
}

// UpdateCredentialPhone func used to update Credential phone by ID & returned ID of updated Credential
func (a AccountService) UpdateCredentialPhone(ctx context.Context, req request.UpdateAccountCredentialPhoneRequest) (string, error) {
	credentialID, err := a.accountRepo.FindCredentialIDByAccountID(ctx, req.ID)
	if err != nil {
		return "", err
	}

	updatedCredentialID, err := a.credentialRepo.UpdatePhone(ctx, credentialID, req.NewCredentialPhone)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential login")
	}

	return updatedCredentialID, nil
}

// UpdateCredentialName func used to update Credential name by ID & returned ID of updated Credential
func (a AccountService) UpdateCredentialName(ctx context.Context, req request.UpdateAccountCredentialNameRequest) (string, error) {
	credentialID, err := a.accountRepo.FindCredentialIDByAccountID(ctx, req.ID)
	if err != nil {
		return "", err
	}

	updatedCredentialID, err := a.credentialRepo.UpdateName(ctx, credentialID, req.NewCredentialName)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential name")
	}

	return updatedCredentialID, nil
}

// UpdateCredentialMiddlename func used to update Credential middleName by ID & returned ID of updated Credential
func (a AccountService) UpdateCredentialMiddlename(ctx context.Context, req request.UpdateAccountCredentialMiddlenameRequest) (string, error) {
	credentialID, err := a.accountRepo.FindCredentialIDByAccountID(ctx, req.ID)
	if err != nil {
		return "", err
	}

	updatedCredentialID, err := a.credentialRepo.UpdateMiddleName(ctx, credentialID, req.NewCredentialMiddlename)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential middle name")
	}

	return updatedCredentialID, nil
}

// UpdateCredentialSurname func used to update Credential surname by ID & returned ID of updated Credential
func (a AccountService) UpdateCredentialSurname(ctx context.Context, req request.UpdateAccountCredentialSurnameRequest) (string, error) {
	credentialID, err := a.accountRepo.FindCredentialIDByAccountID(ctx, req.ID)
	if err != nil {
		return "", err
	}

	updatedCredentialID, err := a.credentialRepo.UpdateSurname(ctx, credentialID, req.NewCredentialSurname)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential surname")
	}

	return updatedCredentialID, nil
}

