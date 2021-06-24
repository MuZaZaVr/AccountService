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

type accountRouter struct {
	*mux.Router
	services *service.Services
	manager  auth.TokenManager
}

func newAccountHandler(services *service.Services, manager auth.TokenManager) accountRouter {
	router := mux.NewRouter().PathPrefix(accountPathPrefix).Subrouter()

	handler := accountRouter{
		Router:   router,
		services: services,
		manager:  manager,
	}

	handler.Path("/create").Methods(http.MethodPost).HandlerFunc(handler.createAccount)

	handler.Path("/find/name").Methods(http.MethodGet).HandlerFunc(handler.findAccountByName)

	handler.Path("/delete").Methods(http.MethodDelete).HandlerFunc(handler.deleteAccount)

	return handler
}

type createAccountRequest struct {
	request.CreateAccountRequest
}

func (req *createAccountRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.CreateAccountRequest)
	if err != nil {
		return err
	}

	return nil
}

func (req createAccountRequest) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("company name can not be nil")
	}

	if req.CompanyID == "" {
		return fmt.Errorf("not valid company ID")
	}

	if req.UserID < 0 && req.UserID == 0 {
		return fmt.Errorf("not valid user ID")
	}

	return nil
}

func (ar *accountRouter) createAccount(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req createAccountRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	id, err := ar.services.AccountService.Create(ctx, req.CreateAccountRequest)
	if err != nil {
		middleware.JSONError(w, http.StatusInternalServerError, err)
	}

	middleware.JSONReturn(w, http.StatusOK, id)
}

type findByNameRequest struct {
	request.FindAccountByNameRequest
}

func (req *findByNameRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.FindAccountByNameRequest)
	if err != nil {
		return err
	}

	return nil
}

func (req findByNameRequest) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("company name can not be nil")
	}

	return nil
}

func (ar *accountRouter) findAccountByName(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req findByNameRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	id, err := ar.services.AccountService.FindByName(ctx, req.FindAccountByNameRequest)
	if err != nil {
		middleware.JSONError(w, http.StatusInternalServerError, err)
	}

	middleware.JSONReturn(w, http.StatusOK, id)
}

type deleteAccountRequest struct {
	request.DeleteAccountRequest
}

func (req *deleteAccountRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	return nil
}

func (req *deleteAccountRequest) Validate() error {
	if req.ID == "" {
		return fmt.Errorf("account ID can not be nil")
	}

	return nil
}

func (ar *accountRouter) deleteAccount(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req deleteAccountRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	id, err := ar.services.AccountService.Delete(ctx, req.DeleteAccountRequest)
	if err != nil {
		middleware.JSONError(w, http.StatusInternalServerError, err)
	}

	middleware.JSONReturn(w, http.StatusOK, id)
}