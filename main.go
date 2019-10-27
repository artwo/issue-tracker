package main

import (
	"issue-tracker/config"
	"issue-tracker/model"
	"issue-tracker/repo"
	"issue-tracker/utils"
	"log"
)

func main() {
	mongoClient := config.NewMongoConnection()
	ticketrepo := repo.NewMongoTicketRepository(mongoClient.Client)

	ticket := ticketrepo.FindByID("1234")
	log.Println(utils.ToString(ticket))

	err := ticketrepo.Delete("1234")
	if err != nil {
		log.Printf("Something wrong happends: %s\n", err)
	}
	err = ticketrepo.Delete("first-id")
	if err != nil {
		log.Printf("Something wrong happends: %s\n", err)
	}

	newTicket := model.Ticket{
		ID:          "1234",
		Title:       "This is a test ticket",
		Description: "I'm a super Test!!",
		Status:      model.StatusNew,
	}
	err = ticketrepo.Add(newTicket)
	if err != nil {
		log.Printf("Something wrong happends: %s\n", err)
	}

	newTicket.Status = model.StatusReady
	err = ticketrepo.Update(newTicket)
	if err != nil {
		log.Printf("Something wrong happends: %s\n", err)
	}

	updateNonExistingTicket := model.Ticket{
		ID:          "567",
		Title:       "This is a test ticket",
		Description: "I'm a super Test!!",
		Status:      model.StatusNew,
	}
	err = ticketrepo.Update(updateNonExistingTicket)
	if err != nil {
		log.Printf("Something wrong happends: %s\n", err)
	}
	err = ticketrepo.Add(updateNonExistingTicket)
	if err != nil {
		log.Printf("Something wrong happends: %s\n", err)
	}

	tickets := ticketrepo.FindAll()
	log.Println(utils.ToString(tickets))

	tickets = ticketrepo.FindAllByStatus(model.StatusReady)
	log.Println(utils.ToString(tickets))

	err = mongoClient.CloseMongoConnection()
	if err != nil {
		log.Fatal(err)
	}
}
