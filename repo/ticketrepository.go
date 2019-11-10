package repo

import (
	"context"
	"errors"
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

func (r *MongoTicketRepository) FindByID(ID string) model.Ticket {
	var result model.Ticket
	err := r.collection.FindOne(context.TODO(), bson.D{{"id", ID}}).Decode(&result)
	if err != nil {
		log.Printf("Unable to retrieve ticket with ID: '%s', error: %s\n", ID, err)
		return model.Ticket{}
	}
	return result
}

func (r *MongoTicketRepository) findMany(filter bson.D) []model.Ticket {
	var results []model.Ticket

	cur, err := r.collection.Find(context.TODO(), filter, nil)
	if err != nil {
		log.Printf("Unable to retrieve tickets, error: %s\n", err)
		return []model.Ticket{}
	}

	for cur.Next(context.TODO()) {
		var elem model.Ticket
		err := cur.Decode(&elem)
		if err != nil {
			log.Printf("Unable to decode ticket, error: %s\n", err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Printf("Something unexpected went wrong while retreaving tickets, error: %s\n", err)
	}
	if err := cur.Close(context.TODO()); err != nil {
		log.Printf("Unable to close cursor while retreaving tickets, error: %s\n", err)
	}

	return results
}

func (r *MongoTicketRepository) FindAll() []model.Ticket {
	return r.findMany(bson.D{{}})
}

func (r *MongoTicketRepository) FindAllByStatus(status model.Status) []model.Ticket {
	return r.findMany(bson.D{{"status", status}})
}

func (r *MongoTicketRepository) Add(ticket model.Ticket) error {
	_, err := r.collection.InsertOne(context.TODO(), ticket)
	return err
}

func (r *MongoTicketRepository) Update(ticket model.Ticket) error {
	filter := bson.D{{"id", ticket.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"title", ticket.Title},
			{"description", ticket.Description},
			{"status", ticket.Status},
		}},
	}

	updateResult, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if updateResult.MatchedCount == 0 || updateResult.ModifiedCount == 0 {
		return errors.New("Ticket '" + ticket.ID + "' cannot be updated because it has not been created yet")
	}
	return err
}

func (r *MongoTicketRepository) Delete(ID string) error {
	deleteResult, err := r.collection.DeleteOne(context.TODO(), bson.D{{"id", ID}}, nil)
	if deleteResult.DeletedCount == 0 {
		return errors.New("Ticket '" + ID + "' cannot be deleted because it does not exist")
	}
	return err
}
