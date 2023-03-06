package models

// Checkout card payment request
type CreateCheckoutCreditCardPaymentRequest struct {
    StatementDescriptor string                                       `json:"statement_descriptor,omitempty"`
    Installments        []CreateCheckoutCardInstallmentOptionRequest `json:"installments,omitempty"`
    Authentication      CreatePaymentAuthenticationRequest           `json:"authentication,omitempty"`
    Capture             bool                                         `json:"capture,omitempty"`
}
