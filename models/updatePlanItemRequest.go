package models

// Request for updating a plan item
type UpdatePlanItemRequest struct {
    Name          string                     `json:"name"`
    Description   string                     `json:"description"`
    Status        string                     `json:"status"`
    PricingScheme UpdatePricingSchemeRequest `json:"pricing_scheme"`
    Quantity      int                        `json:"quantity,omitempty"`
    Cycles        int                        `json:"cycles,omitempty"`
}
