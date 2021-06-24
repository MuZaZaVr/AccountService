package service

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/internal/model/request"
	"github.com/MuZaZaVr/account-service/internal/repository"
	"github.com/pkg/errors"
)

// CompanyService is a company service
type CompanyService struct {
	companyRepo repository.Company
}

// NewCompanyService is a constructor for company service
func NewCompanyService(companyRepo repository.Company) *CompanyService {
	return &CompanyService{
		companyRepo: companyRepo,
	}
}

// Create func creates new company & returns id
func (c CompanyService) Create(ctx context.Context, req request.CreateCompanyRequest) (string, error) {
	companyModel := model.CompanyDTO{
		Name:        req.Name,
		Description: req.Description,
		URL:         req.URL,
	}

	id, err := c.companyRepo.Create(ctx, companyModel)
	if err != nil {
		return "", errors.Wrap(err, "cant create company")
	}

	return id, nil
}

// FindByName func find company by provided name & returns company
func (c CompanyService) FindByName(ctx context.Context, req request.FindCompanyByNameRequest) (*model.CompanyDTO, error) {
	company, err := c.companyRepo.FindByName(ctx, req.Name)
	if err != nil {
		return nil, errors.Wrap(err, "can't find company by name")
	}

	return company, nil
}

// FindByURL func find Company by provided URL & returns Company
func (c CompanyService) FindByURL(ctx context.Context, req request.FindCompanyByURLRequest) (*model.CompanyDTO, error) {
	company, err := c.companyRepo.FindByURL(ctx, req.URL)
	if err != nil {
		return nil, errors.Wrap(err, "can't find company by URL")
	}

	return company, nil
}

// UpdateName func update Company name & returns Company ID
func (c CompanyService) UpdateName(ctx context.Context, req request.UpdateCompanyNameRequest) (string, error) {
	id, err := c.companyRepo.UpdateName(ctx, req.ID, req.NewName)
	if err != nil {
		return "", errors.Wrap(err, "can't update company's name")
	}

	return id, nil
}

// UpdateDescription func update Company description & returns Company ID
func (c CompanyService) UpdateDescription(ctx context.Context, req request.UpdateCompanyDescriptionRequest) (string, error) {
	id, err := c.companyRepo.UpdateDescription(ctx, req.ID, req.NewDescription)
	if err != nil {
		return "", errors.Wrap(err, "can't update company's description")
	}

	return id, nil
}

// UpdateURL func update Company URL & returns Company ID
func (c CompanyService) UpdateURL(ctx context.Context, req request.UpdateCompanyURLRequest) (string, error) {
	id, err := c.companyRepo.UpdateURL(ctx, req.ID, req.NewURL)
	if err != nil {
		return "", errors.Wrap(err, "can't update company's URL")
	}

	return id, nil
}

// Delete func delete Company & returns Company ID
func (c CompanyService) Delete(ctx context.Context, req request.DeleteCompanyRequest) (string, error) {
	id, err := c.companyRepo.Delete(ctx, req.ID)
	if err != nil {
		return "", errors.Wrap(err, "can't delete company")
	}

	return id, nil
}

// IsExist check company existence & returns true if company exists or false of not
func (c CompanyService) IsExist(ctx context.Context, req request.IsCompanyExistRequest) (bool, error) {
	isExist, err := c.companyRepo.IsExist(ctx, req.Name)
	if err != nil {
		return false, errors.Wrap(err, "can't confirm company's existence")
	}

	return isExist, nil
}
