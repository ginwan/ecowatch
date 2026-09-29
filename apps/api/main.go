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

	mux := http.NewServeMux()
	sensorHandler := &handlers.SensorHandler{Pool: pool}

	mux.HandleFunc("/health", handlers.GetHealth)
	// sensors routes
	mux.HandleFunc("GET /api/v1/sensors", sensorHandler.GetSensors)
	mux.HandleFunc("POST /api/v1/sensors", sensorHandler.CreateSensor)
	mux.HandleFunc("PUT /api/v1/sensors/{id}", sensorHandler.UpdateSensor)
	mux.HandleFunc("GET /api/v1/sensors/{id}", sensorHandler.GetSensorByID)
	mux.HandleFunc("DELETE /api/v1/sensors/{id}", sensorHandler.DeleteSensor)

	err = http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
