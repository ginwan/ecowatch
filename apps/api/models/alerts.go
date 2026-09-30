package models

import "time"

type Alert struct {
	ID        int       `json:"id"`
	SensorID  int       `json:"sensor_id"`
	Value     float64   `json:"value"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
