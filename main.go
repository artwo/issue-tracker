package main

import (
	"issue-tracker/config"
	"log"
)

func main() {
	mongoClient := config.NewMongoConnection()

	err := mongoClient.CloseMongoConnection()
	if err != nil {
		log.Fatal(err)
	}
}
