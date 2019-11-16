package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"issue-tracker/config"
	"issue-tracker/controller"
	"issue-tracker/repo"
	"issue-tracker/service"
	"log"
	"net/http"
	"time"
)

func Routes(c *controller.RestController) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		controller.RequestTraceMiddleware,
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.RealIP,
		middleware.Recoverer,
	)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Route("/issue-tracker", func(r chi.Router) {
		r.Mount("/v1", c.Routes())
	})
	return router
}

func walkFunc(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	log.Printf("%s %s\n", method, route)
	return nil
}

func main() {
	log.Println("Starting Issue Tracker API.")
	port := ":8080"

	mongoClient := config.NewMongoConnection()
	defer mongoClient.CloseMongoConnection()

	ticketRepo := repo.NewMongoTicketRepository(mongoClient.Client)
	ticketService := service.NewTicketService(ticketRepo)
	boardRepo := repo.NewMongoBoardRepository(mongoClient.Client)
	boardService := service.NewBoardService(boardRepo)

	c := controller.NewRestController(ticketRepo, ticketService, boardRepo, boardService)
	router := Routes(c)

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Fatalf("Logging err: %s\n", err.Error())
	}

	log.Printf("API available at port '%s'.\n", port)
	log.Fatal(http.ListenAndServe(port, router))
	// TODO: create graceful shutdown
}
