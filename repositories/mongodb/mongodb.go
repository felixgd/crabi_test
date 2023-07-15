package mongodb

import (
	"context"
	"crabi_test/domain"
	"crabi_test/utils/errors"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// other necessary imports
)

const (
	dbName         = "crabi-test"
	collectionName = "user"
)

type MongoClient struct {
	Client *mongo.Client
}

func NewMongoClient(connectionString string) (*MongoClient, error) {
	// Set up the MongoDB client options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	// Create a new MongoDB client
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	return &MongoClient{
		Client: client,
	}, nil
}

type MongoDBRepository interface {
	GetByEmail(email string) (*domain.User, error)
	Create(data domain.User) error
}

func (mc *MongoClient) Close() error {
	err := mc.Client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	log.Println("Disconnected from MongoDB!")

	return nil
}

func (c *MongoClient) GetByEmail(ctx context.Context, email string) (*domain.User, *errors.APIError) {
	var user domain.User

	collection := c.Client.Database(dbName).Collection(collectionName)

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(ctx, filter).Decode(&user)
	fmt.Println(err)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, &errors.APIError{
				Message: "Not found",
				Code:    http.StatusNotFound,
				Err:     fmt.Errorf("Not found"),
			}
		}
		return nil, &errors.APIError{
			Message: "Internal server error",
			Code:    http.StatusInternalServerError,
			Err:     fmt.Errorf("Internal server Error"),
		}
	}

	return &user, nil
}

func (c *MongoClient) Create(ctx context.Context, user *domain.User) (*domain.User, *errors.APIError) {
	collection := c.Client.Database(dbName).Collection(collectionName)
	createdUser, err := c.GetByEmail(ctx, user.Email)
	if err != nil {
		if err.Code != http.StatusNotFound {
			return nil, &errors.APIError{
				Message: "Internal Server Error",
				Code:    http.StatusInternalServerError,
				Err:     fmt.Errorf("Internal server Error"),
			}
		}
	}

	if createdUser != nil {
		return nil, &errors.APIError{
			Message: "User already created",
			Code:    http.StatusBadRequest,
			Err:     fmt.Errorf("Bad Request"),
		}
	}

	_, colErr := collection.InsertOne(ctx, user)
	if colErr != nil {
		return nil, &errors.APIError{
			Message: "Internal Server Error",
			Code:    http.StatusInternalServerError,
			Err:     fmt.Errorf("Internal server Error"),
		}
	}

	return user, nil
}
