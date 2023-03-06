package models

// Request for updating a price bracket
type UpdatePriceBracketRequest struct {
    StartQuantity int `json:"start_quantity"`
    Price         int `json:"price"`
    EndQuantity   int `json:"end_quantity,omitempty"`
    OveragePrice  int `json:"overage_price,omitempty"`
}
