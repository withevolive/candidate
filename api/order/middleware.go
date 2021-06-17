package order

import (
	"bookstore/lib/response"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

// NoteContextKey is a concrete type used as a key, the point is to avoid collisions between packages using context
type OrderContextKey struct{}

// NoteContext middleware depends on a datastore
type OrderContext struct {
	Ds Orders
}

// Handler injects a note into the request context
func (m *OrderContext) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		role := string(claims["role"].(string))
		if role != "admin" {
			render.Render(w, r, response.NotFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
