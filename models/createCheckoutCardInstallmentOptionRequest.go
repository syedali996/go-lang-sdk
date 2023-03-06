package models

// Options for card installment
type CreateCheckoutCardInstallmentOptionRequest struct {
    Number int `json:"number"`
    Total  int `json:"total"`
}
