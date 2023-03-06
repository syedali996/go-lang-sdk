package models

// Response object for getting a price bracket
type GetPriceBracketResponse struct {
    StartQuantity *int `json:"start_quantity"`
    Price         *int `json:"price"`
    EndQuantity   *int `json:"end_quantity"`
    OveragePrice  *int `json:"overage_price"`
}
