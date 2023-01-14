package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/renaldyhidayatt/crud_blog/dao"
	"github.com/renaldyhidayatt/crud_blog/dto"
	"github.com/renaldyhidayatt/crud_blog/utils"
)

type userHandler struct {
	user dao.DaoUser
}

func NewUserHandler(user dao.DaoUser) *userHandler {
	return &userHandler{user: user}
}

func (h *userHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.user.GetAll()

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	utils.ResponseWithJSON(w, http.StatusOK, res)
}

func (h *userHandler) GetID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.user.GetID(int(id))

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if (dto.Users{}) == res {
		utils.ResponseWithJSON(w, http.StatusNotFound, res)
	} else {
		utils.ResponseWithJSON(w, http.StatusOK, res)
	}
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	var users dto.Users

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.user.Insert(users)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	utils.ResponseWithJSON(w, http.StatusCreated, res)

}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var users dto.Users

	err = json.NewDecoder(r.Body).Decode(&users)

	users.ID = int(id)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.user.Update(users)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, res)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user dto.Users

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		panic(err.Error())
	}

	err = h.user.Delete(id)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusNoContent, user)
}
