package models

// Response object for getting the setup from a subscription
type GetSetupResponse struct {
    Id          *string `json:"id"`
    Description *string `json:"description"`
    Amount      *int    `json:"amount"`
    Status      *string `json:"status"`
}
