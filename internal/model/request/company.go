package request

type (
	CreateCompanyRequest struct {
		Name        string
		Description string
		URL         string
	}

	DeleteCompanyRequest struct {
		ID string
	}

	FindCompanyByNameRequest struct {
		Name string
	}

	FindCompanyByURLRequest struct {
		URL string
	}

	UpdateCompanyNameRequest struct {
		ID      string
		NewName string
	}

	UpdateCompanyDescriptionRequest struct {
		ID             string
		NewDescription string
	}

	UpdateCompanyURLRequest struct {
		ID     string
		NewURL string
	}

	IsCompanyExistRequest struct {
		Name string
	}
)
