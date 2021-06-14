package handler

import (
	"github.com/MuZaZaVr/account-service/internal/service"
	"github.com/MuZaZaVr/account-service/pkg/auth"
	"github.com/gorilla/mux"
)

type accountRouter struct {
	*mux.Router
	services *service.Services
	manager auth.TokenManager
}

func newAccountHandler(services *service.Services, manager auth.TokenManager) *accountRouter {
	router := mux.NewRouter().PathPrefix(accountPathPrefix).Subrouter()

	handler := accountRouter{
		Router:   router,
		services: services,
		manager:  manager,
	}

	return &handler
}
