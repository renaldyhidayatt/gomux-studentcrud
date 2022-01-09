package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/renaldyhidayatt/crud_blog/dto"
	"github.com/renaldyhidayatt/crud_blog/repository"
	"github.com/renaldyhidayatt/crud_blog/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var user dto.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseJson(w, err, http.StatusBadRequest)
		return
	}

	if err := repository.Insert(ctx, user); err != nil {
		utils.ResponseJson(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJson(w, res, http.StatusCreated)

}
