package model

type Credential struct {
	Id string `json:"id"`

	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`

	Name       string `json:"name"`
	Middlename string `json:"middlename"`
	Surname    string `json:"surname"`
	Age        int    `json:"age"`

	City    string `json:"city"`
	Address string `json:"address"`
}
