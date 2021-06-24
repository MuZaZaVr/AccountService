package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MuZaZaVr/account-service/internal/model/request"
	"github.com/MuZaZaVr/account-service/internal/service"
	"github.com/MuZaZaVr/account-service/pkg/auth"
	"github.com/MuZaZaVr/account-service/pkg/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type credentialRouter struct {
	*mux.Router
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewCredentialRouter(services *service.Services, manager auth.TokenManager) credentialRouter {
	router := mux.NewRouter().PathPrefix(credentialPathPrefix).Subrouter()
	handler := credentialRouter{
		Router:       router,
		services:     services,
		tokenManager: manager,
	}

	handler.Path("/create").Methods(http.MethodPost).HandlerFunc(handler.create)

	return handler
}

type createCredentialRequest struct {
	request.CreateCredentialRequest
}

func (req *createCredentialRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	return nil
}

func (req *createCredentialRequest) Validate() error {
	if req.Password == "" {
		return fmt.Errorf("credential password can not be nil")
	}

	if req.Email == "" {
		return fmt.Errorf("credential email can not be nil")
	}

	if req.Login == "" {
		return fmt.Errorf("credential login can not be nil")
	}

	return nil
}

func (cr *credentialRouter) create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req createCredentialRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	id, err := cr.services.CredentialService.Create(ctx, req.CreateCredentialRequest)
	if err != nil {
		middleware.JSONError(w, http.StatusInternalServerError, err)
	}

	middleware.JSONReturn(w, http.StatusOK, id)
}