package companyhttp

import (
	"encoding/json"
	"net/http"

	response "github.com/joshuaalpuerto/go-rest-api/internal/common/http/response"
	commonvalidator "github.com/joshuaalpuerto/go-rest-api/internal/common/validator"
	companydomain "github.com/joshuaalpuerto/go-rest-api/internal/company/domain"
	companyusecases "github.com/joshuaalpuerto/go-rest-api/internal/company/usecases"
)

type CompanyHandler struct {
	companyService companyusecases.CompanyUsecase
	validator      commonvalidator.Validator
}

func NewCompanyHandler(companyRepository companyusecases.CompanyRepository, validator commonvalidator.Validator) CompanyHandler {
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
	companiesDTO := ToAppCompanies(companies)
	response.SendSuccessResponse(w, companiesDTO, http.StatusOK)
}

func (h CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var newCompany NewCompany
	response := response.Response{}

	if err := json.NewDecoder(r.Body).Decode(&newCompany); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errs := h.validator.Validate(&newCompany); errs != nil {
		response.SendErrorResponse(w, errs.Error(), http.StatusBadRequest)
		return
	}

	createdCompany, err := h.companyService.CreateCompany(companydomain.NewCompany{
		Name: newCompany.Name,
	})
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, createdCompany, http.StatusCreated)
}

func (h CompanyHandler) GetCompanyByID(w http.ResponseWriter, r *http.Request) {
	companyID := r.PathValue("id")

	company, err := h.companyService.GetCompanyByID(companyID)
	response := response.Response{}

	if err != nil {
		if err == companydomain.ErrNotFound {
			response.SendErrorResponse(w, err.Error(), http.StatusNotFound)
		} else {
			response.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	response.SendSuccessResponse(w, company, http.StatusOK)
}
