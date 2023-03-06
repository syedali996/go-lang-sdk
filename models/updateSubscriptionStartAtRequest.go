package models

import (
    "time"
)

// Request for updating the start date from a subscription
type UpdateSubscriptionStartAtRequest struct {
    StartAt time.Time `json:"start_at"`
}
