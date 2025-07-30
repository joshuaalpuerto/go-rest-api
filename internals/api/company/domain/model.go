package companydomain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotFound     = errors.New("company not found")
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

func ToCompanyDomain(db CompanyDB) (Company, error) {
	// we can use type conversation because properties are the same
	return Company(db), nil
}

func ToCompaniesDomain(dbs []CompanyDB) ([]Company, error) {
	companies := make([]Company, len(dbs))
	// TODO: throw proper error here when conversation happens.
	for i, db := range dbs {
		company, err := ToCompanyDomain(db)
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
