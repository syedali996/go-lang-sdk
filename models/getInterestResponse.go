package models

// Interest Response
type GetInterestResponse struct {
    Days   *int    `json:"days"`
    Type   *string `json:"type"`
    Amount *int    `json:"amount"`
}
