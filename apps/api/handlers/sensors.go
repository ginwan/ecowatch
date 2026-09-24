package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"slices"
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

func CreateSensor(w http.ResponseWriter, r *http.Request) {
	var sensor models.Sensor
	// your fixed decode line here
	err := json.NewDecoder(r.Body).Decode(&sensor)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Unable to parse sensor data: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	data, err := os.ReadFile("data/sensors.json")
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

	maxID := 0
	for _, s := range sensors {
		if s.ID > maxID {
			maxID = s.ID
		}
	}
	newID := maxID + 1

	sensor.ID = newID
	sensors = append(sensors, sensor)
	// write the updated sensors slice back to the file
	addedData, err := json.MarshalIndent(sensors, "", "  ")
	if err != nil {
		http.Error(w, "Failed to marshal updated sensors data", http.StatusInternalServerError)
		return
	}
	err = os.WriteFile("data/sensors.json", addedData, 0644)
	if err != nil {
		http.Error(w, "Failed to write updated sensors data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sensor)
}

func UpdateSensor(w http.ResponseWriter, r *http.Request) {
	sensorID := r.PathValue("id")

	id, err := strconv.Atoi(sensorID)
	if err != nil {
		http.Error(w, "Invalid Sensor ID", http.StatusBadRequest)
		return
	}

	var sensor models.Sensor
	// your fixed decode line here
	err = json.NewDecoder(r.Body).Decode(&sensor)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Unable to parse sensor data: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

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

	for i, s := range sensors {
		if s.ID == id {
			sensor.ID = id
			sensors[i] = sensor

			// write the updated sensors slice back to the file
			addedData, err := json.MarshalIndent(sensors, "", "  ")
			if err != nil {
				http.Error(w, "Failed to marshal updated sensors data", http.StatusInternalServerError)
				return
			}
			err = os.WriteFile("data/sensors.json", addedData, 0644)
			if err != nil {
				http.Error(w, "Failed to write updated sensors data", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(sensor)
			return
		}
	}
	http.Error(w, "Sensor not found", http.StatusNotFound)
}


func DeleteSensor(w http.ResponseWriter, r *http.Request) {
	sensorID := r.PathValue("id")

	id, err := strconv.Atoi(sensorID)
	if err != nil {
		http.Error(w, "Invalid Sensor ID", http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile("data/sensors.json")
	fmt.Println("Reading file...")
	if err != nil{
		http.Error(w, "Can't read sensors data", http.StatusInternalServerError)
		return
	}

	var sensors []models.Sensor
	err = json.Unmarshal(data, &sensors)
	if err != nil{
		http.Error(w, 
			fmt.Sprintf("Unable to parse sensor data %v", err),
			http.StatusBadRequest)
		return
	}

	for i, s := range sensors{
		if(s.ID == id){
			sensors = slices.Delete(sensors, i, i+1)

			// write the updated sensors slice back to the file
			addedData, err := json.MarshalIndent(sensors, "", "  ")
			if err != nil {
				http.Error(w, "Failed to marshal updated sensors data", http.StatusInternalServerError)
				return
			}
			err = os.WriteFile("data/sensors.json", addedData, 0644)
			if err != nil {
				http.Error(w, "Failed to write updated sensors data", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(sensors)
			return
		}
	}
	http.Error(w, "Sensor not found", http.StatusNotFound)
}