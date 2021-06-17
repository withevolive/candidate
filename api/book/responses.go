package book

import (
	"net/http"

	"github.com/go-chi/render"
)

type BookResponse struct {
	*Book
}

// Render takes care of pre-processing before a response is marshalled and sent across the wire
func (response *BookResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewBookResponse returns a new BookResponse
func NewBookResponse(book *Book) *BookResponse {
	response := &BookResponse{Book: book}
	return response
}

// NewBookListResponse returns a new BookListResponse
func NewBookListResponse(books []*Book) []render.Renderer {
	list := []render.Renderer{}
	for _, book := range books {
		list = append(list, NewBookResponse(book))
	}
	return list
}
