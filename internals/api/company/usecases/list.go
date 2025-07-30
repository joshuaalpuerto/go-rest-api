package companyusecases

import (
	"context"
	"fmt"

	companydomain "github.com/joshuaalpuerto/go-rest-api/internals/api/company/domain"
)

// Business logic methods
func (c *CompanyUsecase) GetAllCompanies() ([]companydomain.Company, error) {
	companies, err := c.companyRepository.FindAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Company service: %w", err)
	}
	companiesDomain, err := companydomain.ToCompaniesDomain(companies)
	if err != nil {
		return nil, fmt.Errorf("Company service: %w", err)
	}
	return companiesDomain, nil
}
