package service

import "issue-tracker/model"

type TicketService interface {
	GetAllTickets() []model.Ticket
	GetTicket(ID string) model.Ticket
	AddTicket(ticket model.Ticket) (model.Ticket, error)
	RemoveTicket(ID string) error
}

type BoardService interface {
	GetAllBoards() []model.Board
	GetBoard(ID string) model.Board
	AddBoard(ticket model.Board) (model.Board, error)
	RemoveBoard(ID string) error
}
