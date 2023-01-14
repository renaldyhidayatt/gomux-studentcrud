package test

import (
	"testing"

	"github.com/renaldyhidayatt/crud_blog/dto"
	"github.com/renaldyhidayatt/crud_blog/repository"
)

func TestGetAll(t *testing.T) {
	repository := repository.NewUserRepository(ConnTest, Context)

	users, err := repository.GetAll()

	if err != nil {
		t.Errorf("error was not expected while getting all users: %s", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, but got %d", len(users))
	}

}

func TestGetID(t *testing.T) {
	repository := repository.NewUserRepository(ConnTest, Context)

	user, err := repository.GetID(1)

	// validation
	if err != nil {
		t.Errorf("error was not expected while getting user by ID: %s", err)
	}
	if user.ID != 1 {
		t.Errorf("expected user ID 1, but got %d", user.ID)
	}
	if user.Name != "John Doe" {
		t.Errorf("expected user name 'John Doe', but got '%s'", user.Name)
	}
	if user.Hobby != "Running" {
		t.Errorf("expected user hobby 'Running', but got '%s'", user.Hobby)
	}
}

func TestInsert(t *testing.T) {
	repository := repository.NewUserRepository(ConnTest, Context)

	newUser := dto.Users{Name: "Jane Doe", Hobby: "Swimming"}
	insertedUser, err := repository.Insert(&newUser)

	// validation
	if err != nil {
		t.Errorf("error was not expected while inserting a new user: %s", err)
	}
	if insertedUser.Name != newUser.Name {
		t.Errorf("expected inserted user name '%s', but got '%s'", newUser.Name, insertedUser.Name)
	}
	if insertedUser.Hobby != newUser.Hobby {
		t.Errorf("expected inserted user hobby '%s', but got '%s'", newUser.Hobby, insertedUser.Hobby)
	}
}

func TestUpdate(t *testing.T) {
	repository := repository.NewUserRepository(ConnTest, Context)

	updatedUser := dto.Users{ID: 1, Name: "John Smith", Hobby: "Hiking"}
	updated, err := repository.Update(updatedUser)

	// validation
	if err != nil {
		t.Errorf("error was not expected while updating a user: %s", err)
	}
	if updated.Name != updatedUser.Name {
		t.Errorf("expected updated user name '%s', but got '%s'", updatedUser.Name, updated.Name)
	}
	if updated.Hobby != updatedUser.Hobby {
		t.Errorf("expected updated user hobby '%s', but got '%s'", updatedUser.Hobby, updated.Hobby)
	}
}

func TestDelete(t *testing.T) {
	repository := repository.NewUserRepository(ConnTest, Context)
	err := repository.Delete(1)

	// validation
	if err != nil {
		t.Errorf("error was not expected while deleting a user: %s", err)
	}
	deletedUser, err := repository.GetID(1)

	if err == nil {
		t.Errorf("user with ID 1 should have been deleted, but it still exists")
	}

	if deletedUser == (dto.Users{}) {
		t.Errorf("user with ID 1 should have been deleted, but it still exists")
	}
}
