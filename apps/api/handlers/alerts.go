package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ginwan/ecowatch/apps/api/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AlertHandler struct {
	Pool *pgxpool.Pool
}

func (h *AlertHandler) GetAlerts(w http.ResponseWriter, r *http.Request) {
	rows, err := h.Pool.Query(context.Background(), "SELECT id, sensor_id, value, status, message, created_at FROM alerts")
	if err != nil {
		http.Error(w, "Failed to get alert data", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var alerts []models.Alert

	for rows.Next() {
		var alert models.Alert
		err := rows.Scan(
			&alert.ID,
			&alert.SensorID,
			&alert.Value,
			&alert.Status,
			&alert.Message,
			&alert.CreatedAt,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to access alert data %v", err), http.StatusInternalServerError)
			return
		}
		alerts = append(alerts, alert)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(alerts)
}
