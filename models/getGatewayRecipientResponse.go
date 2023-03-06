package models

// Information about the recipient on the gateway
type GetGatewayRecipientResponse struct {
    Gateway   *string `json:"gateway"`
    Status    *string `json:"status"`
    Pgid      *string `json:"pgid"`
    CreatedAt *string `json:"created_at"`
    UpdatedAt *string `json:"updated_at"`
}
