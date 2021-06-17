package tests

import (
	"bookstore/api/auth"
	"bookstore/api/book"
	"bookstore/api/user"
	"fmt"

	"github.com/go-pg/pg/v10"
)

// Generator generates database records for testing purposes
type Generator struct {
	Pg *pg.DB
}

func (g *Generator) generateUser() (*user.User, string) {
	user := &user.User{
		FirstName: "Tue",
		LastName:  "Tester",
		Email:     "quangtue@tester.com",
		Password:  "quangtue1234",
	}
	if _, err := g.Pg.Model(user).Insert(); err != nil {
		handleError(err)
	}

	tokenAuth := auth.TokenAuth()
	token := tokenAuth.EncodeToken(user)

	return user, token
}

func (g *Generator) generateBooks(u *user.User, n int) []*book.Book {
	books := make([]*book.Book, 0, n)

	for i := 1; i <= n; i++ {
		book := &book.Book{
			Id:       u.Id,
			Name:     fmt.Sprintf("Name %d", i),
			Price:    1.1,
			Rating:   5.0,
			Author:   fmt.Sprintf("Author %d", i),
			Category: fmt.Sprintf("Category %d", i),
		}
		if _, err := g.Pg.Model(book).Insert(); err != nil {
			handleError(err)
		}
		books = append(books, book)
	}

	return books
}

func (g *Generator) countBooks() int {
	count, err := g.Pg.Model((*book.Book)(nil)).Count()
	if err != nil {
		handleError(err)
	}
	return count
}

func (g *Generator) truncateBooks() {
	if _, err := g.Pg.Exec("TRUNCATE TABLE books"); err != nil {
		handleError(err)
	}
}
