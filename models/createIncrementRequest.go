package models

// Request for creating a new increment
type CreateIncrementRequest struct {
    Value         float64 `json:"value"`
    IncrementType string  `json:"increment_type"`
    ItemId        string  `json:"item_id"`
    Cycles        int     `json:"cycles,omitempty"`
    Description   string  `json:"description,omitempty"`
}
