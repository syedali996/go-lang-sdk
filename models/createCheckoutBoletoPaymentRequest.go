package models

import (
    "time"
)

type CreateCheckoutBoletoPaymentRequest struct {
    Bank         string    `json:"bank"`
    Instructions string    `json:"instructions"`
    DueAt        time.Time `json:"due_at"`
}
