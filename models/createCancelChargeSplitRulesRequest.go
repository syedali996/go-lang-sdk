package models

// Creates a refund with split rules
type CreateCancelChargeSplitRulesRequest struct {
    Id     string `json:"id"`
    Amount int    `json:"Amount"`
    Type   string `json:"type"`
}
