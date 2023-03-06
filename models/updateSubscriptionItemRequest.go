package models

// Request for updating a subscription item
type UpdateSubscriptionItemRequest struct {
    Description   string                     `json:"description"`
    Status        string                     `json:"status"`
    PricingScheme UpdatePricingSchemeRequest `json:"pricing_scheme"`
    Name          string                     `json:"name"`
    Cycles        int                        `json:"cycles,omitempty"`
    Quantity      int                        `json:"quantity,omitempty"`
    MinimumPrice  int                        `json:"minimum_price,omitempty"`
}
