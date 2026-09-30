package models

type Facility struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"` // e.g. "factory", "warehouse", "office"
	Location string `json:"location"`
	Status   string `json:"status"` // e.g. "active", "inactive"
}
