package models

type CreateWithdrawRequest struct {
    Amount   int               `json:"amount"`
    Metadata map[string]string `json:"metadata,omitempty"`
}
