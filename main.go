package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renaldyhidayatt/crud_blog/config"
	"github.com/renaldyhidayatt/crud_blog/dto"
	"github.com/renaldyhidayatt/crud_blog/handler"
	"github.com/renaldyhidayatt/crud_blog/migrate"
	"github.com/renaldyhidayatt/crud_blog/repository"
	"github.com/renaldyhidayatt/crud_blog/services"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func main() {
	db, err := config.InitialDatabase()

	context := context.Background()

	if err != nil {
		log.Fatal(err.Error())
	}
	migrate.MigrationTable(db)

	log.Println("Start the dev server at http://127.0.0.1:8000")

	myRouter := mux.NewRouter().StrictSlash(true)

	repositoryUser := repository.NewUserRepository(db, context)
	serviceUser := services.NewUserService(repositoryUser)
	handler := handler.NewUserHandler(serviceUser)

	myRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := dto.Result{
			Code:    404,
			Message: "Method not found",
		}

		response, _ := json.Marshal(res)

		w.Write(response)
	})

	myRouter.MethodNotAllowedHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusMethodNotAllowed)

		res := dto.Result{Code: 403, Message: "Method not allowed"}
		response, _ := json.Marshal(res)
		rw.Write(response)
	})

	myRouter.HandleFunc("/", HomePage).Methods("GET")
	myRouter.HandleFunc("/user", handler.GetAll)
	myRouter.HandleFunc("/create", handler.CreateUser).Methods("POST")
	myRouter.HandleFunc("/{id}", handler.GetID).Methods("GET")
	myRouter.HandleFunc("/{id}", handler.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/{id}", handler.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
