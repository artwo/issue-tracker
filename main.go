package main

import (
	"context"
	"issue-tracker/config"
	"log"
)

func main() {
	mongoClient, err := config.NewMongoConnection()
	if err != nil {
		panic(err)
	}

	err = mongoClient.Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Ping to mongoDB successful")

	err = mongoClient.CloseMongoConnection()
	if err != nil {
		log.Fatal(err)
	}
}
