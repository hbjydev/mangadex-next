package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hbjydev/mangadex-next/controllers"
)

type HealthRouter struct{}

func (h *HealthRouter) healthcheck(w http.ResponseWriter, r *http.Request) {
	controller := controllers.HealthController{}
	controller.Healthy(w, r)
}

func (h *HealthRouter) metrics(w http.ResponseWriter, r *http.Request) {
	controller := controllers.HealthController{}
	controller.Metrics(w, r)
}

func (h *HealthRouter) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/-/healthy", h.healthcheck)
	r.HandleFunc("/-/metrics", h.metrics)
}
