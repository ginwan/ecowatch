package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ginwan/ecowatch/apps/api/models"
)

func GetSensors(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("data/sensors.json")
	fmt.Println("Reading sensors data from file...")
	if err != nil {
		http.Error(w, "Failed to read sensors data", http.StatusInternalServerError)
		return
	}

	var sensors []models.Sensor
	err = json.Unmarshal(data, &sensors)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Unable to parse sensor data: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(sensors)
}

func GetSensorByID(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("data/sensors.json")
	fmt.Println("Reading sensors data from file...")
	if err != nil {
		http.Error(w, "Failed to read sensors data", http.StatusInternalServerError)
		return
	}

	var sensors []models.Sensor
	err = json.Unmarshal(data, &sensors)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Unable to parse sensor data: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	sensorId := r.PathValue("id")
	fmt.Printf("Sensor ID %v\n", sensorId)
	// convert string id to int id
	id, err := strconv.Atoi(sensorId)
	if err != nil {
		http.Error(w, "Invalid Sensor ID", http.StatusBadRequest)
		return
	}

	for _, sensor := range sensors {
		if sensor.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(sensor)
			return
		}
	}
	http.Error(w, "Sensor not found", http.StatusNotFound)
}
