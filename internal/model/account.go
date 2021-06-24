package model

// AccountDTO represents DTO structure for mongo.Account
type AccountDTO struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`

	UserId      int            `json:"user_id"`
	Company     CompanyDTO     `json:"company"`
	Credentials CredentialsDTO `json:"credentials"`
}
