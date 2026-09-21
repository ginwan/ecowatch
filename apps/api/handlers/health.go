package handlers

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	health := Health{
		Service: "EcoWatch API",
		Status:  "OK",
		Version: "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(health)
}
