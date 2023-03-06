package models

// Fine Request
type CreateFineRequest struct {
    Days   int    `json:"days"`
    Type   string `json:"type"`
    Amount int    `json:"amount"`
}
