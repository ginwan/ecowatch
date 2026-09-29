package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	// "slices"
	"strconv"

	"github.com/ginwan/ecowatch/apps/api/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SensorHandler struct {
	Pool *pgxpool.Pool
}

func (h *SensorHandler) GetSensors(w http.ResponseWriter, r *http.Request) {
	rows, err := h.Pool.Query(context.Background(), "SELECT id, name, type, facility_id, unit, min_threshold, max_threshold, status FROM sensors")
	if err != nil {
		http.Error(w, "Failed to sensors data", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var sensors []models.Sensor

	for rows.Next() {
		var sensor models.Sensor
		err := rows.Scan(
			&sensor.ID,
			&sensor.Name,
			&sensor.Type,
			&sensor.FacilityID,
			&sensor.Unit,
			&sensor.MinThreshold,
			&sensor.MaxThreshold,
			&sensor.Status)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Unable to access sensor data: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
		sensors = append(sensors, sensor)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sensors)
}

func (h *SensorHandler) CreateSensor(w http.ResponseWriter, r *http.Request) {
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

	sqlString := `INSERT INTO sensors (name, type, facility_id, unit, min_threshold, max_threshold, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`

	err = h.Pool.QueryRow(context.Background(), sqlString,
		sensor.Name, sensor.Type, sensor.FacilityID, sensor.Unit, sensor.MinThreshold, sensor.MaxThreshold, sensor.Status,
	).Scan(&sensor.ID)

	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Unable to write sensor data: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sensor)
}

func (h *SensorHandler) GetSensorByID(w http.ResponseWriter, r *http.Request) {
	sensorId := r.PathValue("id")

	id, err := strconv.Atoi(sensorId)
	if err != nil {
		http.Error(w, "Invalid Sensor ID", http.StatusBadRequest)
		return
	}

	var sensor models.Sensor
	sqlString := `SELECT name, type, facility_id, unit, min_threshold, max_threshold, status FROM sensors WHERE id = $1`

	err = h.Pool.QueryRow(context.Background(), sqlString, id).Scan(
		&sensor.Name,
		&sensor.Type,
		&sensor.FacilityID,
		&sensor.Unit,
		&sensor.MinThreshold,
		&sensor.MaxThreshold,
		&sensor.Status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Sensor not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get sensor data", http.StatusInternalServerError)
		return
	}
	sensor.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sensor)
}

func (h *SensorHandler) UpdateSensor(w http.ResponseWriter, r *http.Request) {
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

	sqlString := `UPDATE sensors 
		SET 
			name = $1, 
			type = $2, 
			facility_id = $3, 
			unit = $4, 
			min_threshold = $5, 
			max_threshold = $6, 
			status = $7 
		WHERE 
			id = $8`

	row, err := h.Pool.Exec(context.Background(), sqlString,
		sensor.Name, sensor.Type, sensor.FacilityID, sensor.Unit, sensor.MinThreshold, sensor.MaxThreshold, sensor.Status, id,
	)

	if err != nil {
		http.Error(w, "Update sensor data failed", http.StatusInternalServerError)
		return
	}

	if row.RowsAffected() == 0 {
		http.Error(w, "Sensor not found", http.StatusNotFound)
		return
	}

	sensor.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sensor)
}

func (h *SensorHandler) DeleteSensor(w http.ResponseWriter, r *http.Request) {
	sensorID := r.PathValue("id")

	id, err := strconv.Atoi(sensorID)
	if err != nil {
		http.Error(w, "Invalid Sensor ID", http.StatusBadRequest)
		return
	}

	sqlString := `DELETE FROM sensors WHERE id = $1`

	row, err := h.Pool.Exec(context.Background(), sqlString, id)

	if err != nil {
		http.Error(w, "Delete sensor failed", http.StatusInternalServerError)
		return
	}

	if row.RowsAffected() == 0 {
		http.Error(w, "Sensor not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
