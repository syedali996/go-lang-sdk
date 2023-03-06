package models

type CreateSubscriptionSplitRequest struct {
    Enabled bool                 `json:"enabled"`
    Rules   []CreateSplitRequest `json:"rules"`
}
