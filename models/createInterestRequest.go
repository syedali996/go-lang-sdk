package models

// Interest Request
type CreateInterestRequest struct {
    Days   int    `json:"days"`
    Type   string `json:"type"`
    Amount int    `json:"amount"`
}
