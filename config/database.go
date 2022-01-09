package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

const (
	username string = "root"
	password string = ""
	database string = "crud_blog"
)

var (
	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

func InitialDatabase() (*sql.DB, error) {
	db, err = sql.Open("mysql", dsn)

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
