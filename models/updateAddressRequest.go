package models

// Request for updating an address
type UpdateAddressRequest struct {
    Number     string            `json:"number"`
    Complement string            `json:"complement"`
    Metadata   map[string]string `json:"metadata"`
    Line2      string            `json:"line_2"`
}
