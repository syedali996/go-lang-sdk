package models

// Request for creating a new subscription item
type CreateSubscriptionItemRequest struct {
    Description   string                     `json:"description"`
    PricingScheme CreatePricingSchemeRequest `json:"pricing_scheme"`
    Id            string                     `json:"id"`
    PlanItemId    string                     `json:"plan_item_id"`
    Discounts     []CreateDiscountRequest    `json:"discounts"`
    Name          string                     `json:"name"`
    Cycles        int                        `json:"cycles,omitempty"`
    Quantity      int                        `json:"quantity,omitempty"`
    MinimumPrice  int                        `json:"minimum_price,omitempty"`
}
