package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hbjydev/mangadex-next/controllers"
)

type UserRouter struct{}

func (u *UserRouter) getAll(w http.ResponseWriter, r *http.Request) {
	controller := controllers.UserController{}
	controller.GetAll(w, r)
}

func (u *UserRouter) getOne(w http.ResponseWriter, r *http.Request) {
	controller := controllers.UserController{}
	controller.GetOne(w, r)
}

func (u *UserRouter) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", u.getAll)
	r.HandleFunc("/users/{username}", u.getOne)
}
