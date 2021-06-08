package model

import (
	"github.com/MuZaZaVr/account-service/pkg/database/mongo"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Company represents Company model for mongo
type Company mongo.Company

// CompanyDTO represents DTO structure for mongo.Company
type CompanyDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// ConvertFromDTOToMongoModel func convert CompanyDTO model into mongo.Company model
func (c CompanyDTO) ConvertFromDTOToMongoModel() (*Company, error) {
	company := Company{
		Name:        c.Name,
		Description: c.Description,
		URL:         c.URL,
	}

	var err error
	if c.ID != "" {
		company.ID, err = primitive.ObjectIDFromHex(c.ID)
		if err != nil {
			return nil, errors.Wrap(err, "invalid ID")
		}
	}

	return &company, err
}

// ConvertFromMongoModelToDTO func convert Company model into CompanyDTO model
func (c Company) ConvertFromMongoModelToDTO() *CompanyDTO {
	companyDTO := CompanyDTO{
		ID:          c.ID.Hex(),
		Name:        c.Name,
		Description: c.Description,
		URL:         c.URL,
	}

	return &companyDTO
}
