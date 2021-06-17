package book

import (
	"bookstore/lib/middleware"
	"bookstore/lib/response"
	"net/http"

	"github.com/go-chi/render"
)

func (env *Env) getBooks(w http.ResponseWriter, r *http.Request) {
	books, err := env.Ds.GetAllBooks()
	if err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
	if err := render.RenderList(w, r, NewBookListResponse(books)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}

func (env *Env) getBooksAvailable(w http.ResponseWriter, r *http.Request) {
	books, err := env.Ds.GetAllAvailableBooks()
	if err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
	if err := render.RenderList(w, r, NewBookListResponse(books)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}

func (env *Env) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.IDContextKey{}).(int)
	book, err := env.Ds.Get(id)
	if err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
	if err := render.Render(w, r, NewBookResponse(book)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}

func (env *Env) createBook(w http.ResponseWriter, r *http.Request) {
	data := &BookRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.UnprocessableEntity(err))
		return
	}

	book := data.Book
	book, err := env.Ds.Create(book)
	if err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
	if err := render.Render(w, r, NewBookResponse(book)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}

func (env *Env) updateBook(w http.ResponseWriter, r *http.Request) {
	data := &BookRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.UnprocessableEntity(err))
		return
	}

	book := data.Book
	book, err := env.Ds.Update(book)
	if err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
	if err := render.Render(w, r, NewBookResponse(book)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}

func (env *Env) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.IDContextKey{}).(int)

	if err := env.Ds.Delete(id); err != nil {
		render.Render(w, r, response.NotFound)
		return
	}
	if err := render.Render(w, r, response.Success); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}
