# Example Go Language Application

This application is a guestbook, it will walk through a basic CRUD application, the logic on the server and how to work with golang templates.

## Setup

### Prerequisites

1. Ensure you know the login credentials for the MySQL instance you will be working with
2. Create a database that will be used to store entry information.
3. Modify `config/config.json` to match your environments credentials so the application can connect to MySQL.

## Run

To run, from the root of the project run the following:

`go run ./cmd/frontend/main.go`