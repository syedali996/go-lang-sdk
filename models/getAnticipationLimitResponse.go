package models

// Anticipation limit
type GetAnticipationLimitResponse struct {
    Amount          *int `json:"amount"`
    AnticipationFee *int `json:"anticipation_fee"`
}
