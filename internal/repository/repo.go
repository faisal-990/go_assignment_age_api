package repository

import (
	"context"
	"database/sql"

	"github.com/faisal-990/age/db/sqlc/generated"
)

// UserRepo defines the interface for database operations.
type UserRepo interface {
	CreateUser(ctx context.Context, user generated.CreateUserParams) (generated.User, error)
	DeleteUser(ctx context.Context, id int32) error
	GetUser(ctx context.Context, id int32) (*generated.User, error)
	ListUsers(ctx context.Context) ([]generated.User, error)
	UpdateUser(ctx context.Context, arg generated.UpdateUserParams) (generated.User, error)
}

// SqlRepo implements UserRepo.
type SqlRepo struct {
	db      *sql.DB
	queries *generated.Queries
}

// New creates a new repository.
func New(db *sql.DB) UserRepo {
	return &SqlRepo{
		db:      db,
		queries: generated.New(db),
	}
}

// --- Implementation ---

func (s *SqlRepo) CreateUser(ctx context.Context, user generated.CreateUserParams) (generated.User, error) {
	// Directly call generated method
	return s.queries.CreateUser(ctx, user)
}

func (s *SqlRepo) DeleteUser(ctx context.Context, id int32) error {
	// Directly call generated method
	return s.queries.DeleteUser(ctx, id)
}

func (s *SqlRepo) GetUser(ctx context.Context, id int32) (*generated.User, error) {
	user, err := s.queries.GetUser(ctx, id)
	if err != nil {
		// If error is sql.ErrNoRows, this returns (nil, sql.ErrNoRows)
		return nil, err
	}
	// Return the address of the user struct to match the interface signature (*generated.User)
	return &user, nil
}

func (s *SqlRepo) ListUsers(ctx context.Context) ([]generated.User, error) {
	return s.queries.ListUsers(ctx)
}

func (s *SqlRepo) UpdateUser(ctx context.Context, arg generated.UpdateUserParams) (generated.User, error) {
	return s.queries.UpdateUser(ctx, arg)
}

