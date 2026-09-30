package models

import "time"

type Reading struct {
	ID         int       `json:"id"`
	SensorID   int       `json:"sensor_id"`
	Value      float64   `json:"value"`
	RecordedAt time.Time `json:"recorded_at"`
	ReceivedAt time.Time `json:"received_at"`
}