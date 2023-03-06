package models

// Request for creating a Setup for a subscription. The setup is an order that will be created at the subscription creation.
type CreateSetupRequest struct {
    Amount      int                  `json:"amount"`
    Description string               `json:"description"`
    Payment     CreatePaymentRequest `json:"payment"`
}
