package service

import (
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/internal/model/request"
	"github.com/MuZaZaVr/account-service/internal/repository"
)

type CredentialService struct {
	credentialRepo repository.Credential
}

func (c CredentialService) Create(req request.CreateCredentialRequest) (int, error) {
	panic("implement me")
}

func (c CredentialService) FindByLogin(req request.FindCredentialByLoginRequest) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialService) FindByEmail(req request.FindCredentialByEmailRequest) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialService) FindByPhone(req request.FindCredentialByPhoneRequest) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialService) FindByName(req request.FindCredentialByNameRequest) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialService) FindByMiddleName(req request.FindCredentialByMiddlenameRequest) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialService) FindBySurname(req request.FindCredentialBySurnameRequest) ([]model.Credential, error) {
	panic("implement me")
}

func (c CredentialService) UpdateLogin(req request.UpdateCredentialLoginRequest) (string, error) {
	panic("implement me")
}

func (c CredentialService) UpdateEmail(req request.UpdateCredentialEmailRequest) (string, error) {
	panic("implement me")
}

func (c CredentialService) UpdatePhone(req request.UpdateCredentialPhoneRequest) (string, error) {
	panic("implement me")
}

func (c CredentialService) UpdateName(req request.UpdateCredentialNameRequest) (string, error) {
	panic("implement me")
}

func (c CredentialService) UpdateMiddleName(req request.UpdateCredentialMiddlenameRequest) (string, error) {
	panic("implement me")
}

func (c CredentialService) UpdateSurname(req request.UpdateCredentialSurnameRequest) (string, error) {
	panic("implement me")
}

func (c CredentialService) Delete(req request.DeleteCredentialRequest) (bool, error) {
	panic("implement me")
}

func NewCredentialService(credentialRepo repository.Credential) *CredentialService {
	return &CredentialService{
		credentialRepo: credentialRepo,
	}
}
