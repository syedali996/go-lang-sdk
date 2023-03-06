package models

// Request for creating a plan item
type CreatePlanItemRequest struct {
    Name          string                     `json:"name"`
    PricingScheme CreatePricingSchemeRequest `json:"pricing_scheme"`
    Id            string                     `json:"id"`
    Description   string                     `json:"description"`
    Cycles        int                        `json:"cycles,omitempty"`
    Quantity      int                        `json:"quantity,omitempty"`
}
