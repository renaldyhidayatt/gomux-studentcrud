package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/renaldyhidayatt/crud_blog/dto"
	"github.com/renaldyhidayatt/crud_blog/handler"
	"github.com/renaldyhidayatt/crud_blog/repository"
	"github.com/renaldyhidayatt/crud_blog/services"
)

func TestGetAllHandler(t *testing.T) {
	repository := repository.NewUserRepository(ConnTest, Context)
	service := services.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	req, err := http.NewRequest("GET", "/users", nil)

	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()

	handler.GetAll(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var users []dto.Users
	err = json.Unmarshal(rr.Body.Bytes(), &users)
	if err != nil {
		t.Errorf("Error unmarshalling response body: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("Expected 1 user, got %v", len(users))
	}
	if users[0].ID != 1 {
		t.Errorf("Expected user ID to be 1, got %v", users[0].ID)
	}
	if users[0].Name != "John Doe" {
		t.Errorf("Expected user name to be 'John Doe', got %v", users[0].Name)
	}
}

func TestGetIDHandler(t *testing.T) {
	repository := repository.NewUserRepository(ConnTest, Context)
	service := services.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.GetID(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var user dto.Users
	err = json.Unmarshal(rr.Body.Bytes(), &user)
	if err != nil {
		t.Errorf("Error unmarshalling response body: %v", err)
	}
	if user.ID != 1 {
		t.Errorf("Expected user ID to be 1, got %v", user.ID)
	}
	if user.Name != "John Doe" {
		t.Errorf("Expected user name to be 'John Doe', got %v", user.Name)
	}
}

func TestCreateHandler(t *testing.T) {
	repository := repository.NewUserRepository(ConnTest, Context)
	service := services.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	user := dto.Users{Name: "Jane Doe", Hobby: "Swimming"}

	// Convert user to json
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.CreateUser(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var res dto.Users
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the user has been added to the mockDaoUser
	req, err = http.NewRequest("GET", "/users/"+strconv.Itoa(res.ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()

	handler.GetID(rr, req)

	if res.Name != user.Name {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			res.Name, user.Name)
	}

	if res.Hobby != user.Hobby {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			res.Hobby, user.Hobby)
	}
}

func TestUpdateHandler(t *testing.T) {
	repository := repository.NewUserRepository(ConnTest, Context)
	service := services.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	// Insert a user to update
	user := dto.Users{Name: "Jane Doe", Hobby: "Swimming"}
	_, err := service.Insert(user)
	if err != nil {
		t.Fatal(err)
	}

	// Prepare data for update
	userUpdate := dto.Users{Name: "John Doe", Hobby: "Running"}

	// Convert userUpdate to json
	userUpdateJSON, err := json.Marshal(userUpdate)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/users/"+strconv.Itoa(user.ID), bytes.NewBuffer(userUpdateJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.UpdateUser(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var res dto.Users
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatal(err)
	}

	if res.Name != userUpdate.Name {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			res.Name, userUpdate.Name)
	}

	if res.Hobby != userUpdate.Hobby {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			res.Hobby, userUpdate.Hobby)
	}

	// Check if the user has been updated in the mockDaoUser
	userCheck, err := service.GetID(user.ID)
	if err != nil {
		t.Fatal(err)
	}

	if userCheck.Name != userUpdate.Name {
		t.Error("User has not been updated in the mockDaoUser")
	}

	if userCheck.Hobby != userUpdate.Hobby {
		t.Error("User has not been updated in the mockDaoUser")
	}
}

func TestDeleteHandler(t *testing.T) {
	repository := repository.NewUserRepository(ConnTest, Context)
	service := services.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	// Create user to delete
	user := dto.Users{Name: "John Doe", Hobby: "Hiking"}
	res, _ := service.Insert(user)

	// Prepare request to delete the user
	req, err := http.NewRequest("DELETE", "/users/"+strconv.Itoa(res.ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Call the handler to delete the user
	handler.DeleteUser(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

	// Check if the user has been deleted from the mockDaoUser
	_, err = service.GetID(res.ID)
	if err == nil {
		t.Error("User has not been deleted from the userservice")
	}

}
