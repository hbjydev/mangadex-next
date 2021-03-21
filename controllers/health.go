package controllers

import "net/http"

type HealthController struct{}

func (h *HealthController) Healthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
