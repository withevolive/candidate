package book

import (
	"github.com/go-pg/pg/v10"
)

type Books interface {
	GetAllBooks() ([]*Book, error)
	GetAllAvailableBooks() ([]*Book, error)
	Get(id int) (*Book, error)
	Create(book *Book) (*Book, error)
	Update(book *Book) (*Book, error)
	Delete(id int) error
}

type Datastore struct {
	Pg *pg.DB
}

func (ds *Datastore) GetAllBooks() ([]*Book, error) {
	var books []*Book
	err := ds.Pg.Model(&books).Select()
	return books, err
}

func (ds *Datastore) GetAllAvailableBooks() ([]*Book, error) {
	var books []*Book
	err := ds.Pg.Model(&books).Where("available = ?", Is).Select()
	return books, err
}

func (ds *Datastore) Get(id int) (*Book, error) {
	book := &Book{Id: id}
	err := ds.Pg.Model(book).WherePK().Select()
	return book, err
}

func (ds *Datastore) Create(book *Book) (*Book, error) {
	_, err := ds.Pg.Model(book).Insert()
	return book, err
}

func (ds *Datastore) Update(book *Book) (*Book, error) {
	_, err := ds.Pg.Model(book).WherePK().Update()
	return book, err
}

func (ds *Datastore) Delete(id int) error {
	book := &Book{Id: id}
	_, err := ds.Pg.Model(book).WherePK().Delete()
	return err
}
