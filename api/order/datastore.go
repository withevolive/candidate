package order

import (
	"github.com/go-pg/pg/v10"
)

type Orders interface {
	GetAllOrders() ([]*Order, error)
	Create(order *Order) (*Order, error)
}

type Datastore struct {
	Pg *pg.DB
}

func (ds *Datastore) GetAllOrders() ([]*Order, error) {
	var orders []*Order
	err := ds.Pg.Model(&orders).Column(
		"order.*").Relation("User.last_name").Select()
	return orders, err
}

func (ds *Datastore) Create(order *Order) (*Order, error) {
	_, err := ds.Pg.Model(order).Insert()
	return order, err
}
