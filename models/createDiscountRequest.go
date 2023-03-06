package models

// Request for creating a new discount
type CreateDiscountRequest struct {
    Value        float64 `json:"value"`
    DiscountType string  `json:"discount_type"`
    ItemId       string  `json:"item_id"`
    Cycles       int     `json:"cycles,omitempty"`
    Description  string  `json:"description,omitempty"`
}
