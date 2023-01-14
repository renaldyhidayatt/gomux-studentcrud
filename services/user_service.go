package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/renaldyhidayatt/crud_blog/dao"
	"github.com/renaldyhidayatt/crud_blog/dto"
	"github.com/renaldyhidayatt/crud_blog/utils"
)

type userServices struct {
	user dao.DaoUser
}

func NewUserServices(user dao.DaoUser) *userServices {
	return &userServices{user: user}
}

func (s *userServices) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := s.user.GetAll()

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	utils.ResponseWithJSON(w, http.StatusOK, res)
}

func (s *userServices) GetID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := s.user.GetID(int(id))

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

func (s *userServices) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	var users dto.Users

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := s.user.Insert(&users)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	utils.ResponseWithJSON(w, http.StatusCreated, res)

}

func (s *userServices) UpdateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	var users dto.Users

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.Unmarshal([]byte(body), &users)

	res, err := s.user.Update(users)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, res)
}

func (s *userServices) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user dto.Users

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		panic(err.Error())
	}

	err = s.user.Delete(id)

	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusNoContent, user)
}
