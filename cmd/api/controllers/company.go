package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/joshuaalpuerto/go-rest-api/cmd/api/response"
	companyusecases "github.com/joshuaalpuerto/go-rest-api/internals/api/company/usecases"
)

type Validator interface {
	Validate(model any) error
}

type CompanyHandler struct {
	companyService companyusecases.CompanyUsecase
	validator      Validator
}

func NewCompanyHandler(companyRepository companyusecases.CompanyRepository, validator Validator) CompanyHandler {
	return CompanyHandler{
		companyService: companyusecases.NewCompanyUsecase(companyRepository),
		validator:      validator,
	}
}

func (h CompanyHandler) GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := h.companyService.GetAllCompanies()
	response := response.Response{}

	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert domain entities to DTOs
	companiesDTO := companyusecases.ToAppCompanies(companies)
	response.SendSuccessResponse(w, companiesDTO, http.StatusOK)
}

func (h CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company companyusecases.NewCompany
	response := response.Response{}

	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errs := h.validator.Validate(&company); errs != nil {
		response.SendErrorResponse(w, errs.Error(), http.StatusBadRequest)
		return
	}

	createdCompany, err := h.companyService.CreateCompany(company)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, createdCompany, http.StatusCreated)
}
