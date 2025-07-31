package companyusecases

import (
	"context"

	companydomain "github.com/joshuaalpuerto/go-rest-api/internal/company/domain"
)

func (c *CompanyUsecase) CreateCompany(nc companydomain.NewCompany) (*companydomain.Company, error) {
	userId := "c8d9c08f-4f87-4c5c-8862-2f4abac75f1f"
	nc, err := nc.ToDomainEntity(userId)
	if err != nil {
		return nil, err
	}

	result, err := c.companyRepository.Create(context.Background(), nc)
	if err != nil {
		return nil, err
	}

	company, err := result.ToDomain()
	if err != nil {
		return nil, err
	}
	return &company, nil
}
