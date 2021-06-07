package request

import (
	"github.com/MuZaZaVr/account-service/internal/model"
)

type (
	CreateAccountRequest struct {
		Name string
		Description string

		UserId int

		Company model.Company
		Credentials model.Credential
	}

	DeleteAccountRequest struct {
		Id string
	}

	FindAccountByNameRequest struct {
		Name string
	}

	FindAccountsByCredentialLoginRequest struct {
		CredentialLogin string
	}

	FindAccountsByCredentialEmailRequest struct {
		CredentialEmail string
	}

	FindAccountsByCredentialPhoneRequest struct {
		CredentialPhone string
	}

	FindAccountsByCredentialNameRequest struct {
		CredentialName string
	}

	FindAccountsByCredentialMiddlenameRequest struct {
		CredentialMiddlename string
	}

	FindAccountsByCredentialSurnameRequest struct {
		CredentialSurname string
	}

	UpdateAccountCredentialLoginRequest struct {
		NewCredentialLogin string
	}

	UpdateAccountCredentialEmailRequest struct {
		NewCredentialEmail string
	}

	UpdateAccountCredentialPhoneRequest struct {
		NewCredentialPhone string
	}

	UpdateAccountCredentialNameRequest struct {
		NewCredentialName string
	}

	UpdateAccountCredentialMiddlenameRequest struct {
		NewCredentialMiddlename string
	}

	UpdateAccountCredentialSurnameRequest struct {
		NewCredentialSurname string
	}
)

