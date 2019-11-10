package service

import (
	"errors"
	"issue-tracker/model"
	"issue-tracker/repo"
	"log"
)

type boardService struct {
	BoardRepo repo.BoardRepository
}

func NewBoardService(boardRepo repo.BoardRepository) BoardService {
	return &boardService{boardRepo}
}

func (b *boardService) GetAllBoards() []model.Board {
	log.Println("Fetching all boards.")
	return b.BoardRepo.FindAll()
}

func (b *boardService) GetBoard(ID string) model.Board {
	log.Printf("Getting board with ID '%s'.\n", ID)
	return b.BoardRepo.FindByID(ID)
}

func (b *boardService) AddBoard(board model.Board) (model.Board, error) {
	// TODO: Generate UUID
	board.ID = "123"

	if err := b.BoardRepo.Add(board); err != nil {
		log.Println("Unable to create board, error: " + err.Error())
		return model.Board{}, errors.New("something unexpected happened while creating a new board")
	}
	return board, nil
}

func (b *boardService) RemoveBoard(ID string) error {
	log.Printf("Deleting board with ID '%s'.\n", ID)
	return b.BoardRepo.Delete(ID)
}
