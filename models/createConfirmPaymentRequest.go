package models

type CreateConfirmPaymentRequest struct {
    Description string `json:"description"`
    Amount      int    `json:"Amount,omitempty"`
    Code        string `json:"Code"`
}
