package repo

import "issue-tracker/model"

type TicketRepository interface {
	FindByID(ID string) model.Ticket
	FindAll() []model.Ticket
	FindAllByStatus(status model.Status) []model.Ticket
	Add(ticket model.Ticket) error
	Update(ticket model.Ticket) error
	Delete(ID string) error
}
