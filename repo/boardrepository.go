package repo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"issue-tracker/model"
	"log"
)

type MongoBoardRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoBoardRepository(client *mongo.Client) BoardRepository {
	collection := client.Database("issue_tracker").Collection("board")
	return &MongoBoardRepository{
		client,
		collection,
	}
}

func (r *MongoBoardRepository) FindByID(ID string) model.Board {
	var result model.Board
	err := r.collection.FindOne(context.TODO(), bson.D{{"id", ID}}).Decode(&result)
	if err != nil {
		log.Printf("Unable to retrieve board with ID: '%s' from repo, error: %s\n", ID, err)
		return model.Board{}
	}
	return result
}

func (r *MongoBoardRepository) findMany(filter bson.D) []model.Board {
	var results []model.Board

	cur, err := r.collection.Find(context.TODO(), filter, nil)
	if err != nil {
		log.Printf("Unable to retrieve boards from repo, error: %s\n", err)
		return []model.Board{}
	}

	for cur.Next(context.TODO()) {
		var elem model.Board
		err := cur.Decode(&elem)
		if err != nil {
			log.Printf("Unable to decode board from repo, error: %s\n", err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Printf("Something unexpected went wrong while retreaving boards from repo, error: %s\n", err)
	}
	if err := cur.Close(context.TODO()); err != nil {
		log.Printf("Unable to close cursor while retreaving boards from repo, error: %s\n", err)
	}

	if len(results) == 0 {
		return make([]model.Board, 0)
	}

	return results
}

func (r *MongoBoardRepository) FindAll() []model.Board {
	return r.findMany(bson.D{{}})
}

func (r *MongoBoardRepository) Add(board model.Board) error {
	_, err := r.collection.InsertOne(context.TODO(), board)
	return err
}

func (r *MongoBoardRepository) Update(board model.Board) error {
	filter := bson.D{{"id", board.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"title", board.Name},
			{"description", board.Description},
		}},
	}

	updateResult, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if updateResult.MatchedCount == 0 || updateResult.ModifiedCount == 0 {
		return errors.New("Board '" + board.ID + "' cannot be updated from repo because it has not been created yet")
	}
	return err
}

func (r *MongoBoardRepository) Delete(ID string) error {
	deleteResult, err := r.collection.DeleteOne(context.TODO(), bson.D{{"id", ID}}, nil)
	if deleteResult.DeletedCount == 0 {
		return errors.New("Board '" + ID + "' cannot be deleted from repo because it does not exist")
	}
	return err
}
