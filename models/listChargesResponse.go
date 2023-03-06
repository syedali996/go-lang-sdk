package models

// Response object for listing charges
type ListChargesResponse struct {
    Data   *[]GetChargeResponse `json:"data"`
    Paging *PagingResponse      `json:"paging"`
}
