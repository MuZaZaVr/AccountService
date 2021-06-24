package handler

import (
	"github.com/MuZaZaVr/account-service/internal/service"
	"github.com/MuZaZaVr/account-service/pkg/auth"
	"github.com/gorilla/mux"
)

const (
	accountPathPrefix = "/account"
	companyPathPrefix = "/company"
	credentialPathPrefix = "/credential"
)

type API struct {
	*mux.Router
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *API {
	api := API{
		mux.NewRouter(),
	}

	api.PathPrefix(companyPathPrefix).
		Handler(newCompanyRouter(services, tokenManager)).
		Name("companyRouter")

	api.PathPrefix(accountPathPrefix).
		Handler(newAccountHandler(services, tokenManager)).
		Name("accountRouter")

	api.PathPrefix(credentialPathPrefix).
		Handler(NewCredentialRouter(services, tokenManager)).
		Name("credentialRouter")
	return &api
}
