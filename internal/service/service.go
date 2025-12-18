package service

import (
	"context"
	"time"

	"github.com/faisal-990/age/db/sqlc/generated"
	"github.com/faisal-990/age/internal/repository"
)

// --- 1. DTOs (Data Transfer Objects) ---

type CreateUserRequest struct {
	Name string
	Dob  time.Time // Handler parses string -> time.Time before calling this
}

type UpdateUserRequest struct {
	ID   int32
	Name string
	Dob  time.Time
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`           // Output format: "YYYY-MM-DD"
	Age  int32  `json:"age,omitempty"` // omitempty hides this field if it is 0
}

// --- 2. Service Interface ---

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*UserResponse, error)
	DeleteUser(ctx context.Context, id int32) error
	GetUser(ctx context.Context, id int32) (*UserResponse, error)
	ListUsers(ctx context.Context) ([]*UserResponse, error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UserResponse, error)
}

// --- 3. Service Implementation ---

type UserServiceStruct struct {
	r repository.UserRepo
}

func New(repo repository.UserRepo) UserService {
	return &UserServiceStruct{
		r: repo,
	}
}

// --- 4. Methods ---

func (s *UserServiceStruct) CreateUser(ctx context.Context, req CreateUserRequest) (*UserResponse, error) {
	// Map Service Request -> Repo Params
	arg := generated.CreateUserParams{
		Name: req.Name,
		Dob:  req.Dob,
	}

	// Call Repo
	user, err := s.r.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	// Map Repo Result -> Service Response
	// Note: We leave Age as 0 so it is hidden by 'omitempty'
	return &UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"), // Format time.Time to String
	}, nil
}

func (s *UserServiceStruct) DeleteUser(ctx context.Context, id int32) error {
	return s.r.DeleteUser(ctx, id)
}

func (s *UserServiceStruct) GetUser(ctx context.Context, id int32) (*UserResponse, error) {
	user, err := s.r.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// Handle Not Found (nil user)
	if user == nil {
		return nil, nil
	}

	return &UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  calculateAge(user.Dob),
	}, nil
}

func (s *UserServiceStruct) ListUsers(ctx context.Context) ([]*UserResponse, error) {
	users, err := s.r.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	// Iterate and Map
	var response []*UserResponse
	for _, u := range users {
		response = append(response, &UserResponse{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob.Format("2006-01-02"),
			Age:  calculateAge(u.Dob),
		})
	}
	return response, nil
}

func (s *UserServiceStruct) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UserResponse, error) {
	// Map Request -> Repo Params
	arg := generated.UpdateUserParams{
		ID:   req.ID,
		Name: req.Name,
		Dob:  req.Dob,
	}

	user, err := s.r.UpdateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  calculateAge(user.Dob),
	}, nil
}

// --- 5. Helper Function ---

func calculateAge(dob time.Time) int32 {
	now := time.Now()
	age := now.Year() - dob.Year()

	// If the birthday hasn't happened yet this year, subtract 1
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return int32(age)
}

