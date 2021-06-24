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
	"testing"
)

func ConnectToAccountMongo() (context.Context, *AccountRepository, error) {
	ctx := context.Background()

	cfg, err := config.Init(configPath)
	if err != nil {
		return nil, nil, err
	}

	db, err := mongo.NewMongo(ctx, cfg.Mongo)

	repo := NewAccountRepository(db)

	return ctx, repo, nil
}

func TestAccountRepository_Create(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToAccountMongo()
	require.NoError(t, err)

	type test struct {
		name    string
		account model.AccountDTO
	}

	testTable := []test{
		{
			name: "all OK",
			account: model.AccountDTO{
				Name:        "TestAccountName",
				Description: "TestAccountDescription",
				UserId:      1,
				Company: model.CompanyDTO{
					Name:        "TestCompanyName",
					Description: "TestCompanyDescription",
					URL:         "TestCompanyURL",
				},
				Credentials: model.CredentialsDTO{
					Login:        "TestCredentialLogin",
					PasswordHash: "TestCredentialPasswordHash",
					Email:        "TestCredentialEmail",
					Phone:        "TestCredentialPhone",
					Name:         "TestCredentialName",
					Middlename:   "TestCredentialMiddlename",
					Surname:      "TestCredentialSurname",
					Age:          20,
					City:         "TestCredentialCity",
					Address:      "TestCredentialAddress",
				},
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var accountID string

			_, err := repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)

			accountID, err = repo.Create(ctx, testCase.account)
			assert.NoError(err)

			assert.NotEmpty(accountID)

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}

func TestAccountRepository_FindByName(t *testing.T) {
	assert := assertTest.New(t)

	ctx, repo, err := ConnectToAccountMongo()
	require.NoError(t, err)

	type test struct {
		name string

		isOK bool

		accountName string

		account         model.AccountDTO
		expectedAccount *model.AccountDTO

		fc func(data *test) ()

		expectedError error
	}

	testTable := []test{
		{
			name:        "all OK",
			isOK:        true,
			accountName: "TestAccountName",
			account: model.AccountDTO{
				Name:        "TestAccountName",
				Description: "TestAccountDescription",
				UserId:      1,
				Company: model.CompanyDTO{
					Name:        "TestCompanyName",
					Description: "TestCompanyDescription",
					URL:         "TestCompanyURL",
				},
				Credentials: model.CredentialsDTO{
					Login:        "TestCredentialLogin",
					PasswordHash: "TestCredentialPasswordHash",
					Email:        "TestCredentialEmail",
					Phone:        "TestCredentialPhone",
					Name:         "TestCredentialName",
					Middlename:   "TestCredentialMiddlename",
					Surname:      "TestCredentialSurname",
					Age:          20,
					City:         "TestCredentialCity",
					Address:      "TestCredentialAddress",
				},
			},
			expectedAccount: &model.AccountDTO{},

			fc: func(data *test) {
				data.expectedAccount = &data.account
			},
		},
		{
			name:          "not found",
			accountName:   "some invalid account name",
			expectedError: errors.New("mongo: no documents in result"),
		},
		{
			name:          "empty name",
			expectedError: errors.New("empty account name"),
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
				testCase.expectedAccount.ID, err = repo.Create(ctx, testCase.account)
				assert.NoError(err)
			}

			foundedAccount, err := repo.FindByName(ctx, testCase.accountName)
			assert.Equal(testCase.expectedError, err)

			if testCase.isOK {
				assert.Equal(testCase.expectedAccount, foundedAccount)
			}

			_, err = repo.db.DeleteMany(ctx, bson.M{})
			assert.NoError(err)
		})
	}
}