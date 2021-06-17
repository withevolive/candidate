package book

import (
	"errors"
	"net/http"
)

type BookRequest struct {
	*Book
}

func (book *BookRequest) Bind(r *http.Request) error {
	if book.Book == nil {
		return errors.New("Missing required JSON attributes.")
	}
	return book.Book.Validate()
}
