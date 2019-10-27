package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoClient struct {
	clientOptions *options.ClientOptions
	Client        *mongo.Client
}

func NewMongoConnection() (*MongoClient, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Unable to connect to MongoDB: %s\n", err.Error())
		return nil, err
	}
	log.Printf("New connection to MongoDB successful.")
	return &MongoClient{
		clientOptions,
		client,
	}, nil
}

func (m *MongoClient) CloseMongoConnection() error {
	err := m.Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("Unable to close connection to MongoDB: %s\n", err.Error())
		return err
	}
	log.Println("Disconnected from MongoDB successfully.")
	return nil
}
