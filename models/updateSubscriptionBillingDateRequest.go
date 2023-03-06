package models

import (
    "time"
)

// Request for updating the due date from a subscription
type UpdateSubscriptionBillingDateRequest struct {
    NextBillingAt time.Time `json:"next_billing_at"`
}
