package models

// Request for creating a new Access Token
type CreateAccessTokenRequest struct {
    ExpiresIn int `json:"expires_in,omitempty"`
}
