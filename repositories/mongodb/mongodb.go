package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	// other necessary imports
)

type MongoDBRepo struct {
	collection *mongo.Collection
}

type MongoDBRepository interface {
	GetByID(id string) (interface{}, error)
	Create(data interface{}) error
	Update(id string, data interface{}) error
	Delete(id string) error
}

func NewMongoDBRepo(client *mongo.Client, dbName, collectionName string) *MongoDBRepo {
	collection := client.Database(dbName).Collection(collectionName)
	return &MongoDBRepo{collection: collection}
}

// Implement the methods of the MongoDBRepository interface
// ...
