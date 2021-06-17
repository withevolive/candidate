package order

import (
	"net/http"

	"github.com/go-chi/render"
)

type OrderResponse struct {
	*Order
}

// Render takes care of pre-processing before a response is marshalled and sent across the wire
func (response *OrderResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewOrderResponse returns a new OrderResponse
func NewOrderResponse(order *Order) *OrderResponse {
	response := &OrderResponse{Order: order}
	return response
}

// NewOrderListResponse returns a new OrderListResponse
func NewOrderListResponse(orders []*Order) []render.Renderer {
	list := []render.Renderer{}
	for _, order := range orders {
		list = append(list, NewOrderResponse(order))
	}
	return list
}
