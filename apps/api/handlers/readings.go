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

	var minThreshold float64
	var maxThreshold float64

	sqlStr := `SELECT min_threshold, max_threshold FROM sensors WHERE id = $1`;
	err = h.Pool.QueryRow(context.Background(), sqlStr, reading.SensorID).Scan(&minThreshold, &maxThreshold)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't access sensors data %v", err), http.StatusInternalServerError)
		return
	}

	if(reading.Value > maxThreshold){
		sqlString := `INSERT INTO alerts(sensor_id, value, status, message)
		VALUES ($1, $2, $3, $4)`

		_, err := h.Pool.Exec(context.Background(), sqlString, reading.SensorID, reading.Value, "high", "Recorded high value")
		if err != nil {
		http.Error(w, fmt.Sprintf("Unable to create alert %v", err), http.StatusInternalServerError)
		return
		}
	
	}else if (reading.Value < minThreshold){
		sqlString := `INSERT INTO alerts(sensor_id, value, status, message)
		VALUES ($1, $2, $3, $4) `

		_, err := h.Pool.Exec(context.Background(), sqlString, reading.SensorID, reading.Value, "low", "Recorded low value")
		if err != nil {
		http.Error(w, fmt.Sprintf("Unable to create alert %v", err), http.StatusInternalServerError)
		return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reading)
}
