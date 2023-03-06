package models

// Response object for geetting an order location request
type GetLocationResponse struct {
    Latitude  *string `json:"latitude"`
    Longitude *string `json:"longitude"`
}
