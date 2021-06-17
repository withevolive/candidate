package order

import (
	"errors"
	"net/http"
)

type OrderRequest struct {
	*Order
}

func (order *OrderRequest) Bind(r *http.Request) error {
	if order.Order == nil {
		return errors.New("Missing required JSON attributes.")
	}
	return order.Order.Validate()
}
