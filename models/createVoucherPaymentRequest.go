package models

// The settings for creating a voucher payment
type CreateVoucherPaymentRequest struct {
    StatementDescriptor string            `json:"statement_descriptor,omitempty"`
    CardId              string            `json:"card_id,omitempty"`
    CardToken           string            `json:"card_token,omitempty"`
    Card                CreateCardRequest `json:"Card,omitempty"`
    RecurrencyCycle     string            `json:"recurrency_cycle,omitempty"`
}
