package models

// Request for creating a bank transfer payment
type CreateBankTransferPaymentRequest struct {
    Bank    string `json:"bank"`
    Retries int    `json:"retries"`
}
