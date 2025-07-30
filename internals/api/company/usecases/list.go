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
		return nil, fmt.Errorf("FindAll: %w", err)
	}
	companiesDomain, err := companydomain.ToCompaniesDomain(companies)
	if err != nil {
		return nil, fmt.Errorf("FindAll: DB to Domain: %w", err)
	}
	return companiesDomain, nil
}

func (c *CompanyUsecase) GetCompanyByID(id string) (*companydomain.Company, error) {
	company, err := c.companyRepository.FindOneByID(context.Background(), id)
	if err != nil {
		if err == companydomain.ErrNotFound {
			return nil, fmt.Errorf("Company [%s] %w", id, err)
		}
		return nil, fmt.Errorf("FindOneByID: %w", err)
	}
	companyDomain, err := companydomain.ToCompanyDomain(*company)
	if err != nil {
		return nil, fmt.Errorf("FindOneByID: DB to Domain: %w", err)
	}
	return &companyDomain, nil
}
