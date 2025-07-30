package companyusecases

import (
	"context"
	"fmt"

	companydomain "github.com/joshuaalpuerto/go-rest-api/internals/api/company/domain"
)

func (c *CompanyUsecase) CreateCompany(nc NewCompany) (*companydomain.Company, error) {
	// this should come from request
	userId := "c8d9c08f-4f87-4c5c-8862-2f4abac75f1f"
	// this is validated
	domainCompany, err := nc.ToDomainCompany(userId)
	if err != nil {
		return nil, fmt.Errorf("Invalid company: [%+v]: %w", nc, err)
	}

	companyDB, err := c.companyRepository.Create(context.Background(), domainCompany)
	if err != nil {
		return nil, err
	}
	companyDomain, err := companydomain.ToCompanyDomain(*companyDB)
	if err != nil {
		return nil, err
	}
	return &companyDomain, nil
}
