package service

import "crabi_test/repositories/pld"

// Service represents a service or dependency that the API depends on.
type Service interface {
	FetchUser() string
	CreateUser() string
	AuthUser() string
}

// UserService is an implementation of the Service interface.
type UserService struct {
	repositories pld.PLD
}

// NewUserService creates a new instance of the Handler.
func NewUserService(rp pld.PLD) *UserService {
	return &UserService{
		repositories: rp,
	}
}

// FetchData is a method of UserService that fetches some data.
func (s *UserService) FetchUser() string {
	return "Hello, world!"
}

func (s *UserService) CreateUser() string {
	return "Hello, world!"
}

func (s *UserService) AuthUser() string {
	return "Hello, world!"
}
