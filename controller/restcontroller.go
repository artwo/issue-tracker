package controller

import (
	"issue-tracker/model"
	"issue-tracker/service"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

type RestController struct {
	*render.Render
	TicketService service.TicketService
	BoardService  service.BoardService
}

func newJSONRender() *render.Render {
	return render.New(render.Options{
		//IndentJSON: true,
	})
}

func NewRestController(ticketService service.TicketService, boardService service.BoardService) *RestController {
	return &RestController{
		newJSONRender(),
		ticketService,
		boardService,
	}
}

func (c *RestController) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ticket", c.getAllTickets)
	router.Get("/ticket/{ticketID}", c.getTicket)
	router.Post("/ticket", c.postTicket)
	router.Delete("/ticket/{ticketID}", c.deleteTicket)

	router.Get("/board", c.getAllBoards)
	router.Get("/board/{boardID}", c.getBoard)
	router.Post("/board", c.postBoard)
	router.Delete("/board/{boardID}", c.deleteBoard)

	return router
}

func (c *RestController) error(w http.ResponseWriter, status int, message string, errors []error) {
	var errorMessages []string
	for _, err := range errors {
		errorMessages = append(errorMessages, err.Error())
	}
	_ = c.JSON(w, status, model.Error{status, message, errorMessages})
}
