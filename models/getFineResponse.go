package models

// Fine Response
type GetFineResponse struct {
    Days   *int    `json:"days"`
    Type   *string `json:"type"`
    Amount *int    `json:"amount"`
}
