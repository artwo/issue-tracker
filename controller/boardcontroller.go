package controller

import (
	"encoding/json"
	"issue-tracker/model"
	"issue-tracker/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (c *RestController) getAllBoards(w http.ResponseWriter, r *http.Request) {
	boards := c.BoardService.GetAllBoards()
	_ = c.JSON(w, http.StatusOK, boards)
}

func (c *RestController) getBoard(w http.ResponseWriter, r *http.Request) {
	boardID := chi.URLParam(r, "boardID")
	board := c.BoardService.GetBoard(boardID)
	if board.IsNull() {
		log.Printf("Unable to find board with ID '%s'.\n", boardID)
		c.error(w, http.StatusNotFound, "Board not found", nil)
		return
	}
	_ = c.JSON(w, http.StatusOK, board)
}

func (c *RestController) postBoard(w http.ResponseWriter, r *http.Request) {
	var board model.Board
	err := json.NewDecoder(r.Body).Decode(&board)
	if err != nil {
		log.Printf("Unable to parse postBoard requet body, error: %s\n", err.Error())
		c.error(w, http.StatusBadRequest, "Unable to read the request body because it is invalid, empty or malformed", nil)
		return
	}

	if errs := board.Validate(); len(errs) > 0 {
		log.Printf("There is something wrong with the request body, errors: %s", utils.ErrorsToString(errs))
		c.error(w, http.StatusBadRequest, "There is something wrong with the request body", errs)
		return
	}

	createdBoard, err := c.BoardService.AddBoard(board)
	if err != nil {
		c.error(w, http.StatusInternalServerError, "Unable to create board due to an unexpected server error", nil)
		return
	}
	_ = c.JSON(w, http.StatusCreated, createdBoard)
}

func (c *RestController) deleteBoard(w http.ResponseWriter, r *http.Request) {
	boardID := chi.URLParam(r, "boardID")
	if boardID == "" {
		log.Println("Received a delete ticket request with no boardID.")
		c.error(w, http.StatusBadRequest, "The path parameter 'boardID' is missing", nil)
		return
	}

	board := c.BoardService.GetBoard(boardID)
	if board.IsNull() {
		log.Printf("Unable to find ticket with ID '%s'.\n", boardID)
		c.error(w, http.StatusNotFound, "Ticket not found", nil)
		return
	}

	if err := c.BoardService.RemoveBoard(board.ID); err != nil {
		log.Printf("Unable to delete board with ID '%s'.\n", boardID)
		c.error(w, http.StatusInternalServerError, "Unable to delete board due to an unexpected server error", nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}
