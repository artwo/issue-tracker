package repo

import "issue-tracker/model"

type TicketRepository interface {
	FindByID(ID string) (model.Ticket, error)
	FindAllByStatus(status model.Status) ([]model.Ticket, error)
	Add(ticket model.Ticket) error
	Remove(ticket model.Ticket) error
}
