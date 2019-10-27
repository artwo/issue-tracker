package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"issue-tracker/model"
	"log"
)

type MongoTicketRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoTicketRepository(client *mongo.Client) TicketRepository {
	collection := client.Database("issue_tracker").Collection("ticket")
	return &MongoTicketRepository{
		client,
		collection,
	}
}

func (r *MongoTicketRepository) FindByID(ID string) (model.Ticket, error) {
	var result model.Ticket
	err := r.collection.FindOne(context.TODO(), bson.D{{"id", ID}}).Decode(&result)
	if err != nil {
		log.Printf("Unable to retrieve ticket with ID: '%s', error: %s\n", ID, err)
		return model.Ticket{}, err
	}
	return result, nil
}

func (r *MongoTicketRepository) FindAllByStatus(status model.Status) ([]model.Ticket, error) {
	panic("implement me")
}

func (r *MongoTicketRepository) Add(ticket model.Ticket) error {
	panic("implement me")
}

func (r *MongoTicketRepository) Remove(ticket model.Ticket) error {
	panic("implement me")
}
