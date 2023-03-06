package models

// Checkout bank transfer payment request
type CreateCheckoutBankTransferRequest struct {
    Bank    []string `json:"bank"`
    Retries int      `json:"retries"`
}
