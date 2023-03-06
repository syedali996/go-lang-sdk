package models

// Request for updating card data
type UpdateChargeCardRequest struct {
    UpdateSubscription bool              `json:"update_subscription"`
    CardId             string            `json:"card_id"`
    Card               CreateCardRequest `json:"card"`
    Recurrence         bool              `json:"recurrence"`
}
