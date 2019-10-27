package main

import (
	"issue-tracker/config"
	"issue-tracker/repo"
	"issue-tracker/utils"
	"log"
)

func main() {
	mongoClient := config.NewMongoConnection()

	ticketrepo := repo.NewMongoTicketRepository(mongoClient.Client)
	ticket, err := ticketrepo.FindByID("first-id")
	if err != nil {
		log.Fatal("Something went wrong")
	}
	log.Println(utils.ToString(ticket))

	err = mongoClient.CloseMongoConnection()
	if err != nil {
		log.Fatal(err)
	}
}
