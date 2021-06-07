package service

import (
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/internal/model/request"
	"github.com/MuZaZaVr/account-service/internal/repository"
)

type CompanyService struct {
	companyRepo repository.Company
}

func (c CompanyService) Create(req request.CreateCompanyRequest) (string, error) {
	panic("implement me")
}

func (c CompanyService) FindByName(req request.FindCompanyByNameRequest) (model.Company, error) {
	panic("implement me")
}

func (c CompanyService) FindByURL(req request.FindCompanyByURLRequest) (model.Company, error) {
	panic("implement me")
}

func (c CompanyService) UpdateName(req request.UpdateCompanyNameRequest) (string, error) {
	panic("implement me")
}

func (c CompanyService) UpdateDescription(req request.UpdateCompanyDescriptionRequest) (string, error) {
	panic("implement me")
}

func (c CompanyService) UpdateURL(req request.UpdateCompanyURLRequest) (string, error) {
	panic("implement me")
}

func (c CompanyService) Delete(req request.DeleteCompanyRequest) (bool, error) {
	panic("implement me")
}

func (c CompanyService) IsExist(req request.IsCompanyExistRequest) (bool, error) {
	panic("implement me")
}

func NewCompanyService(companyRepo repository.Company) *CompanyService {
	return &CompanyService{
		companyRepo: companyRepo,
	}
}
