package model

type Account struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	UserId       int    `json:"user_id"`
	CompanyId    string `json:"company_id"`
	CredentialId string `json:"credential_id"`
}
