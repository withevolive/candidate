package order

import (
	"bookstore/api/auth"
	"bookstore/api/user"
	"bookstore/lib/response"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func (env *Env) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := env.Ds.GetAllOrders()
	if err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
	if err := render.RenderList(w, r, NewOrderListResponse(orders)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}

func (env *Env) createOrder(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(auth.UserContextKey{}).(*user.User)
	data := &OrderRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.UnprocessableEntity(err))
		return
	}
	for k, v := range data.Books {
		fmt.Println(k, v)
	}
	order := data.Order
	order.UserId = user.Id
	order.User = user
	order, err := env.Ds.Create(order)
	if err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
	if err := render.Render(w, r, NewOrderResponse(order)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}
