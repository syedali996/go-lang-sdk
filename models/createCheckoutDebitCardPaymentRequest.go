package models

// Checkout credit card payment request
type CreateCheckoutDebitCardPaymentRequest struct {
    StatementDescriptor string                             `json:"statement_descriptor,omitempty"`
    Authentication      CreatePaymentAuthenticationRequest `json:"authentication"`
}
