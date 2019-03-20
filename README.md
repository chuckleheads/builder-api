# Builder API

A port of the Habitat Builder API to Go.

## Running
`go run cmd/builder-api/main.go`

## Routes that work

* POST /depot/origins
* GET /depot/origin/{origin}
* PUT /depot/origin/{origin}

## Updating the db models

The models are generated using [sqlboiler](https://github.com/volatiletech/sqlboiler). if you make db changes they'll need to be regenerated. To do so just run `go generate ./...` from the root of the repo.

#### Note
You'll need sqlboiler and sqlboiler-psql in your gopath. Instructions can be found [here](https://github.com/volatiletech/sqlboiler#getting-started)
