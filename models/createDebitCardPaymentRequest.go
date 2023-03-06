package models

// The settings for creating a debit card payment
type CreateDebitCardPaymentRequest struct {
    StatementDescriptor string                              `json:"statement_descriptor,omitempty"`
    Card                CreateCardRequest                   `json:"card,omitempty"`
    CardId              string                              `json:"card_id,omitempty"`
    CardToken           string                              `json:"card_token,omitempty"`
    Recurrence          bool                                `json:"recurrence,omitempty"`
    Authentication      CreatePaymentAuthenticationRequest  `json:"authentication,omitempty"`
    Token               CreateCardPaymentContactlessRequest `json:"token,omitempty"`
}
