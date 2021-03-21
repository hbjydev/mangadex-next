package controllers

import (
	"fmt"
	"net/http"

	"github.com/hbjydev/mangadex-next/database"
)

type HealthController struct{}

func (h *HealthController) Healthy(w http.ResponseWriter, r *http.Request) {
	w.Header().Del("Content-Type")
	w.WriteHeader(http.StatusOK)
}

func (h *HealthController) Metrics(w http.ResponseWriter, r *http.Request) {
	var output string

	output += fmt.Sprintf("mangadex_item_count{type=\"user\"} %v\n", countUsers())

	w.Header().Del("Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

func countUsers() int {
	row := database.DB.QueryRow("SELECT count(id) FROM users")
	var count int
	if err := row.Scan(&count); err != nil {
		return -1
	}
	return count
}
