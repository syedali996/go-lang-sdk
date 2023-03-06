package models

// Bank transfer checkout response
type GetCheckoutBankTransferPaymentResponse struct {
    Bank *[]string `json:"bank"`
}
