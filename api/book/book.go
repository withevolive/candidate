package book

import (
	"context"
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	Id        int       `pg:",pk"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Rating    float32   `json:"rating"`
	Author    string    `json:"author"`
	Category  string    `json:"category"`
	Available Available `json:"available"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required),
		validation.Field(&b.Price, validation.Required),
		validation.Field(&b.Rating, validation.Required),
		validation.Field(&b.Author, validation.Required),
		validation.Field(&b.Category, validation.Required),
	)
}

func (b *Book) BeforeInsert(ctx context.Context) (context.Context, error) {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return ctx, nil
}

func (b *Book) BeforeUpdate(ctx context.Context) (context.Context, error) {
	b.UpdatedAt = time.Now()
	return ctx, nil
}

func (b *Book) IsAvailable(ctx context.Context) (context.Context, error) {
	switch b.Available {
	case Is, IsNot:
		return ctx, nil
	}
	return nil, errors.New("Invalid Availabe")
}
