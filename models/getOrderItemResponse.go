package models

// Response object for getting an order item
type GetOrderItemResponse struct {
    Id          *string `json:"id"`
    Amount      *int    `json:"amount"`
    Description *string `json:"description"`
    Quantity    *int    `json:"quantity"`
    Category    *string `json:"category"`
    Code        *string `json:"code"`
}
