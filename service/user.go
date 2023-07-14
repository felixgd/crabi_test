package service

import (
	"context"
	"crabi_test/domain"
	"crabi_test/repositories/mongodb"
	"crabi_test/repositories/pld"
	"crabi_test/utils/constants"
	"crabi_test/utils/cypher"
	"crabi_test/utils/errors"
	"crabi_test/utils/jwt"
	"log"
	"net/http"
)

// Service represents a service or dependency that the API depends on.
type Service interface {
	FetchUser(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	AuthUser(ctx context.Context, user *domain.User) (map[string]string, error)
}

// UserService is an implementation of the Service interface.
type UserService struct {
	pldRepository pld.PLD
	mongoClient   *mongodb.MongoClient
}

// NewUserService creates a new instance of the Handler.
func NewUserService(pldRepository pld.PLD, mongoClient *mongodb.MongoClient) *UserService {
	return &UserService{
		pldRepository: pldRepository,
		mongoClient:   mongoClient,
	}
}

// FetchData is a method of UserService that fetches some data.
func (s *UserService) FetchUser(ctx context.Context, email string) (*domain.User, error) {
	user, err := s.mongoClient.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	pldPayload := domain.PLDPayload{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	pld, err := s.pldRepository.GetUserInPLD(&pldPayload)
	if err != nil {
		return nil, err
	}

	if pld.IsBlacklisted {
		return nil, errors.APIError{
			Code:    http.StatusBadRequest,
			Message: "Error while validating data",
		}
	}

	encryptedPassword, err := cypher.Encrypt(user.Password, constants.AES_KEY)
	if err != nil {
		return nil, errors.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	user.Password = encryptedPassword

	user, err = s.mongoClient.Create(ctx, user)
	if err != nil {
		log.Println("here3")
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func (s *UserService) AuthUser(ctx context.Context, user *domain.User) (map[string]string, error) {
	DBUser, err := s.mongoClient.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if user.Password != DBUser.Password {
		return nil, errors.APIError{
			Code:    http.StatusBadGateway,
			Message: "Authentication failed",
		}
	}

	token, err := jwt.GenerateToken(user.Email, []byte(constants.SECRET_KEY))
	if err != nil {
		return nil, err
	}

	response := make(map[string]string, 1)
	response["auth_token"] = token

	return response, nil
}
