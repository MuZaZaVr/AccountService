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

func ConnectToCompanyMongo() (context.Context, *CompanyRepository, error) {
	ctx := context.Background()

	cfg, err := config.Init(configPath)
	if err != nil {
		return nil, nil, err
	}

	db, err := mongo.NewMongo(ctx, cfg.Mongo)

	repo := NewCompanyRepository(db)

	return ctx, repo, nil
}

func TestCompanyRepository_Create(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCompanyMongo()
	require.NoError(t, err)

	type test struct {
		name    string
		company model.CompanyDTO
	}

	testTable := []test{
		{
			name: "all OK",
			company: model.CompanyDTO{
				Name:        "TestCompanyName",
				Description: "TestCompanyDescription",
				URL:         "TestCompanyURL",
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var companyID string

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			companyID, err = repo.Create(ctx, testCase.company)
			assert.NoError(err)

			assert.NotEmpty(companyID)

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCompanyRepository_FindByName(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCompanyMongo()
	require.NoError(t, err)

	type test struct {
		name string

		isOK bool

		companyName string

		company         model.CompanyDTO
		expectedCompany *model.CompanyDTO

		fc func(data *test) ()

		expectedError error
	}

	testTable := []test{
		{
			name:        "all OK",
			isOK:        true,
			companyName: "TestCompanyName",
			company: model.CompanyDTO{
				Name:        "TestCompanyName",
				Description: "TestCompanyDescription",
				URL:         "TestCompanyURL",
			},
			expectedCompany: &model.CompanyDTO{},

			fc: func(data *test) {
				data.expectedCompany = &data.company
			},
		},
		{
			name:          "not found",
			companyName:   "some invalid company name",
			expectedError: errors.New("mongo: no documents in result"),
		},
		{
			name:          "empty name",
			expectedError: errors.New("empty company name"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.fc != nil {
				testCase.fc(&testCase)
			}

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				testCase.expectedCompany.ID, err = repo.Create(ctx, testCase.company)
				assert.NoError(err)
			}

			foundedCompany, err := repo.FindByName(ctx, testCase.companyName)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				assert.Equal(testCase.expectedCompany, foundedCompany)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCompanyRepository_FindByURL(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCompanyMongo()
	require.NoError(t, err)

	type test struct {
		name string
		URL  string

		company model.CompanyDTO

		isOK            bool
		expectedError   error
		expectedCompany *model.CompanyDTO

		fc func(data *test) ()
	}

	testTable := []test{
		{
			name: "all OK",
			isOK: true,
			URL:  "TestCompanyURL",
			company: model.CompanyDTO{
				Name:        "TestCompanyName",
				Description: "TestCompanyDescription",
				URL:         "TestCompanyURL",
			},
			expectedCompany: &model.CompanyDTO{},

			fc: func(data *test) {
				data.expectedCompany = &data.company
			},
		},
		{
			name:          "not found",
			URL:           "some invalid company URL",
			expectedError: errors.New("mongo: no documents in result"),
		},
		{
			name:          "empty URL",
			expectedError: errors.New("empty company URL"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.fc != nil {
				testCase.fc(&testCase)
			}

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOK {
				testCase.expectedCompany.ID, err = repo.Create(ctx, testCase.company)
				assert.NoError(err)
			}

			foundedCompany, err := repo.FindByURL(ctx, testCase.URL)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				assert.Equal(testCase.expectedCompany, foundedCompany)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCompanyRepository_UpdateName(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCompanyMongo()
	require.NoError(t, err)

	type test struct {
		name           string
		companyID      string
		companyNewName string

		company model.CompanyDTO

		isOk           bool
		expectedError  error
		expectedResult string

		fc func(data *test)
	}

	testTable := []test{
		{
			name:           "all OK",
			companyNewName: "new company name",

			company: model.CompanyDTO{
				Name:        "company name",
				Description: "company description",
				URL:         "company URL",
			},

			isOk: true,
		},
		{
			name:          "invalid ID",
			companyID:     "",
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
		{
			name:          "not found",
			companyID:     primitive.NewObjectID().Hex(),
			expectedError: errors.New("mongo: no documents in result"),

			fc: func(data *test) {
				data.expectedResult = data.companyID
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.fc != nil {
				testCase.fc(&testCase)
			}

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOk {
				testCase.companyID, err = repo.Create(ctx, testCase.company)
				testCase.expectedResult = testCase.companyID
				assert.NoError(err)
			}

			responseID, err := repo.UpdateName(ctx, testCase.companyID, testCase.companyNewName)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOk {
				assert.Equal(testCase.expectedResult, responseID)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCompanyRepository_UpdateDescription(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCompanyMongo()
	require.NoError(t, err)

	type test struct {
		name                  string
		companyID             string
		companyNewDescription string

		company model.CompanyDTO

		isOk           bool
		expectedError  error
		expectedResult string

		fc func(data *test)
	}

	testTable := []test{
		{
			name:                  "all OK",
			companyNewDescription: "new company description",

			company: model.CompanyDTO{
				Name:        "company name",
				Description: "company description",
				URL:         "company URL",
			},

			isOk: true,
		},
		{
			name:          "invalid ID",
			companyID:     "",
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
		{
			name:          "not found",
			companyID:     primitive.NewObjectID().Hex(),
			expectedError: errors.New("mongo: no documents in result"),

			fc: func(data *test) {
				data.expectedResult = data.companyID
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.fc != nil {
				testCase.fc(&testCase)
			}

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOk {
				testCase.companyID, err = repo.Create(ctx, testCase.company)
				testCase.expectedResult = testCase.companyID
				assert.NoError(err)
			}

			responseID, err := repo.UpdateDescription(ctx, testCase.companyID, testCase.companyNewDescription)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOk {
				assert.Equal(testCase.expectedResult, responseID)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCompanyRepository_UpdateURL(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCompanyMongo()
	require.NoError(t, err)

	type test struct {
		name          string
		companyID     string
		companyNewURL string

		company model.CompanyDTO

		isOk           bool
		expectedError  error
		expectedResult string

		fc func(data *test)
	}

	testTable := []test{
		{
			name:          "all OK",
			companyNewURL: "new company description",

			company: model.CompanyDTO{
				Name:        "company name",
				Description: "company description",
				URL:         "company URL",
			},

			isOk: true,
		},
		{
			name:          "invalid ID",
			companyID:     "",
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
		{
			name:          "not found",
			companyID:     primitive.NewObjectID().Hex(),
			expectedError: errors.New("mongo: no documents in result"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOk {
				testCase.companyID, err = repo.Create(ctx, testCase.company)
				testCase.expectedResult = testCase.companyID
				assert.NoError(err)
			}

			responseID, err := repo.UpdateURL(ctx, testCase.companyID, testCase.companyNewURL)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOk {
				assert.Equal(testCase.expectedResult, responseID)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCompanyRepository_Delete(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCompanyMongo()
	require.NoError(t, err)

	type test struct {
		name      string
		companyID string

		company model.CompanyDTO

		isOk           bool
		expectedError  error
		expectedResult string

		fc func(data *test)
	}

	testTable := []test{
		{
			name: "all OK",

			company: model.CompanyDTO{
				Name:        "company name",
				Description: "company description",
				URL:         "company URL",
			},

			isOk: true,
		},
		{
			name:          "invalid ID",
			companyID:     "",
			expectedError: errors.New("the provided hex string is not a valid ObjectID"),
		},
		{
			name:          "not found",
			companyID:     primitive.NewObjectID().Hex(),
			expectedError: errors.New("mongo: no documents in result"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOk {
				testCase.companyID, err = repo.Create(ctx, testCase.company)
				testCase.expectedResult = testCase.companyID
				assert.NoError(err)
			}

			responseID, err := repo.Delete(ctx, testCase.companyID)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOk {
				assert.Equal(testCase.expectedResult, responseID)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestCompanyRepository_IsExist(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToCompanyMongo()
	require.NoError(t, err)

	type test struct {
		name        string
		companyName string

		company model.CompanyDTO

		isOk           bool
		expectedError  error
		expectedResult bool

		fc func(data *test)
	}

	testTable := []test{
		{
			name: "all OK",
			companyName: "company name",
			company: model.CompanyDTO{
				Name:        "company name",
				Description: "company description",
				URL:         "company URL",
			},

			isOk: true,
		},
		{
			name:          "invalid company name",
			companyName:     "",
			expectedError: errors.New("empty company name"),
		},
		{
			name:          "not found",
			companyName:     primitive.NewObjectID().Hex(),
			expectedError: errors.New("mongo: no documents in result"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			if testCase.isOk {
				testCase.company.ID, err = repo.Create(ctx, testCase.company)
				testCase.expectedResult = true
				assert.NoError(err)
			}

			responseID, err := repo.IsExist(ctx, testCase.companyName)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOk {
				assert.Equal(testCase.expectedResult, responseID)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}
