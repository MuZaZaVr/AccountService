package service

import (
	"context"
	"github.com/MuZaZaVr/account-service/internal/model"
	"github.com/MuZaZaVr/account-service/internal/model/request"
	m "github.com/MuZaZaVr/account-service/internal/service/mock"
	"github.com/pkg/errors"
	assertTest "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestCompanyService_Create(t *testing.T) {
	assert := assertTest.New(t)

	type test struct {
		name     string
		req      request.CreateCompanyRequest
		fn       func(company *m.Company, data test)
		expID    string
		expError error
	}

	testTable := []test{
		{
			name: "all OK",
			req: request.CreateCompanyRequest{
				Name:        "Amazon",
				Description: "Some description",
				URL:         "https://amazon.com",
			},
			expID: primitive.NewObjectID().Hex(),
			fn: func(company *m.Company, data test) {
				company.On("Create", mock.Anything, model.CompanyDTO{
					Name:        data.req.Name,
					Description: data.req.Description,
					URL:         data.req.URL,
				}).Return(data.expID, nil)
			},
		},
		{
			name: "errors",
			req: request.CreateCompanyRequest{
				Name:        "Amazon",
				Description: "Some description",
				URL:         "https://amazon.com",
			},
			fn: func(company *m.Company, data test) {
				company.On("Create", mock.Anything, model.CompanyDTO{
					Name:        data.req.Name,
					Description: data.req.Description,
					URL:         data.req.URL,
				}).Return(data.expID, errors.New(""))
			},
			expError: errors.Wrap(errors.New(""), "cant create company"),
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			companyRepoMock := new(m.Company)
			ctx := context.Background()
			service := NewCompanyService(companyRepoMock)

			if tc.fn != nil {
				tc.fn(companyRepoMock, tc)
			}

			id, err := service.Create(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expError.Error(), err.Error())
			}

			assert.Equal(tc.expID, id)
		})
	}
}

func TestCompanyService_Delete(t *testing.T) {
	assert := assertTest.New(t)

	type test struct {
		name   string
		isOk   bool
		req    request.DeleteCompanyRequest
		fn     func(company *m.Company, data test)
		expID  string
		expErr error
	}

	testTable := []test{
		{
			name: "all OK",
			isOk: true,

			req:   request.DeleteCompanyRequest{},
			expID: primitive.NewObjectID().Hex(),

			fn: func(company *m.Company, data test) {
				company.On("Delete", mock.Anything, data.req.ID).Return(data.expID, nil)
			},
		},
		{
			name: "errors",

			req:    request.DeleteCompanyRequest{ID: ""},
			expErr: errors.Wrap(errors.New(""), "can't delete company"),

			fn: func(company *m.Company, data test) {
				company.On("Delete", mock.Anything, data.req.ID).Return("", errors.New(""))
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			companyRepoMock := new(m.Company)
			service := NewCompanyService(companyRepoMock)

			if tc.isOk {
				tc.expID = tc.req.ID
			}

			if tc.fn != nil {
				tc.fn(companyRepoMock, tc)
			}

			id, err := service.Delete(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}

			assert.Equal(tc.expID, id)
		})
	}
}

func TestCompanyService_FindByName(t *testing.T) {
	assert := assertTest.New(t)

	type test struct {
		name       string
		isOk       bool
		req        request.FindCompanyByNameRequest
		fn         func(company *m.Company, data test)
		expCompany *model.CompanyDTO
		expErr     error
	}

	testTable := []test{
		{
			name: "all OK",
			isOk: true,
			req:  request.FindCompanyByNameRequest{},

			expCompany: &model.CompanyDTO{
				ID:          "",
				Name:        "companyName",
				Description: "companyDescription",
				URL:         "companyURL",
			},

			fn: func(company *m.Company, data test) {
				company.On("FindByName", mock.Anything, data.req.Name).Return(data.expCompany, nil)
			},
		},
		{
			name: "errors",
			req:  request.FindCompanyByNameRequest{},

			expCompany: nil,
			expErr:     errors.Wrap(errors.New(""), "can't find company by name"),

			fn: func(company *m.Company, data test) {
				company.On("FindByName", mock.Anything, data.req.Name).Return(nil, errors.New(""))
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			companyRepoMock := new(m.Company)
			service := NewCompanyService(companyRepoMock)

			if tc.isOk {
				tc.req.Name = tc.expCompany.Name
			}

			if tc.fn != nil {
				tc.fn(companyRepoMock, tc)
			}

			companyResult, err := service.FindByName(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}

			assert.Equal(tc.expCompany, companyResult)
		})
	}
}

func TestCompanyService_FindByURL(t *testing.T) {
	assert := assertTest.New(t)

	type test struct {
		name       string
		isOk       bool
		req        request.FindCompanyByURLRequest
		fn         func(company *m.Company, data test)
		expCompany *model.CompanyDTO
		expErr     error
	}

	testTable := []test{
		{
			name: "all OK",
			isOk: true,
			req:  request.FindCompanyByURLRequest{},

			expCompany: &model.CompanyDTO{
				Name:        "companyName",
				Description: "companyDescription",
				URL:         "companyURL",
			},

			fn: func(company *m.Company, data test) {
				company.On("FindByURL", mock.Anything, data.req.URL).Return(data.expCompany, nil)
			},
		},
		{
			name: "errors",
			req:  request.FindCompanyByURLRequest{},

			expCompany: nil,
			expErr:     errors.Wrap(errors.New(""), "can't find company by URL"),

			fn: func(company *m.Company, data test) {
				company.On("FindByURL", mock.Anything, data.req.URL).Return(nil, errors.New(""))
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			companyRepoMock := new(m.Company)
			service := NewCompanyService(companyRepoMock)

			if tc.isOk {
				tc.req.URL = tc.expCompany.URL
			}

			if tc.fn != nil {
				tc.fn(companyRepoMock, tc)
			}

			company, err := service.FindByURL(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}

			assert.Equal(tc.expCompany, company)
		})
	}
}

func TestCompanyService_UpdateName(t *testing.T) {
	assert := assertTest.New(t)

	type test struct {
		name   string
		isOk   bool
		req    request.UpdateCompanyNameRequest
		fn     func(company *m.Company, data test)
		expID  string
		expErr error
	}

	testTable := []test{
		{
			name: "all OK",
			isOk: true,
			req: request.UpdateCompanyNameRequest{
				NewName: "someName",
			},

			expID: primitive.NewObjectID().Hex(),

			fn: func(company *m.Company, data test) {
				company.On("UpdateName", mock.Anything, data.req.ID, data.req.NewName).Return(data.expID, nil)
			},
		},
		{
			name: "errors",
			req: request.UpdateCompanyNameRequest{
				NewName: "someName",
			},

			expErr: errors.Wrap(errors.New(""), "can't update company's name"),

			fn: func(company *m.Company, data test) {
				company.On("UpdateName", mock.Anything, data.req.ID, data.req.NewName).Return("", errors.New(""))
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			companyRepoMock := new(m.Company)
			service := NewCompanyService(companyRepoMock)

			if tc.isOk {
				tc.req.ID = tc.expID
			}

			if tc.fn != nil {
				tc.fn(companyRepoMock, tc)
			}

			company, err := service.UpdateName(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}

			assert.Equal(tc.expID, company)
		})
	}
}

func TestCompanyService_UpdateDescription(t *testing.T) {
	assert := assertTest.New(t)

	type test struct {
		name   string
		isOk   bool
		req    request.UpdateCompanyDescriptionRequest
		fn     func(company *m.Company, data test)
		expID  string
		expErr error
	}

	testTable := []test{
		{
			name: "all OK",
			isOk: true,
			req: request.UpdateCompanyDescriptionRequest{
				NewDescription: "someDescription",
			},

			expID: primitive.NewObjectID().Hex(),

			fn: func(company *m.Company, data test) {
				company.On("UpdateDescription", mock.Anything, data.req.ID, data.req.NewDescription).Return(data.expID, nil)
			},
		},
		{
			name: "errors",
			req: request.UpdateCompanyDescriptionRequest{
				NewDescription: "someDescription",
			},

			expErr: errors.Wrap(errors.New(""), "can't update company's description"),

			fn: func(company *m.Company, data test) {
				company.On("UpdateDescription", mock.Anything, data.req.ID, data.req.NewDescription).Return("", errors.New(""))
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			companyRepoMock := new(m.Company)
			service := NewCompanyService(companyRepoMock)

			if tc.isOk {
				tc.req.ID = tc.expID
			}

			if tc.fn != nil {
				tc.fn(companyRepoMock, tc)
			}

			company, err := service.UpdateDescription(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}

			assert.Equal(tc.expID, company)
		})
	}
}

func TestCompanyService_UpdateURL(t *testing.T) {
	assert := assertTest.New(t)

	type test struct {
		name   string
		isOk   bool
		req    request.UpdateCompanyURLRequest
		fn     func(company *m.Company, data test)
		expID  string
		expErr error
	}

	testTable := []test{
		{
			name: "all OK",
			isOk: true,
			req: request.UpdateCompanyURLRequest{
				NewURL: "someURL",
			},

			expID: primitive.NewObjectID().Hex(),

			fn: func(company *m.Company, data test) {
				company.On("UpdateURL", mock.Anything, data.req.ID, data.req.NewURL).Return(data.expID, nil)
			},
		},
		{
			name: "errors",
			req: request.UpdateCompanyURLRequest{
				NewURL: "someURL",
			},

			expErr: errors.Wrap(errors.New(""), "can't update company's URL"),

			fn: func(company *m.Company, data test) {
				company.On("UpdateURL", mock.Anything, data.req.ID, data.req.NewURL).Return("", errors.New(""))
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			companyRepoMock := new(m.Company)
			service := NewCompanyService(companyRepoMock)

			if tc.isOk {
				tc.req.ID = tc.expID
			}

			if tc.fn != nil {
				tc.fn(companyRepoMock, tc)
			}

			company, err := service.UpdateURL(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}

			assert.Equal(tc.expID, company)
		})
	}
}

func TestCompanyService_IsExist(t *testing.T) {
	assert := assertTest.New(t)

	type test struct {
		name      string
		isOk      bool
		req       request.IsCompanyExistRequest
		fn        func(company *m.Company, data test)
		expResult bool
		expErr    error
	}

	testTable := []test{
		{
			name: "all OK",
			isOk: true,
			req:  request.IsCompanyExistRequest{},

			expResult: true,

			fn: func(company *m.Company, data test) {
				company.On("IsExist", mock.Anything, data.req.Name).Return(data.expResult, nil)
			},
		},
		{
			name: "errors",
			req:  request.IsCompanyExistRequest{},

			expResult: false,
			expErr:    errors.Wrap(errors.New(""), "can't confirm company's existence"),

			fn: func(company *m.Company, data test) {
				company.On("IsExist", mock.Anything, data.req.Name).Return(data.expResult, errors.New(""))
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			companyRepoMock := new(m.Company)
			service := NewCompanyService(companyRepoMock)

			if tc.isOk {
				tc.req.Name = "someName"
			}

			if tc.fn != nil {
				tc.fn(companyRepoMock, tc)
			}

			result, err := service.IsExist(ctx, tc.req)
			if err != nil {
				assert.Equal(tc.expErr.Error(), err.Error())
			}

			assert.Equal(tc.expResult, result)
		})
	}
}
