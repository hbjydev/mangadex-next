package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hbjydev/mangadex-next/controllers"
)

type HealthRouter struct {
	Controller controllers.HealthController
}

func (h *HealthRouter) healthcheck(w http.ResponseWriter, r *http.Request) {
	controller := controllers.HealthController{}
	controller.Healthy(w, r)
}

func (h *HealthRouter) RegisterRoutes(r *mux.Router) {
	pathPrefix := r.PathPrefix("/-")

	pathPrefix.Path("/healthy").HandlerFunc(h.healthcheck).Methods("GET")
}
