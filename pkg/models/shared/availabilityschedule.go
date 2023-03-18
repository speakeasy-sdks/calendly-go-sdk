package shared

// AvailabilitySchedule
// The availability schedule set by the user.
type AvailabilitySchedule struct {
	Default  bool               `json:"default"`
	Name     string             `json:"name"`
	Rules    []AvailabilityRule `json:"rules"`
	Timezone string             `json:"timezone"`
	URI      string             `json:"uri"`
	User     string             `json:"user"`
}
