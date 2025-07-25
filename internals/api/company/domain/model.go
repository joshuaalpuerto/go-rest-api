package companydomain

import (
	"time"

	"github.com/google/uuid"
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

func (bus *Company) ToDB() CompanyDB {
	return CompanyDB{
		ID:        bus.ID,
		Name:      bus.Name,
		CreatedAt: bus.CreatedAt,
		UpdatedAt: bus.UpdatedAt,
		CreatedBy: bus.CreatedBy,
		UpdatedBy: bus.UpdatedBy,
	}
}

func ToCompanyDomain(db CompanyDB) (Company, error) {
	c := Company{
		ID:        db.ID,
		Name:      db.Name,
		CreatedAt: db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
		CreatedBy: db.CreatedBy,
		UpdatedBy: db.UpdatedBy,
	}

	return c, nil
}

func ToCompaniesDomain(dbs []CompanyDB) ([]Company, error) {
	companies := make([]Company, len(dbs))
	for i, db := range dbs {
		company, err := ToCompanyDomain(db)
		if err != nil {
			return nil, err
		}
		companies[i] = company
	}
	return companies, nil
}
