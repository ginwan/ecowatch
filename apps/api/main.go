package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	health := Health{
		Service: "EcoWatch API",
		Status:  "OK",
		Version: "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(health)
}
func main() {
	fmt.Println("EcoWatch API is starting...")

	http.HandleFunc("/health", getHealth)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
