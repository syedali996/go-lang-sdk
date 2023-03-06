package models

// Request for updating the card from a subscription
type UpdateSubscriptionCardRequest struct {
    Card   CreateCardRequest `json:"card"`
    CardId string            `json:"card_id"`
}
