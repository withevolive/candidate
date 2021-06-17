package tests

import (
	"bookstore/api/app"
	"os"
)

func makeApp() app.App {
	var a app.App
	a.Initialize(
		os.Getenv("POSTGRES_DB_HOST"),
		os.Getenv("POSTGRES_DB_PORT"),
		os.Getenv("POSTGRES_TEST_DB_NAME"),
		os.Getenv("POSTGRES_DB_USER"),
		os.Getenv("POSTGRES_DB_PASSWORD"),
		false,
		false)

	return a
}
