package infrarepositories

import (
	"context"
	"database/sql"
	"fmt"

	companydomain "github.com/joshuaalpuerto/go-rest-api/internals/api/company/domain"
	"github.com/joshuaalpuerto/go-rest-api/internals/infra/db"
)

type CompanyRepository struct {
	storer *db.Postgres
}

func NewCompanyRepository(db *db.Postgres) *CompanyRepository {
	return &CompanyRepository{
		storer: db,
	}
}

func (r *CompanyRepository) FindAll(ctx context.Context) ([]companydomain.CompanyDB, error) {
	rows, err := r.storer.GetDB().QueryContext(ctx, "SELECT * FROM companies")
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	var companies []companydomain.CompanyDB
	for rows.Next() {
		var company companydomain.CompanyDB
		err := rows.Scan(
			&company.ID,
			&company.Name,
			&company.CreatedAt,
			&company.UpdatedAt,
			&company.CreatedBy,
			&company.UpdatedBy,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return companies, nil
}

func (r *CompanyRepository) FindOne(ctx context.Context, id string) (*companydomain.CompanyDB, error) {
	var company companydomain.CompanyDB
	err := r.storer.GetDB().QueryRowContext(ctx, "SELECT * FROM companies WHERE id = $1", id).Scan(
		&company.ID,
		&company.Name,
		&company.CreatedAt,
		&company.UpdatedAt,
		&company.CreatedBy,
		&company.UpdatedBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query row: %w", err)
	}

	return &company, nil
}

func (r *CompanyRepository) Create(ctx context.Context, company companydomain.Company) (*companydomain.CompanyDB, error) {
	var companyDB companydomain.CompanyDB
	err := r.storer.GetDB().QueryRowContext(ctx, "INSERT INTO companies (name) VALUES ($1) RETURNING *", company.Name).Scan(
		&companyDB.ID,
		&companyDB.Name,
		&companyDB.CreatedAt,
		&companyDB.UpdatedAt,
		&companyDB.CreatedBy,
		&companyDB.UpdatedBy,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create: %w", err)
	}

	return &companyDB, nil
}

func (r *CompanyRepository) Update(ctx context.Context, query string, args ...interface{}) (*companydomain.CompanyDB, error) {
	var company companydomain.CompanyDB
	err := r.storer.GetDB().QueryRowContext(ctx, query, args...).Scan(
		&company.ID,
		&company.Name,
		&company.CreatedAt,
		&company.UpdatedAt,
		&company.CreatedBy,
		&company.UpdatedBy,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update: %w", err)
	}

	return &company, nil
}

func (r *CompanyRepository) Delete(ctx context.Context, query string, args ...interface{}) (*companydomain.CompanyDB, error) {
	var company companydomain.CompanyDB
	err := r.storer.GetDB().QueryRowContext(ctx, query, args...).Scan(
		&company.ID,
		&company.Name,
		&company.CreatedAt,
		&company.UpdatedAt,
		&company.CreatedBy,
		&company.UpdatedBy,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete: %w", err)
	}

	return &company, nil
}
