package companydomain

import (
	"context"
)

type CompanyRepository interface {
	FindAll(ctx context.Context) ([]CompanyDB, error)
	FindOne(ctx context.Context, id string) (*CompanyDB, error)
	Create(ctx context.Context, company Company) (*CompanyDB, error)
	Update(ctx context.Context, company Company) (*CompanyDB, error)
	Delete(ctx context.Context, id string) (*CompanyDB, error)
}
