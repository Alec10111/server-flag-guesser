# api-flag-guesser

Backend for Flag Guesser web application written in Go.
The API uses built in net/http module, httprouter.
The database is handled by the gorm package.
Webscraping makes use of colly package.

## TODO

-   Unit tests
-   JWT authentication
-   Country information on database
-   Save game scores in database

## Setup

### Docker

Run `docker compose up`

### Local

Run `go mod download` and then `go run .` (For some reason `go run main.go` does not work for me)
