package test

import (
	"context"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/renaldyhidayatt/crud_blog/config"
)

var ConnTest *sqlx.DB
var Context = context.Background()

func TestMain(m *testing.M) {
	db, err := config.InitialDatabase()

	ConnTest = db

	if err != nil {
		log.Fatal(err.Error())
	}
}
