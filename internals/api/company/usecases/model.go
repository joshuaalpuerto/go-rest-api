package companyusecases

import (
	"time"

	"github.com/google/uuid"
	companydomain "github.com/joshuaalpuerto/go-rest-api/internals/api/company/domain"
)

// Contract of what company looks like to the client
type CompanyDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
}

func ToAppCompany(c companydomain.Company) CompanyDTO {
	return CompanyDTO{
		ID:        c.ID.String(),
		Name:      c.Name,
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
		UpdatedAt: c.UpdatedAt.Format(time.RFC3339),
		CreatedBy: c.CreatedBy,
		UpdatedBy: c.UpdatedBy,
	}
}

func ToAppCompanies(c []companydomain.Company) []CompanyDTO {
	if len(c) == 0 {
		return []CompanyDTO{}
	}

	companies := make([]CompanyDTO, len(c))
	for i, v := range c {
		companies[i] = ToAppCompany(v)
	}
	return companies
}

type NewCompany struct {
	Name string `json:"name" validate:"required"`
}

// ToDomainCompany converts PostCompany to domain Company
func (p *NewCompany) ToDomainCompany(userId string) (companydomain.NewCompany, error) {
	return companydomain.NewCompany{
		Name:      p.Name,
		CreatedBy: uuid.MustParse(userId),
		UpdatedBy: uuid.MustParse(userId),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
