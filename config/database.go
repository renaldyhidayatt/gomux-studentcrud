package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	username string = "postgres"
	password string = ""
	host     string = "localhost"
	port     int    = 5432
	database string = "crud_blog"
)

var (
	dsn = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, port, database)
)

func InitialDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dsn)

	checkErr(err)
	err = db.Ping()
	checkErr(err)

	return db, nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
