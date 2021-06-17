package tests

import (
	"bookstore/api/app"
	"bookstore/api/book"
	"bookstore/api/user"
	"bookstore/lib/env"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var a app.App
var u *user.User
var token string
var generator *Generator

func TestUnauthorized(t *testing.T) {
	request, errRequest := http.NewRequest("GET", "/v1/book", nil)
	if errRequest != nil {
		t.Error(errRequest)
	}
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusUnauthorized, response.Code)
}

func TestNoResults(t *testing.T) {
	request, errRequest := http.NewRequest("GET", "/v1/book", nil)
	if errRequest != nil {
		t.Error(errRequest)
	}
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	assert.Equal(t, "[]", strings.TrimSpace(response.Body.String()))
}

func TestGetBooks(t *testing.T) {
	generator.generateBooks(u, 3)
	defer generator.truncateBooks()

	request, errRequest := http.NewRequest("GET", "/v1/book", nil)
	if errRequest != nil {
		t.Error(errRequest)
	}
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	var books []*book.Book
	if errJSON := json.Unmarshal([]byte(response.Body.String()), &books); errJSON != nil {
		t.Error(errJSON)
	}

	assert.Equal(t, 3, len(books), "Expected 6 books, got %v", len(books))
}

func TestCreateBook(t *testing.T) {
	defer generator.truncateBooks()

	body, errBody := json.Marshal(map[string]interface{}{
		"name":     "New name",
		"price":    5000.5,
		"rating":   6.0,
		"author":   "New Author",
		"category": "New Category",
	})
	if errBody != nil {
		t.Error(errBody)
	}

	request, errRequest := http.NewRequest("POST", "/v1/book", bytes.NewBuffer(body))
	if errRequest != nil {
		t.Error(errRequest)
	}
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	count := generator.countBooks()
	assert.Equal(t, 1, count, "Expected 1 book, got %v", count)
}

func TestUpdateNonExistentBook(t *testing.T) {
	request, errRequest := http.NewRequest("PUT", "/v1/book/999", nil)
	if errRequest != nil {
		t.Error(errRequest)
	}
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusNotFound, response.Code)
}

func TestUpdateBook(t *testing.T) {
	books := generator.generateBooks(u, 1)
	defer generator.truncateBooks()

	body, errBody := json.Marshal(map[string]interface{}{
		"name":     "New name",
		"price":    5000.5,
		"rating":   6.0,
		"author":   "New Author",
		"category": "New Category",
	})
	if errBody != nil {
		t.Error(errBody)
	}

	id := strconv.Itoa(books[0].Id)
	request, errRequest := http.NewRequest("PUT", "/v1/book/"+id, bytes.NewBuffer(body))
	if errRequest != nil {
		t.Error(errRequest)
	}
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)
}

func TestDeleteNonExistentBook(t *testing.T) {
	request, errRequest := http.NewRequest("DELETE", "/v1/book/999", nil)
	if errRequest != nil {
		t.Error(errRequest)
	}
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusNotFound, response.Code)
}

func TestDeleteBook(t *testing.T) {
	books := generator.generateBooks(u, 1)
	defer generator.truncateBooks()

	id := strconv.Itoa(books[0].Id)
	request, errRequest := http.NewRequest("DELETE", "/v1/book/"+id, nil)
	if errRequest != nil {
		t.Error(errRequest)
	}
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	count := generator.countBooks()
	assert.Equal(t, 0, count, "Expected 0 books, got %v", count)
}

func TestMain(m *testing.M) {
	env.LoadEnv("../.env")
	a = makeApp()
	generator = &Generator{Pg: a.Pg}
	u, token = generator.generateUser()
	code := m.Run()
	os.Exit(code)
}
