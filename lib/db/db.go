package db

import (
	"bookstore/api/book"
	"bookstore/api/order"
	"bookstore/api/user"
	"context"
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	formattedQuery, _ := q.FormattedQuery()
	log.Printf("%s\n", formattedQuery)
	return nil
}

// Connection establishes a new database connection
func Connection(dbHost, dbPort, dbName, dbUser, dbPassword string, logQueries bool) *pg.DB {
	db := pg.Connect(&pg.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		Database: dbName,
		User:     dbUser,
		Password: dbPassword,
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Panicf("Error: %s\n", err.Error())
	}

	err = createSchema(db)
	if err != nil {
		log.Panicf("Error: %s\n", err.Error())
	}

	if logQueries {
		db.AddQueryHook(dbLogger{})
	}

	return db
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*user.User)(nil),
		(*book.Book)(nil),
		(*order.Order)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
