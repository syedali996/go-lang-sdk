package models

type GetSubscriptionSplitResponse struct {
    Enabled *bool               `json:"enabled"`
    Rules   *[]GetSplitResponse `json:"rules"`
}
