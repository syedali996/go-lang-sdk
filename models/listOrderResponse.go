package models

// Response object for listing order objects
type ListOrderResponse struct {
    Data   *[]GetOrderResponse `json:"data"`
    Paging *PagingResponse     `json:"paging"`
}
