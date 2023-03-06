package models

// Update Order item Request
type UpdateOrderItemRequest struct {
    Amount      int    `json:"amount"`
    Description string `json:"description"`
    Quantity    int    `json:"quantity"`
    Category    string `json:"category"`
}
