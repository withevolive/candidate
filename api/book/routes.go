package book

import (
	"bookstore/lib/middleware"

	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v10"
)

// Env is used to inject a datastore into request handlers
type Env struct {
	Ds Books
}

// Routes sets up the router
func Routes(pgDb *pg.DB) *chi.Mux {
	ds := &Datastore{Pg: pgDb}
	env := &Env{Ds: ds}

	router := chi.NewRouter()

	bookContext := &BookContext{Ds: ds}

	// Routes for admin
	router.Route("/", func(router chi.Router) {
		router.Use(bookContext.Handler)
		router.Get("/", env.getBooks)
		router.Post("/", env.createBook)
	})

	// Routes for customer
	router.Get("/available", env.getBooksAvailable)

	router.Route("/{id}", func(router chi.Router) {
		router.Use(middleware.ID)
		// Route view a detailed description of the book
		router.Get("/", env.getBook)
		router.Put("/", env.updateBook)
		router.Delete("/", env.deleteBook)
	})

	return router
}
