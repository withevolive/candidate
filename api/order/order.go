package order

import (
	"bookstore/api/book"
	"bookstore/api/user"
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Order struct {
	Id        int          `json:"id" pg:",pk"`
	Price     float64      `json:"price,omitempty"`
	Address   string       `json:"address"`
	Phone     string       `json:"phone"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	UserId    int          `json:"user_id"`
	User      *user.User   `pg:"rel:has-one"`
	Books     []*book.Book `json:"books" pg:"rel:has-many"`
}

func (o Order) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Address, validation.Required),
		validation.Field(&o.Phone, validation.Required),
		validation.Field(&o.Books, validation.Required),
	)
}

func (o *Order) BeforeInsert(ctx context.Context) (context.Context, error) {
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	return ctx, nil
}

func (o *Order) BeforeUpdate(ctx context.Context) (context.Context, error) {
	o.UpdatedAt = time.Now()
	return ctx, nil
}
