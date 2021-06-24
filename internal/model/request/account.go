package request

import "github.com/MuZaZaVr/account-service/internal/model"

type (
	CreateAccountRequest struct {
		Name        string
		Description string

		UserID int

		Company    model.CompanyDTO
		Credential model.CredentialsDTO
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

	FindAccountsByCredentialAgeRequest struct {
		CredentialAge int
	}

	FindAccountsByCredentialCityRequest struct {
		CredentialCity string
	}

	FindAccountsByCredentialAddressRequest struct {
		CredentialAddress string
	}

	FindAllAccountsByCompanyIDRequest struct {
		CompanyID string
	}

	FindAllAccountsByUserIDRequest struct {
		UserID int
	}

	UpdateAccountRequest struct {
		ID             string
		UpdatedAccount model.AccountDTO
	}
)
