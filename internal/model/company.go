package model

// CompanyDTO represents DTO structure for mongo.Company
type CompanyDTO struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
