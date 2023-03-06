package models

// Request for updating a pricing scheme
type UpdatePricingSchemeRequest struct {
    SchemeType    string                      `json:"scheme_type"`
    PriceBrackets []UpdatePriceBracketRequest `json:"price_brackets"`
    Price         int                         `json:"price,omitempty"`
    MinimumPrice  int                         `json:"minimum_price,omitempty"`
    Percentage    float64                     `json:"percentage,omitempty"`
}
