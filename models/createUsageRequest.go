package models

import (
    "time"
)

// Request for creating a usage
type CreateUsageRequest struct {
    Quantity    int       `json:"quantity"`
    Description string    `json:"description"`
    UsedAt      time.Time `json:"used_at"`
    Code        string    `json:"code,omitempty"`
    Group       string    `json:"group,omitempty"`
    Amount      int       `json:"amount,omitempty"`
}
