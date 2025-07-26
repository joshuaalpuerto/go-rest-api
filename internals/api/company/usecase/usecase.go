package companyusecase

import (
	"context"

	companydomain "github.com/joshuaalpuerto/go-rest-api/internals/api/company/domain"
)

type CompanyRepository interface {
	FindAll(ctx context.Context) ([]companydomain.CompanyDB, error)
	FindOne(ctx context.Context, id string) (*companydomain.CompanyDB, error)
	Create(ctx context.Context, company companydomain.Company) (*companydomain.CompanyDB, error)
	Update(ctx context.Context, company companydomain.Company) (*companydomain.CompanyDB, error)
	Delete(ctx context.Context, id string) (*companydomain.CompanyDB, error)
}

type CompanyUsecase struct {
	companyRepository CompanyRepository
}

func NewCompanyUsecase(companyRepository CompanyRepository) CompanyUsecase {
	return CompanyUsecase{
		companyRepository: companyRepository,
	}
}

// Business logic methods
func (c *CompanyUsecase) GetAllCompanies() ([]companydomain.Company, error) {
	companies, err := c.companyRepository.FindAll(context.Background())
	if err != nil {
		return nil, err
	}
	companiesDomain, err := companydomain.ToCompaniesDomain(companies)
	if err != nil {
		return nil, err
	}
	return companiesDomain, nil
}

func (c *CompanyUsecase) GetCompanyByID(id string) (*companydomain.Company, error) {
	company, err := c.companyRepository.FindOne(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, nil
	}
	companyDomain, err := companydomain.ToCompanyDomain(*company)
	if err != nil {
		return nil, err
	}
	return &companyDomain, nil
}

func (c *CompanyUsecase) CreateCompany(company companydomain.Company) (*companydomain.Company, error) {
	//TODO: how can we create validation here?
	companyDB, err := c.companyRepository.Create(context.Background(), company)
	if err != nil {
		return nil, err
	}
	companyDomain, err := companydomain.ToCompanyDomain(*companyDB)
	if err != nil {
		return nil, err
	}
	return &companyDomain, nil
}

func (c *CompanyUsecase) UpdateCompany(company companydomain.Company) (*companydomain.Company, error) {
	companyDB, err := c.companyRepository.Update(context.Background(), company)
	if err != nil {
		return nil, err
	}
	companyDomain, err := companydomain.ToCompanyDomain(*companyDB)
	if err != nil {
		return nil, err
	}
	return &companyDomain, nil
}
