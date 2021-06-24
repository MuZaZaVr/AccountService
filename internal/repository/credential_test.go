package repository

import (
	"context"
	"errors"
	"github.com/MuZaZaVr/account-service/internal/config"
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/pkg/database/mongo"
	assertTest "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

const configPath = "config/main"

var credentialCreateModel = model.CredentialDTO{
	Login:        "TestLogin",
	PasswordHash: "d1e8a70b5ccab1dc2f56bbf7e99f064a660c08e361a35751b9c483c88943d082",
	Email:        "login-example@gmail.com",
	Phone:        "+375 (29) 299-99-99",
	Name:         "TestName",
	Middlename:   "TestMiddleName",
	Surname:      "TestSurname",
	Age:          20,
	City:         "TestCity",
	Address:      "TestAddress",
}

func ConnectToCredentialMongo() (context.Context, *CredentialRepository, error) {
	ctx := context.Background()

	cfg, err := config.Init(configPath)
	if err != nil {
		return nil, nil, err
	}

	db, err := mongo.NewMongo(ctx, cfg.Mongo)
	if err != nil {
		return nil, nil, err
	}

	return ctx, NewCredentialRepository(db), nil
}

func TestCredentialRepository_Create(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name       string
		credential model.CredentialDTO
	}

	testTable := []test{
		{
			name:       "all OK",
			credential: credentialCreateModel,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var credentialID string

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			credentialID, err = repo.Create(ctx, testCase.credential)
			assert.NoError(err)

			assert.NotEmpty(credentialID)

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}

}

func TestCredentialRepository_UpdateLogin(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name                string
		isOK                bool
		requestCredentialID string
		credential          model.CredentialDTO
		expectedError       error
	}

	testTable := []test{
		{
			name:       "all OK",
			isOK:       true,
			credential: credentialCreateModel,
		},
		{
			name:          "not correct CredentialID",
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
		{
			name:                "not found",
			requestCredentialID: primitive.NewObjectID().Hex(),
			expectedError:       errors.New("mongo: no documents in result"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var credentialID = testCase.requestCredentialID

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				credentialID, err = repo.Create(ctx, testCase.credential)
			}
			updatedID, err := repo.UpdateLogin(ctx, credentialID, "NewTestLogin")
			assert.Equal(err, testCase.expectedError)

			if testCase.isOK {
				assert.Equal(credentialID, updatedID)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_UpdateEmail(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name                string
		isOk                bool
		requestCredentialID string
		credential          model.CredentialDTO
		expectedError       error
	}

	testTable := []test{
		{
			name:       "all OK",
			isOk:       true,
			credential: credentialCreateModel,
		},
		{
			name:                "not found",
			requestCredentialID: primitive.NewObjectID().Hex(),
			expectedError:       errors.New("mongo: no documents in result"),
		},
		{
			name:          "not correct CredentialID",
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var credentialID = testCase.requestCredentialID

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOk {
				credentialID, err = repo.Create(ctx, testCase.credential)
				assert.NoError(err)
			}

			updatedCredentialID, err := repo.UpdateEmail(ctx, credentialID, "NewLogin")
			assert.Equal(testCase.expectedError, err)

			if testCase.isOk {
				require.Equal(t, credentialID, updatedCredentialID)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_UpdatePhone(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name                string
		isOK                bool
		credentialRequestID string
		credential          model.CredentialDTO
		expectedError       error
	}

	testTable := []test{
		{
			name:       "all OK",
			isOK:       true,
			credential: credentialCreateModel,
		},
		{
			name:                "not found",
			credentialRequestID: primitive.NewObjectID().Hex(),
			expectedError:       errors.New("mongo: no documents in result"),
		},
		{
			name:          "not correct CredentialID",
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var credentialID = testCase.credentialRequestID

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				credentialID, err = repo.Create(ctx, testCase.credential)
				assert.NoError(err)
			}

			responseCredentialID, err := repo.UpdatePhone(ctx, credentialID, "+375 (33) 333-33-33")
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				assert.Equal(credentialID, responseCredentialID)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_UpdateName(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name                string
		isOK                bool
		credentialRequestID string
		credential          model.CredentialDTO
		expectedError       error
	}

	testTable := []test{
		{
			name:       "all OK",
			isOK:       true,
			credential: credentialCreateModel,
		},
		{
			name:                "not found",
			credentialRequestID: primitive.NewObjectID().Hex(),
			expectedError:       errors.New("mongo: no documents in result"),
		},
		{
			name:          "not correct CredentialID",
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var credentialID = testCase.credentialRequestID

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				credentialID, err = repo.Create(ctx, testCase.credential)
				assert.NoError(err)
			}

			responseCredentialID, err := repo.UpdateName(ctx, credentialID, "NewName")
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				assert.Equal(credentialID, responseCredentialID)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_UpdateMiddleName(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name                string
		isOK                bool
		credentialRequestID string
		credential          model.CredentialDTO
		expectedError       error
	}

	testTable := []test{
		{
			name:       "all OK",
			isOK:       true,
			credential: credentialCreateModel,
		},
		{
			name:                "not found",
			credentialRequestID: primitive.NewObjectID().Hex(),
			expectedError:       errors.New("mongo: no documents in result"),
		},
		{
			name:          "not correct CredentialID",
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var credentialID = testCase.credentialRequestID

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				credentialID, err = repo.Create(ctx, testCase.credential)
				assert.NoError(err)
			}

			responseCredentialID, err := repo.UpdateMiddleName(ctx, credentialID, "NewMiddlename")
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				assert.Equal(credentialID, responseCredentialID)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_UpdateSurname(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name                string
		isOK                bool
		credentialRequestID string
		credential          model.CredentialDTO
		expectedError       error
	}

	testTable := []test{
		{
			name:       "all OK",
			isOK:       true,
			credential: credentialCreateModel,
		},
		{
			name:                "not found",
			credentialRequestID: primitive.NewObjectID().Hex(),
			expectedError:       errors.New("mongo: no documents in result"),
		},
		{
			name:          "not correct CredentialID",
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var credentialID = testCase.credentialRequestID

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				credentialID, err = repo.Create(ctx, testCase.credential)
				assert.NoError(err)
			}

			responseCredentialID, err := repo.UpdateSurname(ctx, credentialID, "NewSurname")
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				assert.Equal(credentialID, responseCredentialID)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_FindByLogin(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name  string
		isOK  bool
		login string
		fn    func(data *test)

		credentials        []model.CredentialDTO
		credentialExpected []model.CredentialDTO

		expectedError error
	}

	testTable := []test{
		{
			name: "all OK",
			isOK: true,

			fn: func(data *test) {
				data.credentialExpected = data.credentials
			},

			login: "TestLogin",
			credentials: []model.CredentialDTO{
				credentialCreateModel,
			},
			credentialExpected: []model.CredentialDTO{},
		},
		{
			name:  "not found",
			login: "Invalid login",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.fn != nil {
				testCase.fn(&testCase)
			}

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				for _, credential := range testCase.credentials {
					_, err := repo.Create(ctx, credential)
					assert.NoError(err)
				}
			}

			foundedCredentials, err := repo.FindByLogin(ctx, testCase.login)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				for i := range testCase.credentialExpected {
					testCase.credentialExpected[i].ID = foundedCredentials[i].ID
				}
				assert.Equal(testCase.credentialExpected, foundedCredentials)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_FindByEmail(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name  string
		isOK  bool
		email string
		fn    func(data *test)

		credentials        []model.CredentialDTO
		credentialExpected []model.CredentialDTO

		expectedError error
	}

	testTable := []test{
		{
			name: "all OK",
			isOK: true,

			fn: func(data *test) {
				data.credentialExpected = data.credentials
			},

			email: "login-example@gmail.com",
			credentials: []model.CredentialDTO{
				credentialCreateModel,
			},
			credentialExpected: []model.CredentialDTO{},
		},
		{
			name:  "not found",
			email: "Invalid login",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.fn != nil {
				testCase.fn(&testCase)
			}

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				for _, credential := range testCase.credentials {
					_, err := repo.Create(ctx, credential)
					assert.NoError(err)
				}
			}

			foundedCredentials, err := repo.FindByEmail(ctx, testCase.email)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				for i := range testCase.credentialExpected {
					testCase.credentialExpected[i].ID = foundedCredentials[i].ID
				}
				assert.Equal(testCase.credentialExpected, foundedCredentials)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_FindByPhone(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name  string
		isOK  bool
		phone string
		fn    func(data *test)

		credentials        []model.CredentialDTO
		credentialExpected []model.CredentialDTO

		expectedError error
	}

	testTable := []test{
		{
			name: "all OK",
			isOK: true,

			fn: func(data *test) {
				data.credentialExpected = data.credentials
			},

			phone: "+375 (29) 299-99-99",
			credentials: []model.CredentialDTO{
				credentialCreateModel,
			},
			credentialExpected: []model.CredentialDTO{},
		},
		{
			name:  "not found",
			phone: "Invalid login",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.fn != nil {
				testCase.fn(&testCase)
			}

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				for _, credential := range testCase.credentials {
					_, err := repo.Create(ctx, credential)
					assert.NoError(err)
				}
			}

			foundedCredentials, err := repo.FindByPhone(ctx, testCase.phone)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				for i := range testCase.credentialExpected {
					testCase.credentialExpected[i].ID = foundedCredentials[i].ID
				}
				assert.Equal(testCase.credentialExpected, foundedCredentials)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_FindByName(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name     string
		isOK     bool
		userName string
		fn       func(data *test)

		credentials        []model.CredentialDTO
		credentialExpected []model.CredentialDTO

		expectedError error
	}

	testTable := []test{
		{
			name: "all OK",
			isOK: true,

			fn: func(data *test) {
				data.credentialExpected = data.credentials
			},

			userName: "TestName",
			credentials: []model.CredentialDTO{
				credentialCreateModel,
			},
			credentialExpected: []model.CredentialDTO{},
		},
		{
			name:     "not found",
			userName: "Invalid login",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.fn != nil {
				testCase.fn(&testCase)
			}

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				for _, credential := range testCase.credentials {
					_, err := repo.Create(ctx, credential)
					assert.NoError(err)
				}
			}

			foundedCredentials, err := repo.FindByName(ctx, testCase.userName)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				for i := range testCase.credentialExpected {
					testCase.credentialExpected[i].ID = foundedCredentials[i].ID
				}
				assert.Equal(testCase.credentialExpected, foundedCredentials)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_FindByMiddleName(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name       string
		isOK       bool
		middlename string
		fn         func(data *test)

		credentials        []model.CredentialDTO
		credentialExpected []model.CredentialDTO

		expectedError error
	}

	testTable := []test{
		{
			name: "all OK",
			isOK: true,

			fn: func(data *test) {
				data.credentialExpected = data.credentials
			},

			middlename: "TestMiddleName",
			credentials: []model.CredentialDTO{
				credentialCreateModel,
			},
			credentialExpected: []model.CredentialDTO{},
		},
		{
			name:       "not found",
			middlename: "Invalid login",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.fn != nil {
				testCase.fn(&testCase)
			}

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				for _, credential := range testCase.credentials {
					_, err := repo.Create(ctx, credential)
					assert.NoError(err)
				}
			}

			foundedCredentials, err := repo.FindByMiddleName(ctx, testCase.middlename)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				for i := range testCase.credentialExpected {
					testCase.credentialExpected[i].ID = foundedCredentials[i].ID
				}
				assert.Equal(testCase.credentialExpected, foundedCredentials)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_FindBySurname(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name    string
		isOK    bool
		surname string
		fn      func(data *test)

		credentials        []model.CredentialDTO
		credentialExpected []model.CredentialDTO

		expectedError error
	}

	testTable := []test{
		{
			name: "all OK",
			isOK: true,

			fn: func(data *test) {
				data.credentialExpected = data.credentials
			},

			surname: "TestSurname",
			credentials: []model.CredentialDTO{
				credentialCreateModel,
			},
			credentialExpected: []model.CredentialDTO{},
		},
		{
			name:    "not found",
			surname: "Invalid login",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.fn != nil {
				testCase.fn(&testCase)
			}

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				for _, credential := range testCase.credentials {
					_, err := repo.Create(ctx, credential)
					assert.NoError(err)
				}
			}

			foundedCredentials, err := repo.FindBySurname(ctx, testCase.surname)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				for i := range testCase.credentialExpected {
					testCase.credentialExpected[i].ID = foundedCredentials[i].ID
				}
				assert.Equal(testCase.credentialExpected, foundedCredentials)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCredentialRepository_Delete(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repos, err := ConnectToCredentialMongo()
	require.NoError(t, err)

	type test struct {
		name       string
		isOk       bool
		credential model.CredentialDTO

		requestID  string

		expectedError error
	}

	testTable := []test{
		{
			name:       "all OK",
			isOk:       true,
			credential: credentialCreateModel,
		},
		{
			name:          "not correct Credential id",
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
		{
			name:      "not found",
			requestID: primitive.NewObjectID().Hex(),
			expectedError: errors.New("mongo: no documents in result"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var credentialID = testCase.requestID

			_, err := repos.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOk {
				credentialID, err = repos.Create(ctx, testCase.credential)
				assert.NoError(err)
			}

			responseID, err := repos.Delete(ctx, credentialID)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOk {
				assert.Equal(credentialID, responseID)
			}

			_, err = repos.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}
