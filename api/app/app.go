package app

import (
	"bookstore/api/auth"
	"bookstore/api/book"
	"bookstore/api/order"
	"bookstore/api/user"
	"bookstore/lib/db"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-pg/pg/v10"
)

// App represents the application
type App struct {
	Router *chi.Mux
	Pg     *pg.DB
}

// Initialize sets up the database connection and router
func (a *App) Initialize(dbHost, dbPort, dbName, dbUser, dpPassword string, logRequests, logQueries bool) {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.SetHeader("Content-Type", "application/json"),
		middleware.Compress(5),
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	if logRequests {
		router.Use(middleware.Logger)
	}

	a.Pg = db.Connection(dbHost, dbPort, dbName, dbUser, dpPassword, logQueries)

	// Public routes
	router.Group(func(router chi.Router) {
		router.Mount("/auth", auth.Routes(a.Pg))
	})

	// Protected routes
	router.Route("/v1", func(router chi.Router) {
		tokenAuth := auth.TokenAuth()
		ds := &user.Datastore{Pg: a.Pg}
		userContext := &auth.UserContext{Ds: ds}

		router.Group(func(router chi.Router) {
			router.Use(tokenAuth.Verifier())
			router.Use(tokenAuth.Authenticator())
			router.Use(userContext.Handler)
			router.Mount("/book", book.Routes(a.Pg))
			router.Mount("/order", order.Routes(a.Pg))
		})
	})

	a.Router = router
}

// Serve serves the app on the specified port
func (a *App) Serve(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}
