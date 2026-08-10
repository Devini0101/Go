package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Conn() *sql.DB {

	connection := "user=postgres dbname=loja password=@Heloisa0404 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
