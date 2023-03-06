package models

// The settings for creating a private label payment
type CreatePrivateLabelPaymentRequest struct {
    Installments         int               `json:"installments,omitempty"`
    StatementDescriptor  string            `json:"statement_descriptor,omitempty"`
    Card                 CreateCardRequest `json:"card,omitempty"`
    CardId               string            `json:"card_id,omitempty"`
    CardToken            string            `json:"card_token,omitempty"`
    Recurrence           bool              `json:"recurrence,omitempty"`
    Capture              bool              `json:"capture,omitempty"`
    ExtendedLimitEnabled bool              `json:"extended_limit_enabled,omitempty"`
    ExtendedLimitCode    string            `json:"extended_limit_code,omitempty"`
    RecurrencyCycle      string            `json:"recurrency_cycle,omitempty"`
}
