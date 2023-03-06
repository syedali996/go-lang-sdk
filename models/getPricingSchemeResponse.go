package models

// Response object for getting a pricing scheme
type GetPricingSchemeResponse struct {
    Price         *int                       `json:"price"`
    SchemeType    *string                    `json:"scheme_type"`
    PriceBrackets *[]GetPriceBracketResponse `json:"price_brackets"`
    MinimumPrice  *int                       `json:"minimum_price"`
    Percentage    *float64                   `json:"percentage"`
}
