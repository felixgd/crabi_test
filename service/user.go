package service

// Service represents a service or dependency that the API depends on.
type Service interface {
	FetchUser() string
	CreateUser() string
	AuthUser() string
}

// MyService is an implementation of the Service interface.
type MyService struct{}

// FetchData is a method of MyService that fetches some data.
func (s *MyService) FetchUser() string {
	return "Hello, world!"
}

func (s *MyService) CreateUser() string {
	return "Hello, world!"
}

func (s *MyService) AuthUser() string {
	return "Hello, world!"
}
