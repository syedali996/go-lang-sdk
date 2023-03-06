package models

// Request for creating a location
type CreateLocationRequest struct {
    Latitude  string `json:"latitude"`
    Longitude string `json:"longitude"`
}
