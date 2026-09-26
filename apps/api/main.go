package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ginwan/ecowatch/apps/api/handlers"
)

func main() {
	pool, err := connectDB()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()
	fmt.Println("Connected to database!")
	
	http.HandleFunc("/health", handlers.GetHealth)
	// sensors routes
	http.HandleFunc("GET /api/v1/sensors", handlers.GetSensors)
	http.HandleFunc("POST /api/v1/sensors", handlers.CreateSensor)
	http.HandleFunc("PUT /api/v1/sensors/{id}", handlers.UpdateSensor)
	http.HandleFunc("GET /api/v1/sensors/{id}", handlers.GetSensorByID)
	http.HandleFunc("DELETE /api/v1/sensors/{id}", handlers.DeleteSensor)

	err = http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
