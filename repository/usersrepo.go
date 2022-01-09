package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/renaldyhidayatt/crud_blog/config"
	"github.com/renaldyhidayatt/crud_blog/dto"
)

const (
	table          = "users"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]dto.Users, error) {
	var users []dto.Users

	db, err := config.InitialDatabase()

	if err != nil {
		log.Fatal("Tidak bisa connect mysql")
	}
	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var user dto.Users

		if err = rowQuery.Scan(
			&user.ID,
			&user.Name,
			&user.Hobby,
		); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func Insert(ctx context.Context, user dto.Users) error {
	db, err := config.InitialDatabase()

	if err != nil {
		log.Fatal("Tidak Connect ke database ", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (name, hobby, created_at, updated_at) values('%s','%s','%v','%v')",
		table,
		user.Name,
		user.Hobby,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime),
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	return nil
}

func Update(ctx context.Context, user dto.Users) error {
	db, err := config.InitialDatabase()

	if err != nil {
		log.Fatal("Tidak Connect ke database ", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name='%s',hobby='%s' WHERE id='%d'", table, user.Name, user.Hobby, user.ID)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	return nil
}

func Delete(ctx context.Context, user dto.Users) error {
	db, err := config.InitialDatabase()
	if err != nil {
		log.Fatal("Tidak Connect ke database ", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, user.ID)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	check, err := s.RowsAffected()

	if check == 0 {

		return errors.New("Not found your id")
	}

	return nil

}
