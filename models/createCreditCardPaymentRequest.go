package models

// The settings for creating a credit card payment
type CreateCreditCardPaymentRequest struct {
    Installments         int                                 `json:"installments,omitempty"`
    StatementDescriptor  string                              `json:"statement_descriptor,omitempty"`
    Card                 CreateCardRequest                   `json:"card,omitempty"`
    CardId               string                              `json:"card_id,omitempty"`
    CardToken            string                              `json:"card_token,omitempty"`
    Recurrence           bool                                `json:"recurrence,omitempty"`
    Capture              bool                                `json:"capture,omitempty"`
    ExtendedLimitEnabled bool                                `json:"extended_limit_enabled,omitempty"`
    ExtendedLimitCode    string                              `json:"extended_limit_code,omitempty"`
    MerchantCategoryCode int64                               `json:"merchant_category_code,omitempty"`
    Authentication       CreatePaymentAuthenticationRequest  `json:"authentication,omitempty"`
    Contactless          CreateCardPaymentContactlessRequest `json:"contactless,omitempty"`
    AutoRecovery         bool                                `json:"auto_recovery,omitempty"`
    OperationType        string                              `json:"operation_type,omitempty"`
    RecurrencyCycle      string                              `json:"recurrency_cycle,omitempty"`
}
