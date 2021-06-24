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

	handler.Path("/find/one/name").Methods(http.MethodGet).HandlerFunc(handler.findAccountByName)

	handler.Path("/").Methods(http.MethodPost).HandlerFunc(handler.createAccount)
	handler.Path("/").Methods(http.MethodPut).HandlerFunc(handler.updateAccount)
	handler.Path("/").Methods(http.MethodDelete).HandlerFunc(handler.deleteAccount)

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

	if req.UserID < 0 && req.UserID == 0 {
		return fmt.Errorf("not valid user ID")
	}

	return nil
}

// @Summary Create account
// @Tags account
// @Description Create account
// @Accept  json
// @Produce  json
// @Param purchase body request.CreateAccountRequest true "Account"
// @Success 200 {string} string id
// @Failure 400 {object} middleware.SwagError
// @Failure 500 {object} middleware.SwagError
// @Router /account/ [post]
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
		return fmt.Errorf("account name can not be nil")
	}

	return nil
}

// @Summary Find account by name
// @Tags account
// @Description Find account by provided account name
// @Accept  json
// @Produce  json
// @Param purchase body request.FindAccountByNameRequest true "Account"
// @Success 200 {string} string id
// @Failure 400 {object} middleware.SwagError
// @Failure 500 {object} middleware.SwagError
// @Router /account/ [put]
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

type updateAccountRequest struct {
	request.UpdateAccountRequest
}

func (req *updateAccountRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.UpdateAccountRequest)
	if err != nil {
		return err
	}

	return nil
}

func (req *updateAccountRequest) Validate() error {
	if req.ID == "" {
		return fmt.Errorf("account ID can not be nil")
	}

	if req.UpdatedAccount.UserId < 0 && req.UpdatedAccount.UserId == 0 {
		return fmt.Errorf("not valid user ID")
	}

	return nil
}

// @Summary Delete account
// @Tags account
// @Description Update account
// @Accept  json
// @Produce  json
// @Param purchase body request.UpdateAccountRequest true "Account"
// @Success 200 {string} string id
// @Failure 400 {object} middleware.SwagError
// @Failure 500 {object} middleware.SwagError
// @Router /account/ [put]
func (ar *accountRouter) updateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req updateAccountRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	id, err := ar.services.AccountService.Update(ctx, req.UpdateAccountRequest)
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

// @Summary Delete
// @Tags account
// @Description Delete account
// @Accept  json
// @Produce  json
// @Param purchase body request.DeleteAccountRequest true "Account"
// @Success 200 {string} string id
// @Failure 400 {object} middleware.SwagError
// @Failure 500 {object} middleware.SwagError
// @Router /account/ [delete]
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