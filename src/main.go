package main

import (
	"database/sql"
	"github.com/k13w/go_sqs_consumer/src/core/domain/entities"
	_ "github.com/lib/pq"
)

func dbConnect(dbUrl string) (*sql.DB, error) {
	//const dbUrl = "postgres://postgres:postgres@localhost:5300/postgres?sslmode=disable"

	open, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	err = open.Ping()
	if err != nil {
		return nil, err
	}

	return open, nil
}

//type app struct {
//
//}

//func handlerRequest (w http.ResponseWriter, r *http.Request) {
//	 fmt.Fprintf(w, "resposta 200")
//}

func main() {
	dbPollConnections, err := dbConnect("postgres://postgres:postgres@localhost:5300/postgres?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer func(dbPollConnections *sql.DB) {
		err := dbPollConnections.Close()
		if err != nil {

		}
	}(dbPollConnections)

	user := entities.UserEntity{
		DbConnect: dbPollConnections,
	}

	err = user.CreateUser("kiew", 1)
	if err != nil {
		panic(err.Error())
	}
}
