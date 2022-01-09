package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renaldyhidayatt/crud_blog/config"
	"github.com/renaldyhidayatt/crud_blog/controllers"
	"github.com/renaldyhidayatt/crud_blog/dto"
)

func main() {
	config.InitialDatabase()

	log.Println("Start the dev server at http://127.0.0.1:8000")

	myRouter := mux.NewRouter().StrictSlash(true)

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

	myRouter.HandleFunc("/", controllers.HomePage).Methods("GET")
	myRouter.HandleFunc("/create", controllers.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
