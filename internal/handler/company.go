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

type companyRouter struct {
	*mux.Router
	services     *service.Services
	tokenManager auth.TokenManager
}

func newCompanyRouter(services *service.Services, manager auth.TokenManager) companyRouter {
	router := mux.NewRouter().PathPrefix(companyPathPrefix).Subrouter()
	handler := companyRouter{
		Router:       router,
		services:     services,
		tokenManager: manager,
	}

	handler.Path("/create").Methods(http.MethodPost).HandlerFunc(handler.createCompany)

	handler.Path("/find/name").Methods(http.MethodGet).HandlerFunc(handler.findByName)
	handler.Path("/find/URL").Methods(http.MethodGet).HandlerFunc(handler.findByURL)

	handler.Path("/update/name").Methods(http.MethodPut).HandlerFunc(handler.updateName)
	handler.Path("/update/description").Methods(http.MethodPut).HandlerFunc(handler.updateDescription)
	handler.Path("/update/URL").Methods(http.MethodPut).HandlerFunc(handler.updateURL)

	handler.Path("/delete").Methods(http.MethodDelete).HandlerFunc(handler.delete)

	return handler
}

type createCompanyRequest struct {
	request.CreateCompanyRequest
}

// Build builds request for create company.
func (req *createCompanyRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.CreateCompanyRequest)
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request for create company.
func (req *createCompanyRequest) Validate() error {

	if req.Name == "" {
		return fmt.Errorf("company name can not be nil")
	}

	if req.URL == "" {
		return fmt.Errorf("company URL can not be nil")
	}

	return nil
}

func (cr *companyRouter) createCompany(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req createCompanyRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONReturn(w, http.StatusBadRequest, err)
	}

	id, err := cr.services.CompanyService.Create(ctx, req.CreateCompanyRequest)
	if err != nil {
		middleware.JSONReturn(w, http.StatusInternalServerError, err)
	}

	if id != "" {
		middleware.JSONReturn(w, http.StatusOK, id)
	}
}

type findCompanyByNameRequest struct {
	request.FindCompanyByNameRequest
}

// Build builds request to find all companies by name.
func (req *findCompanyByNameRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request to find all companies by name.
func (req *findCompanyByNameRequest) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("company name can not be nil")
	}

	return nil
}

func (cr companyRouter) findByName(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req findCompanyByNameRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	company, err := cr.services.CompanyService.FindByName(ctx, req.FindCompanyByNameRequest)
	if err != nil {
		middleware.JSONError(w, http.StatusInternalServerError, err)
	}

	middleware.JSONReturn(w, http.StatusOK, company)
}

type findCompanyByURLRequest struct {
	request.FindCompanyByURLRequest
}

// Build builds request to find all companies by URL.
func (req *findCompanyByURLRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request to find all companies by URL.
func (req *findCompanyByURLRequest) Validate() error {
	if req.URL == "" {
		return fmt.Errorf("comapny URL can not be nil")
	}

	return nil
}

func (cr companyRouter) findByURL(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req findCompanyByURLRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	companies, err := cr.services.CompanyService.FindByURL(ctx, req.FindCompanyByURLRequest)
	if err != nil {
		middleware.JSONError(w, http.StatusInternalServerError, err)
	}

	middleware.JSONReturn(w, http.StatusOK, companies)
}

type updateCompanyNameRequest struct {
	request.UpdateCompanyNameRequest
}

// Build builds request to update company's name.
func (req *updateCompanyNameRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request to update company's name.
func (req *updateCompanyNameRequest) Validate() error {

	if req.ID == "" {
		return fmt.Errorf("company ID can not be nil")
	}

	if req.NewName == "" {
		return fmt.Errorf("comapny name can not be nil")
	}

	return nil
}

func (cr *companyRouter) updateName(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req updateCompanyNameRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	id, err := cr.services.CompanyService.UpdateName(ctx, req.UpdateCompanyNameRequest)
	if err != nil {
		middleware.JSONError(w, http.StatusInternalServerError, err)
	}

	if id != "" {
		middleware.JSONReturn(w, http.StatusOK, id)
	}
}

type updateCompanyDescriptionRequest struct {
	request.UpdateCompanyDescriptionRequest
}

// Build builds request to update company's description.
func (req *updateCompanyDescriptionRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request to update company's description.
func (req *updateCompanyDescriptionRequest) Validate() error {

	if req.ID == "" {
		return fmt.Errorf("company ID can not be nil")
	}

	if req.NewDescription == "" {
		return fmt.Errorf("comapny description can not be nil")
	}

	return nil
}

func (cr *companyRouter) updateDescription(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req updateCompanyDescriptionRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	id, err := cr.services.CompanyService.UpdateDescription(ctx, req.UpdateCompanyDescriptionRequest)
	if err != nil {
		middleware.JSONError(w, http.StatusInternalServerError, err)
	}

	if id != "" {
		middleware.JSONReturn(w, http.StatusOK, id)
	}
}

type updateCompanyURLRequest struct {
	request.UpdateCompanyURLRequest
}

// Build builds request to update company's URL.
func (req *updateCompanyURLRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request to update company's URL.
func (req *updateCompanyURLRequest) Validate() error {

	if req.ID == "" {
		return fmt.Errorf("company ID can not be nil")
	}

	if req.NewURL == "" {
		return fmt.Errorf("comapny URL can not be nil")
	}

	return nil
}

func (cr *companyRouter) updateURL(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req updateCompanyURLRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	id, err := cr.services.CompanyService.UpdateURL(ctx, req.UpdateCompanyURLRequest)
	if err != nil {
		middleware.JSONError(w, http.StatusInternalServerError, err)
	}

	if id != "" {
		middleware.JSONReturn(w, http.StatusOK, id)
	}
}

type deleteCompanyRequest struct {
	request.DeleteCompanyRequest
}

// Build builds request to delete company.
func (req *deleteCompanyRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request to delete company.
func (req *deleteCompanyRequest) Validate() error {
	if req.ID == "" {
		return fmt.Errorf("company ID can not be nil")
	}

	return nil
}

func (cr *companyRouter) delete(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var req deleteCompanyRequest

	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, http.StatusBadRequest, err)
	}

	id, err := cr.services.CompanyService.Delete(ctx, req.DeleteCompanyRequest)
	if err != nil {
		middleware.JSONError(w, http.StatusInternalServerError, err)
	}

	if id != "" {
		middleware.JSONReturn(w, http.StatusOK, id)
	}
}
