package service

import (
	"errors"
	"issue-tracker/model"
	"issue-tracker/repo"
	"log"
)

type ticketService struct {
	TicketRepo repo.TicketRepository
}

func NewTicketService(ticketRepo repo.TicketRepository) TicketService {
	return &ticketService{ticketRepo}
}

func (s *ticketService) GetAllTickets() []model.Ticket {
	log.Println("Fetching all tickets.")
	return s.TicketRepo.FindAll()
}

func (s *ticketService) GetTicket(ID string) model.Ticket {
	log.Printf("Getting ticket with ID '%s'.\n", ID)
	return s.TicketRepo.FindByID(ID)
}

func (s *ticketService) AddTicket(ticket model.Ticket) error {
	// TODO: Generate UUID
	ticket.ID = "123"

	if err := s.TicketRepo.Add(ticket); err != nil {
		log.Println("Unable to create ticket, error: " + err.Error())
		return errors.New("something unexpected happened while creating a new ticket")
	}
	return nil
}

func (s *ticketService) RemoveTicket(ID string) error {
	log.Printf("Deleting ticket with ID '%s'.\n", ID)
	return s.TicketRepo.Delete(ID)
}
