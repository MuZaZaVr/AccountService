package mongo

type Credentials struct {
	Login        string `bson:"login"`
	PasswordHash string `bson:"password_hash"`
	Email        string `bson:"email"`
	Phone        string `bson:"phone"`

	Name       string `bson:"name"`
	Middlename string `bson:"middlename"`
	Surname    string `bson:"surname"`
	Age        int    `bson:"age"`

	City    string `bson:"city"`
	Address string `bson:"address"`
}
