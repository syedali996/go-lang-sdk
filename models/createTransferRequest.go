package models

// Request for creating a transfer
type CreateTransferRequest struct {
    Amount   int               `json:"amount"`
    Metadata map[string]string `json:"metadata"`
}
