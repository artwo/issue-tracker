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

type BoardRepository interface {
	FindByID(ID string) model.Board
	FindAll() []model.Board
	Add(board model.Board) error
	Update(board model.Board) error
	Delete(ID string) error
}
