package models

// Atualização do valor mínimo da assinatura
type UpdateSubscriptionMinimumPriceRequest struct {
    MinimumPrice int `json:"minimum_price,omitempty"`
}
