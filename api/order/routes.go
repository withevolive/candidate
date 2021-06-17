package order

import (
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v10"
)

// Env is used to inject a datastore into request handlers
type Env struct {
	Ds Orders
}

// Routes sets up the router
func Routes(pgDb *pg.DB) *chi.Mux {
	ds := &Datastore{Pg: pgDb}
	env := &Env{Ds: ds}

	router := chi.NewRouter()

	orderContext := &OrderContext{Ds: ds}

	// Routes for admin
	router.Route("/", func(router chi.Router) {
		router.Use(orderContext.Handler)
		router.Get("/", env.GetAllOrders)
	})

	// Routes for customer
	router.Post("/", env.createOrder)

	return router
}
