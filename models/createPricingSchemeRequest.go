package models

// Request for creating a pricing scheme
type CreatePricingSchemeRequest struct {
    SchemeType    string                      `json:"scheme_type"`
    PriceBrackets []CreatePriceBracketRequest `json:"price_brackets,omitempty"`
    Price         int                         `json:"price,omitempty"`
    MinimumPrice  int                         `json:"minimum_price,omitempty"`
    Percentage    float64                     `json:"percentage,omitempty"`
}
