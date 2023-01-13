package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/renaldyhidayatt/crud_blog/dto"
)

type usersRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *usersRepository {
	return &usersRepository{db: db}
}

func (r *usersRepository) GetAll(ctx context.Context) ([]dto.Users, error) {
	var user dto.Users
	var users []dto.Users

	rowQuery, err := r.db.QueryContext(ctx, "SELECT * FROM users ORDER BY id DESC")

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
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

func (r *usersRepository) GetID(ctx context.Context, id int) (dto.Users, error) {
	var user dto.Users

	result, err := r.db.QueryContext(ctx, "SELECT id, name, hobby FROM users WHERE id = ?", id)

	if err != nil {
		log.Fatal("Error Query User: " + err.Error())
		return user, err
	}

	for result.Next() {
		err := result.Scan(&user.ID, &user.Name, &user.Hobby)
		if err != nil {
			return user, err
		}
	}

	return user, nil

}

func (r *usersRepository) Insert(ctx context.Context, usr *dto.Users) (dto.Users, error) {

	var user dto.Users

	crt, err := r.db.PrepareContext(ctx, "INSERT INTO users (name, hobby) VALUES (?, ?)")

	if err != nil {
		return user, err
	}

	res, err := crt.ExecContext(ctx, usr.Name, usr.Hobby)

	if err != nil {
		return user, err
	}

	rowID, err := res.LastInsertId()

	if err != nil {
		return user, err
	}

	user.ID = int(rowID)

	result, err := r.GetID(ctx, user.ID)

	if err != nil {
		return user, err
	}

	return result, nil
}

func (r *usersRepository) Update(ctx context.Context, usr dto.Users) (dto.Users, error) {

	crt, err := r.db.PrepareContext(ctx, "UPDATE users set name=?,hobby=? WHERE id=?")

	var user dto.Users

	if err != nil {
		return user, err
	}

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	_, queryError := crt.Exec(user.Name, usr.Hobby, usr.ID)
	if queryError != nil {
		return user, err
	}

	res, err := r.GetID(ctx, user.ID)
	if err != nil {
		return user, err
	}

	return res, nil
}

func (r *usersRepository) Delete(ctx context.Context, id int64) error {
	crt, err := r.db.PrepareContext(ctx, "DELETE FROM users WHERE id = ?")

	if err != nil {
		return err
	}

	_, queryError := crt.ExecContext(ctx, id)

	if queryError != nil {
		return err
	}

	return nil

}
