package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ginwan/ecowatch/apps/api/handlers"
)

func main() {
	fmt.Println("EcoWatch API is starting...")

	http.HandleFunc("/health", handlers.GetHealth)
	// sensors routes
	http.HandleFunc("GET /api/v1/sensors", handlers.GetSensors)
	http.HandleFunc("POST /api/v1/sensors", handlers.CreateSensor)
	http.HandleFunc("PUT /api/v1/sensors/{id}", handlers.UpdateSensor)
	http.HandleFunc("GET /api/v1/sensors/{id}", handlers.GetSensorByID)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
