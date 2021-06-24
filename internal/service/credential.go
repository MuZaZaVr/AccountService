package service

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/internal/model/request"
	"github.com/MuZaZaVr/account-service/internal/repository"
	"github.com/pkg/errors"
)

// CredentialService is a company service
type CredentialService struct {
	credentialRepo repository.Credential
}

// NewCredentialService is a constructor for company service
func NewCredentialService(credentialRepo repository.Credential) *CredentialService {
	return &CredentialService{
		credentialRepo: credentialRepo,
	}
}

// Create func creates new credential & returns id
func (c CredentialService) Create(ctx context.Context, req request.CreateCredentialRequest) (string, error) {
	credentialModel := model.CredentialDTO{
		Login:        req.Login,
		PasswordHash: req.Password,
		Email:        req.Email,
		Phone:        req.Phone,
		Name:         req.Name,
		Middlename:   req.Middlename,
		Surname:      req.Surname,
		Age:          req.Age,
		City:         req.City,
		Address:      req.Address,
	}

	id, err := c.credentialRepo.Create(ctx, credentialModel)
	if err != nil {
		return "", errors.Wrap(err, "cant create credential")
	}

	return id, nil
}

// FindByLogin func find credentials by provided login & returns slice of credentials
func (c CredentialService) FindByLogin(ctx context.Context, req request.FindCredentialByLoginRequest) ([]model.CredentialDTO, error) {
	credentials, err := c.credentialRepo.FindByLogin(ctx, req.Login)
	if err != nil {
		return nil, errors.Wrap(err, "can't find credentials by name")
	}

	return credentials, nil
}

// FindByEmail func find credentials by provided email & returns slice of credentials
func (c CredentialService) FindByEmail(ctx context.Context, req request.FindCredentialByEmailRequest) ([]model.CredentialDTO, error) {
	credentials, err := c.credentialRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.Wrap(err, "can't find credentials by email")
	}

	return credentials, nil
}

// FindByPhone func find credentials by provided phone & returns slice of credentials
func (c CredentialService) FindByPhone(ctx context.Context, req request.FindCredentialByPhoneRequest) ([]model.CredentialDTO, error) {
	credentials, err := c.credentialRepo.FindByPhone(ctx, req.Phone)
	if err != nil {
		return nil, errors.Wrap(err, "can't find credentials by phone")
	}

	return credentials, nil
}

// FindByName func find credentials by provided name & returns slice of credentials
func (c CredentialService) FindByName(ctx context.Context, req request.FindCredentialByNameRequest) ([]model.CredentialDTO, error) {
	credentials, err := c.credentialRepo.FindByName(ctx, req.Name)
	if err != nil {
		return nil, errors.Wrap(err, "can't find credentials by name")
	}

	return credentials, nil
}

// FindByMiddleName func find credentials by provided middleName & returns slice of credentials
func (c CredentialService) FindByMiddleName(ctx context.Context, req request.FindCredentialByMiddlenameRequest) ([]model.CredentialDTO, error) {
	credentials, err := c.credentialRepo.FindByMiddleName(ctx, req.Middlename)
	if err != nil {
		return nil, errors.Wrap(err, "can't find credentials by meddle name")
	}

	return credentials, nil
}

// FindBySurname func find credentials by provided surname & returns slice of credentials
func (c CredentialService) FindBySurname(ctx context.Context, req request.FindCredentialBySurnameRequest) ([]model.CredentialDTO, error) {
	credentials, err := c.credentialRepo.FindBySurname(ctx, req.Surname)
	if err != nil {
		return nil, errors.Wrap(err, "can't find credentials by surname")
	}

	return credentials, nil
}

// UpdateLogin func update Credential login & returns Credential ID
func (c CredentialService) UpdateLogin(ctx context.Context, req request.UpdateCredentialLoginRequest) (string, error) {
	id, err := c.credentialRepo.UpdateLogin(ctx, req.ID, req.NewLogin)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential's name")
	}

	return id, nil
}

// UpdateEmail func update Credential email & returns Credential ID
func (c CredentialService) UpdateEmail(ctx context.Context, req request.UpdateCredentialEmailRequest) (string, error) {
	id, err := c.credentialRepo.UpdateEmail(ctx, req.ID, req.NewEmail)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential's email")
	}

	return id, nil
}

// UpdatePhone func update Credential phone & returns Credential ID
func (c CredentialService) UpdatePhone(ctx context.Context, req request.UpdateCredentialPhoneRequest) (string, error) {
	id, err := c.credentialRepo.UpdatePhone(ctx, req.ID, req.NewPhone)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential's phone")
	}

	return id, nil
}

// UpdateName func update Credential name & returns Credential ID
func (c CredentialService) UpdateName(ctx context.Context, req request.UpdateCredentialNameRequest) (string, error) {
	id, err := c.credentialRepo.UpdateName(ctx, req.ID, req.NewName)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential's name")
	}

	return id, nil
}

// UpdateMiddleName func update Credential middleName & returns Credential ID
func (c CredentialService) UpdateMiddleName(ctx context.Context, req request.UpdateCredentialMiddlenameRequest) (string, error) {
	id, err := c.credentialRepo.UpdateMiddleName(ctx, req.ID, req.NewMiddlename)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential's middle name")
	}

	return id, nil
}

// UpdateSurname func update Credential surname & returns Credential ID
func (c CredentialService) UpdateSurname(ctx context.Context, req request.UpdateCredentialSurnameRequest) (string, error) {
	id, err := c.credentialRepo.UpdateSurname(ctx, req.ID, req.NewSurname)
	if err != nil {
		return "", errors.Wrap(err, "can't update credential's surname")
	}

	return id, nil
}

// Delete func delete Credential & returns Credential ID
func (c CredentialService) Delete(ctx context.Context, req request.DeleteCredentialRequest) (string, error) {
	id, err := c.credentialRepo.Delete(ctx, req.ID)
	if err != nil {
		return "", errors.Wrap(err, "can't delete credentials")
	}

	return id, nil
}
