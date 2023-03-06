package models

// Split response
type GetSplitResponse struct {
    Type      *string                  `json:"type"`
    Amount    *int                     `json:"amount"`
    Recipient *GetRecipientResponse    `json:"recipient"`
    GatewayId *string                  `json:"gateway_id"`
    Options   *GetSplitOptionsResponse `json:"options"`
    Id        *string                  `json:"id"`
}
