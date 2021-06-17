# bookstore

## Requirements

- As a customer, I want to be able to view a list of books available for sale.
- As a customer, I want to be able to select the book I want, to view a detailed description of the book.
- As a customer, I want to be able to place an order for the book and provide delivery instructions. Note that this requirement does not require completing - the payment of the book.
- As a shop admin, I want to view the list of the books I currently sell in the store.
- As a shop admin, I want to see the list of the orders placed by the customers so that I can begin to fulfil the orders.

## Features

- modular code organization
- examples of standard CRUD operations
- environment dependent configuration (`.env` file)
- authentication using JSON Web Tokens
- request validation
- query and request logs
- route/struct binding using middleware
- integration tests
- Docker support


## Starting the server with docker

docker-compose up --build

## Running tests

`go test ./...`

