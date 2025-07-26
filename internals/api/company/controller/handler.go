package companycontroller

import (
	"net/http"

	companyusecase "github.com/joshuaalpuerto/go-rest-api/internals/api/company/usecase"
)

type CompanyHandler struct {
	companyService companyusecase.CompanyUsecase
}

func NewCompanyHandler(companyRepository companyusecase.CompanyRepository) CompanyHandler {
	return CompanyHandler{
		companyService: companyusecase.NewCompanyUsecase(companyRepository),
	}
}

func (h *CompanyHandler) GetAllCompanies(w http.ResponseWriter, r *http.Request) (any, error) {
	companies, err := h.companyService.GetAllCompanies()
	if err != nil {

		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// throw internal server error here
		return nil, err
	}

	// Convert domain entities to DTOs
	response := companyusecase.ToAppCompanies(companies)

	return response, nil
}
