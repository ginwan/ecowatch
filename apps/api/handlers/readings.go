package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ginwan/ecowatch/apps/api/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReadingHandler struct {
	Pool *pgxpool.Pool
}

func (h *ReadingHandler) GetReadings(w http.ResponseWriter, r *http.Request) {
	rows, err := h.Pool.Query(context.Background(), "SELECT id, sensor_id, value, recorded_at, received_at FROM readings")
	if err != nil {
		http.Error(w, "Failed to get readings data", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var readings []models.Reading

	for rows.Next() {
		var reading models.Reading
		err := rows.Scan(
			&reading.ID,
			&reading.SensorID,
			&reading.Value,
			&reading.RecordedAt,
			&reading.ReceivedAt,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to access reading	 data %v", err), http.StatusInternalServerError)
			return
		}
		readings = append(readings, reading)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(readings)
}

func (h *ReadingHandler) CreateReading(w http.ResponseWriter, r *http.Request) {
	var reading models.Reading

	err := json.NewDecoder(r.Body).Decode(&reading)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to parse reading data %v", err), http.StatusInternalServerError)
		return
	}

	sqlString := `INSERT INTO readings(sensor_id, value, recorded_at)
		VALUES ($1, $2, $3) 
		RETURNING id, received_at`

	err = h.Pool.QueryRow(context.Background(), sqlString, reading.SensorID, reading.Value, reading.RecordedAt).Scan(&reading.ID, &reading.ReceivedAt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to create reading %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reading)
}
