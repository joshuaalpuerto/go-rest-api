package infrarepositories

import (
	"context"
	"fmt"

	"github.com/joshuaalpuerto/go-rest-api/internal/infra/db"
	onboardingdomain "github.com/joshuaalpuerto/go-rest-api/internal/onboarding/domain"
)

type OnboardingRepository struct {
	storer db.Postgres
}

func NewOnboardingRepository(db db.Postgres) *OnboardingRepository {
	return &OnboardingRepository{
		storer: db,
	}
}

func (r *OnboardingRepository) Create(ctx context.Context, newUserCompany onboardingdomain.UserCompany) (*onboardingdomain.UserCompanyDB, error) {
	var userCompanyDB onboardingdomain.UserCompanyDB
	err := r.storer.GetDB().QueryRowContext(ctx, "INSERT INTO user_companies (company_id, user_id, created_by, updated_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *", newUserCompany.CompanyID, newUserCompany.UserID, newUserCompany.CreatedBy, newUserCompany.UpdatedBy, newUserCompany.CreatedAt, newUserCompany.UpdatedAt).Scan(
		&userCompanyDB.ID,
		&userCompanyDB.CompanyID,
		&userCompanyDB.UserID,
		&userCompanyDB.CreatedAt,
		&userCompanyDB.UpdatedAt,
		&userCompanyDB.CreatedBy,
		&userCompanyDB.UpdatedBy,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user company: %w", err)
	}

	return &userCompanyDB, nil
}
