package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hbjydev/mangadex-next/models"
)

type UserController struct{}

func (u *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		log.Println("/users: failed to marshal users array")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}

func (u *UserController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := models.UserByUsername(vars["username"])
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	userJSON, err := user.Normalize()
	if err != nil {
		log.Printf("/user/%v: failed to marshal user", vars["username"])
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(*userJSON))
}
