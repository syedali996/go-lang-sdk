package models

type ListDiscountsResponse struct {
    Data   *[]GetDiscountResponse `json:"data"`
    Paging *PagingResponse        `json:"paging"`
}
