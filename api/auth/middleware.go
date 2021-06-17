package auth

import (
	"bookstore/api/user"
	"bookstore/lib/response"
	"context"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

// UserContextKey is a concrete type used as a key, the point is to avoid collisions between packages using context
type UserContextKey struct{}

// UserContext middleware depends on a datastore
type UserContext struct {
	Ds user.Users
}

// Handler injects authenticated user into the request context
func (m *UserContext) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		id := int(claims["id"].(float64))

		user, err := m.Ds.Get(id)
		if err != nil {
			render.Render(w, r, response.InternalServerError(err))
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey{}, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
