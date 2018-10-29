package main

import (
	"log"
	"net/http"
	"os"

	"github.com/0110101001110011/blade2/lib"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/todo", todo.Routes())
	})

	return router
}

func main() {
	// Get argument 1 (argument 0 appears to be the gopath) to use as the port
	port := os.Args[1]

	// Set up router
	router := Routes()

	// TODO probably remove this
	// Print out a list of all available routes and their allowed methods
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	// Panic if there is an error
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	// Log before init: Port that will be used
	log.Printf("Spinning up webserver on port [%s]", port)

	// Run the server. A returned value indicates a fatal error or manual exit
	log.Fatal(http.ListenAndServe(":"+port, router))
}
