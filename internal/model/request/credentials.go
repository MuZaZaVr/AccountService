package request

type (
	CreateCredentialRequest struct {
		Login string
		Password string
		Email string
		Phone string
		Name string
		Middlename string
		Surname string
		Age int
		City string
		Address string
	}

	DeleteCredentialRequest struct {
		Id int
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
		NewLogin string
	}

	UpdateCredentialEmailRequest struct {
		NewEmail string
	}

	UpdateCredentialPhoneRequest struct {
		Phone string
	}

	UpdateCredentialNameRequest struct {
		Name string
	}

	UpdateCredentialMiddlenameRequest struct {
		Middlename string
	}

	UpdateCredentialSurnameRequest struct {
		Surname string
	}
)
