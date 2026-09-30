package models

type Reading struct {
	SensorID   int     `json:"sensor_id"`
	Value      float64 `json:"value"`
	RecordedAt string  `json:"recorded_at"`
}
