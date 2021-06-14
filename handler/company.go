package handler

import (
	"github.com/MuZaZaVr/account-service/internal/service"
	"github.com/MuZaZaVr/account-service/pkg/auth"
	"github.com/gorilla/mux"
)

type companyRouter struct {
	*mux.Router
	services *service.Services
	tokenManager auth.TokenManager
}

func newCompanyRouter(services *service.Services, manager auth.TokenManager) companyRouter {
	router := mux.NewRouter().PathPrefix(companyPathPrefix).Subrouter()
	handler := companyRouter{
		Router:       router,
		services:     services,
		tokenManager: manager,
	}

	return handler
}