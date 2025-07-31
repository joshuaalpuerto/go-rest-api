package companydomain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrNameNotValid = errors.New("name is not valid")
)

// represent the company entity in the domain
type Company struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}

type CompanyDB struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedBy string    `db:"created_by"`
	UpdatedBy string    `db:"updated_by"`
}

func (c CompanyDB) ToDomain() (Company, error) {
	// we can use type conversation because properties are the same
	return Company(c), nil
}

func ToCompanyEntities(dbs []CompanyDB) ([]Company, error) {
	companies := make([]Company, len(dbs))
	for i, db := range dbs {
		company, err := db.ToDomain()
		if err != nil {
			return nil, err
		}
		companies[i] = company
	}
	return companies, nil
}

type NewCompany struct {
	Name      string
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ToDomainCompany converts PostCompany to domain Company
func (p *NewCompany) ToDomainEntity(userId string) (NewCompany, error) {
	return NewCompany{
		Name:      p.Name,
		CreatedBy: uuid.MustParse(userId),
		UpdatedBy: uuid.MustParse(userId),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
