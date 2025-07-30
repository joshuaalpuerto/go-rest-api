package companyusecases

import (
	"context"

	companydomain "github.com/joshuaalpuerto/go-rest-api/internals/api/company/domain"
)

type CompanyRepository interface {
	FindAll(ctx context.Context) ([]companydomain.CompanyDB, error)
	FindOne(ctx context.Context, id string) (*companydomain.CompanyDB, error)
	Create(ctx context.Context, company companydomain.NewCompany) (*companydomain.CompanyDB, error)
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
