package handler

import (
	"github.com/MuZaZaVr/account-service/internal/service"
	"github.com/MuZaZaVr/account-service/pkg/auth"
	"github.com/gorilla/mux"
)

const (
	accountPathPrefix = "/account"
	companyPathPrefix = "/company"
)

type API struct {
	*mux.Router
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *API {
	api := API{
		mux.NewRouter(),
	}

	api.PathPrefix(accountPathPrefix).Handler(newCompanyRouter(services, tokenManager))
	api.PathPrefix(companyPathPrefix).Handler(newAccountHandler(services, tokenManager))

	return &api
}
