package controller

import (
	"encoding/json"
	"issue-tracker/model"
	"issue-tracker/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (c *RestController) getAllTickets(w http.ResponseWriter, r *http.Request) {
	tickets := c.TicketService.GetAllTickets()
	_ = c.JSON(w, http.StatusOK, tickets)
}

func (c *RestController) getTicket(w http.ResponseWriter, r *http.Request) {
	ticketID := chi.URLParam(r, "ticketID")
	ticket := c.TicketService.GetTicket(ticketID)
	if (ticket == model.Ticket{}) {
		log.Printf("Unable to find ticket with ID '%s'.\n", ticketID)
		c.error(w, http.StatusNotFound, "Ticket not found", nil)
		return
	}
	_ = c.JSON(w, http.StatusOK, ticket)
}

func (c *RestController) postTicket(w http.ResponseWriter, r *http.Request) {
	var ticket model.Ticket
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		log.Printf("Unable to parse postTicket requet body, error: %s\n", err.Error())
		c.error(w, http.StatusBadRequest, "Unable to read the request body because it is invalid, empty or malformed", nil)
		return
	}

	if errs := ticket.Validate(); len(errs) > 0 {
		log.Printf("There is something wrong with the request body, errors: %s", utils.ErrorsToString(errs))
		c.error(w, http.StatusBadRequest, "There is something wrong with the request body", errs)
		return
	}

	createdTicket, err := c.TicketService.AddTicket(ticket)
	if err != nil {
		c.error(w, http.StatusInternalServerError, "Unable to create ticket due to an unexpected server error", nil)
		return
	}
	_ = c.JSON(w, http.StatusCreated, createdTicket)
}

func (c *RestController) deleteTicket(w http.ResponseWriter, r *http.Request) {
	ticketID := chi.URLParam(r, "ticketID")
	if ticketID == "" {
		log.Println("Received a delete ticket request with no ticketID.")
		c.error(w, http.StatusBadRequest, "The path parameter 'ticketID' is missing", nil)
		return
	}

	ticket := c.TicketService.GetTicket(ticketID)
	if (ticket == model.Ticket{}) {
		log.Printf("Unable to find ticket with ID '%s'.\n", ticketID)
		c.error(w, http.StatusNotFound, "Ticket not found", nil)
		return
	}

	if err := c.TicketService.RemoveTicket(ticket.ID); err != nil {
		log.Printf("Unable to delete ticket with ID '%s'.\n", ticketID)
		c.error(w, http.StatusInternalServerError, "Unable to delete ticket due to an unexpected server error", nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}
