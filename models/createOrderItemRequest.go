package models

// Request for creating an order item
type CreateOrderItemRequest struct {
    Amount      int    `json:"amount"`
    Description string `json:"description"`
    Quantity    int    `json:"quantity"`
    Category    string `json:"category"`
    Code        string `json:"code,omitempty"`
}
