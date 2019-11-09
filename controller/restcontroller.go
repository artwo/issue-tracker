package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/unrolled/render"
	"issue-tracker/model"
	"issue-tracker/repo"
	"issue-tracker/service"
	"issue-tracker/utils"
	"log"
	"net/http"
)

type RestController struct {
	*render.Render
	TicketRepo    repo.TicketRepository
	TicketService service.TicketService
}

func newJsonRender() *render.Render {
	return render.New(render.Options{
		//IndentJSON: true,
	})
}

func NewRestController(ticketRepo repo.TicketRepository, ticketService service.TicketService) *RestController {
	return &RestController{
		newJsonRender(),
		ticketRepo,
		ticketService,
	}
}

func (c *RestController) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/ticket", c.GetAllTickets)
	router.Get("/ticket/{ticketID}", c.GetTicket)
	router.Post("/ticket", c.PostTicket)
	router.Delete("/ticket/{ticketID}", c.DeleteTicket)
	return router
}

func (c *RestController) GetAllTickets(w http.ResponseWriter, r *http.Request) {
	tickets := c.TicketService.GetAllTickets()
	_ = c.JSON(w, http.StatusOK, tickets)
}

func (c *RestController) GetTicket(w http.ResponseWriter, r *http.Request) {
	ticketID := chi.URLParam(r, "ticketID")
	if ticketID == "" { // This should not happen due to the GetAll endpoint
		log.Println("Received a get ticket request with no ticketID.")
		c.Error(w, http.StatusBadRequest, "The path parameter 'ticketID' is missing", nil)
		return
	}

	ticket := c.TicketService.GetTicket(ticketID)
	if (ticket == model.Ticket{}) {
		log.Printf("Unable to find ticket with ID '%s'.\n", ticketID)
		c.Error(w, http.StatusNotFound, "Ticket not found", nil)
		return
	}
	_ = c.JSON(w, http.StatusOK, ticket)
}

func (c *RestController) PostTicket(w http.ResponseWriter, r *http.Request) {
	var ticket model.Ticket
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		log.Printf("Unable to parse PostTicket requet body, error: %s\n", err.Error())
		c.Error(w, http.StatusBadRequest, "Unable to read the request body because it is invalid, empty or malformed", nil)
		return
	}

	if errs := ticket.Validate(); len(errs) > 0 {
		log.Printf("There is something wrong with the request body, errors: %s", utils.ErrorsToString(errs))
		c.Error(w, http.StatusBadRequest, "There is something wrong with the request body", errs)
		return
	}

	createdTicket, err := c.TicketService.AddTicket(ticket)
	if err != nil {
		c.Error(w, http.StatusInternalServerError, "Unable to create ticket due to an unexpected server error", nil)
		return
	}
	_ = c.JSON(w, http.StatusCreated, createdTicket)
}

func (c *RestController) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	ticketID := chi.URLParam(r, "ticketID")
	if ticketID == "" {
		log.Println("Received a delete ticket request with no ticketID.")
		c.Error(w, http.StatusBadRequest, "The path parameter 'ticketID' is missing", nil)
		return
	}

	ticket := c.TicketService.GetTicket(ticketID)
	if (ticket == model.Ticket{}) {
		log.Printf("Unable to find ticket with ID '%s'.\n", ticketID)
		c.Error(w, http.StatusNotFound, "Ticket not found", nil)
		return
	}

	if err := c.TicketService.RemoveTicket(ticket.ID); err != nil {
		log.Printf("Unable to delete ticket with ID '%s'.\n", ticketID)
		c.Error(w, http.StatusInternalServerError, "Unable to delete ticket due to an unexpected server error", nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *RestController) Error(w http.ResponseWriter, status int, message string, errors []error) {
	var errorMessages []string
	for _, err := range errors {
		errorMessages = append(errorMessages, err.Error())
	}
	_ = c.JSON(w, status, model.Error{status, message, errorMessages})
}
