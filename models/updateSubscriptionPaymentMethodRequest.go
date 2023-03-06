package models

// Request for updating a subscription's payment method
type UpdateSubscriptionPaymentMethodRequest struct {
    PaymentMethod string                          `json:"payment_method"`
    CardId        string                          `json:"card_id"`
    Card          CreateCardRequest               `json:"card"`
    CardToken     string                          `json:"card_token,omitempty"`
    Boleto        CreateSubscriptionBoletoRequest `json:"boleto,omitempty"`
}
