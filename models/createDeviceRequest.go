package models

// Request for creating a device
type CreateDeviceRequest struct {
    Platform string `json:"platform,omitempty"`
}
