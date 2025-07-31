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
type UserCompany struct {
	ID        uuid.UUID
	CompanyID uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}

type UserCompanyDB struct {
	ID        uuid.UUID `db:"id"`
	CompanyID uuid.UUID `db:"company_id"`
	UserID    uuid.UUID `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedBy string    `db:"created_by"`
	UpdatedBy string    `db:"updated_by"`
}

func (c UserCompanyDB) ToDomain() (UserCompany, error) {
	// we can use type conversation because properties are the same
	return UserCompany(c), nil
}

func ToUserCompanyEntities(dbs []UserCompanyDB) ([]UserCompany, error) {
	companies := make([]UserCompany, len(dbs))
	for i, db := range dbs {
		company, err := db.ToDomain()
		if err != nil {
			return nil, err
		}
		companies[i] = company
	}
	return companies, nil
}

type NewUserCompany struct {
	CompanyName  string
	UserName     string
	UserEmail    string
	UserPassword string
}
