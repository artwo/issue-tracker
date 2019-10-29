package service

import "issue-tracker/model"

type TicketService interface {
	GetAllTickets() []model.Ticket
	GetTicket(ID string) model.Ticket
	AddTicket(ticket model.Ticket) error
	RemoveTicket(ID string) error
}
