package controller

import (
	"github.com/go-chi/chi"
	"github.com/unrolled/render"
	"issue-tracker/repo"
	"issue-tracker/service"
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
	return router
}

func (c *RestController) GetAllTickets(w http.ResponseWriter, r *http.Request) {
	tickets := c.TicketService.GetAllTickets()
	_ = c.JSON(w, http.StatusOK, tickets)
}
