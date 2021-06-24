package request

import "github.com/MuZaZaVr/account-service/internal/model"

type (
	CreateAccountRequest struct {
		Name        string
		Description string

		UserID int

		Company     model.CompanyDTO
		Credentials model.CredentialDTO
	}

	DeleteAccountRequest struct {
		ID string
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

	FindAllAccountsByCompanyIDRequest struct {
		CompanyID string
	}

	FindAllAccountsByUserIDRequest struct {
		UserID int
	}

	UpdateAccountCredentialLoginRequest struct {
		ID                 string
		NewCredentialLogin string
	}

	UpdateAccountCredentialEmailRequest struct {
		ID                 string
		NewCredentialEmail string
	}

	UpdateAccountCredentialPhoneRequest struct {
		ID                 string
		NewCredentialPhone string
	}

	UpdateAccountCredentialNameRequest struct {
		ID                string
		NewCredentialName string
	}

	UpdateAccountCredentialMiddlenameRequest struct {
		ID                      string
		NewCredentialMiddlename string
	}

	UpdateAccountCredentialSurnameRequest struct {
		ID                   string
		NewCredentialSurname string
	}
)
