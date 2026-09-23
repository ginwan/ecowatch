package models

// Sensor represents a single industrial sensor.
// Adjust fields here to match your existing demo JSON shape if it differs.
type Sensor struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"` // e.g. "temperature", "pressure", "gas"
	FacilityID   int     `json:"facility_id"`
	Unit         string  `json:"unit"` // e.g. "°C", "psi", "ppm"
	MinThreshold float64 `json:"min_threshold"`
	MaxThreshold float64 `json:"max_threshold"`
	Status       string  `json:"status"` // e.g. "active", "inactive"
}
