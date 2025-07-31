package companyusecases

import (
	"context"
	"fmt"

	companydomain "github.com/joshuaalpuerto/go-rest-api/internal/api/company/domain"
)

func (c *CompanyUsecase) CreateCompany(nc NewCompany) (*companydomain.Company, error) {
	// this should come from request
	userId := "c8d9c08f-4f87-4c5c-8862-2f4abac75f1f"

	// Validate if incoming payload adhere to the domain model
	companyEntity, err := nc.ToDomainEntity(userId)
	if err != nil {
		return nil, fmt.Errorf("Invalid company: [%+v]: %w", nc, err)
	}

	result, err := c.companyRepository.Create(context.Background(), companyEntity)
	if err != nil {
		return nil, err
	}

	company, err := result.ToDomain()
	if err != nil {
		return nil, err
	}
	return &company, nil
}
