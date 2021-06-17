package middleware

import (
	"bookstore/lib/request"
	"bookstore/lib/response"
	"context"
	"net/http"

	"github.com/go-chi/render"
)

// IDContextKey is a concrete type used as a key, the point is to avoid collisions between packages using context
type IDContextKey struct{}

// ID middleware injects id URL parameter into the request context
func ID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := request.ParamInt(r, "id")
		if err != nil {
			render.Render(w, r, response.BadRequest(err))
			return
		}

		ctx := context.WithValue(r.Context(), IDContextKey{}, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
