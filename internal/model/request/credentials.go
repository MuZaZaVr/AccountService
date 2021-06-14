package request

type (
	CreateCredentialRequest struct {
		Login      string
		Password   string
		Email      string
		Phone      string
		Name       string
		Middlename string
		Surname    string
		Age        int
		City       string
		Address    string
	}

	DeleteCredentialRequest struct {
		ID string
	}

	FindCredentialByLoginRequest struct {
		Login string
	}

	FindCredentialByEmailRequest struct {
		Email string
	}

	FindCredentialByPhoneRequest struct {
		Phone string
	}

	FindCredentialByNameRequest struct {
		Name string
	}

	FindCredentialByMiddlenameRequest struct {
		Middlename string
	}

	FindCredentialBySurnameRequest struct {
		Surname string
	}

	UpdateCredentialLoginRequest struct {
		ID       string
		NewLogin string
	}

	UpdateCredentialEmailRequest struct {
		ID       string
		NewEmail string
	}

	UpdateCredentialPhoneRequest struct {
		ID       string
		NewPhone string
	}

	UpdateCredentialNameRequest struct {
		ID      string
		NewName string
	}

	UpdateCredentialMiddlenameRequest struct {
		ID            string
		NewMiddlename string
	}

	UpdateCredentialSurnameRequest struct {
		ID         string
		NewSurname string
	}
)
