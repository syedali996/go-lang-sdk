package models

import (
    "time"
)

// Request for creating an anticipation
type CreateAnticipationRequest struct {
    Amount      int       `json:"amount"`
    Timeframe   string    `json:"timeframe"`
    PaymentDate time.Time `json:"payment_date"`
}
