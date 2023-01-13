package migrate

import (
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	createUsersTable = `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			hobby VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);
	`
)

func MigrationTable(db *sqlx.DB) {
	_, err := db.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Student table migrated.")
}
