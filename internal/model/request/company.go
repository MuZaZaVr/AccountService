package request

type (
	CreateCompanyRequest struct {
		Name        string
		Description string
		Url         string
	}

	DeleteCompanyRequest struct {
		Id int
	}


	FindCompanyByNameRequest struct {
		Name string
	}

	FindCompanyByURLRequest struct {
		URL string
	}

	UpdateCompanyNameRequest struct {
		Name string
	}

	UpdateCompanyDescriptionRequest struct {
		Description string
	}

	UpdateCompanyURLRequest struct {
		URL string
	}


	IsCompanyExistRequest struct {
		Name string
	}
)
