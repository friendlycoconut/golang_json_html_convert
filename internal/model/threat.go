package model

// Variant represents a single variant entry from the input JSON.
// Variant is a low level structure
type Variant struct {
	Name      string `json:"name"`
	DateAdded string `json:"dateAdded"`
}

// Threat is the top-level structure for the expected JSON.
// Missing fields remain at zero values; unknown fields are ignored by encoding/json.
type Threat struct {
	ThreatName    string    `json:"threatName"`
	Category      string    `json:"category"`
	Size          int       `json:"size"`
	DetectionDate string    `json:"detectionDate"`
	Variants      []Variant `json:"variants"`
}
