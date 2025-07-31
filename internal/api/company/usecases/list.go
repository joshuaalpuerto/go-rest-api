package companyusecases

import (
	"context"
	"fmt"

	companydomain "github.com/joshuaalpuerto/go-rest-api/internal/api/company/domain"
)

// Business logic methods
func (c *CompanyUsecase) GetAllCompanies() ([]companydomain.Company, error) {
	result, err := c.companyRepository.FindAll(context.Background())
	if err != nil {
		return nil, fmt.Errorf("FindAll: %w", err)
	}
	companies, err := companydomain.ToCompanyEntities(result)
	if err != nil {
		return nil, fmt.Errorf("FindAll: DB to Domain: %w", err)
	}
	return companies, nil
}

func (c *CompanyUsecase) GetCompanyByID(id string) (*companydomain.Company, error) {
	result, err := c.companyRepository.FindOneByID(context.Background(), id)
	if err != nil {
		if err == companydomain.ErrNotFound {
			return nil, fmt.Errorf("Company [%s] %w", id, err)
		}
		return nil, fmt.Errorf("FindOneByID: %w", err)
	}
	company, err := result.ToDomain()
	if err != nil {
		return nil, fmt.Errorf("FindOneByID: DB to Domain: %w", err)
	}
	return &company, nil
}
