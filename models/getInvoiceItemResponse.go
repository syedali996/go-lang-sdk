package models

// Response object for getting an invoice item
type GetInvoiceItemResponse struct {
    Amount             *int                      `json:"amount"`
    Description        *string                   `json:"description"`
    PricingScheme      *GetPricingSchemeResponse `json:"pricing_scheme"`
    PriceBracket       *GetPriceBracketResponse  `json:"price_bracket"`
    Quantity           *int                      `json:"quantity"`
    Name               *string                   `json:"name"`
    SubscriptionItemId *string                   `json:"subscription_item_id"`
}
