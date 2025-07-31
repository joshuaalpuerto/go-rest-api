package infrarepositories

import (
	"context"
	"database/sql"
	"fmt"

	userdomain "github.com/joshuaalpuerto/go-rest-api/internal/api/user/domain"
	"github.com/joshuaalpuerto/go-rest-api/internal/infra/db"
)

type UserRepository struct {
	storer db.Postgres
}

func NewUserRepository(db db.Postgres) *UserRepository {
	return &UserRepository{
		storer: db,
	}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]userdomain.UserDB, error) {
	rows, err := r.storer.GetDB().QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	var users []userdomain.UserDB
	for rows.Next() {
		var user userdomain.UserDB
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return users, nil
}

func (r *UserRepository) FindOneByID(ctx context.Context, id string) (*userdomain.UserDB, error) {
	var user userdomain.UserDB
	err := r.storer.GetDB().QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, userdomain.ErrNotFound
		}
		return nil, fmt.Errorf("failed to query row: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user userdomain.NewUser) (*userdomain.UserDB, error) {
	var userDB userdomain.UserDB
	err := r.storer.GetDB().QueryRowContext(ctx, "INSERT INTO users (name, email, password, created_at, updated_at, created_by, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *", user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy).Scan(
		&userDB.ID,
		&userDB.Name,
		&userDB.Email,
		&userDB.Password,
		&userDB.CreatedAt,
		&userDB.UpdatedAt,
		&userDB.CreatedBy,
		&userDB.UpdatedBy,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create: %w", err)
	}

	return &userDB, nil
}

func (r *UserRepository) Update(ctx context.Context, user userdomain.User) (*userdomain.UserDB, error) {
	var userDB userdomain.UserDB
	err := r.storer.GetDB().QueryRowContext(ctx, "UPDATE users SET name = $1 WHERE id = $2 RETURNING *", user.Name, user.ID).Scan(
		&userDB.ID,
		&userDB.Name,
		&userDB.Email,
		&userDB.CreatedAt,
		&userDB.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update: %w", err)
	}

	return &userDB, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) (*userdomain.UserDB, error) {
	var user userdomain.UserDB
	err := r.storer.GetDB().QueryRowContext(ctx, "DELETE FROM users WHERE id = $1 RETURNING *", id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to delete: %w", err)
	}

	return &user, nil
}
