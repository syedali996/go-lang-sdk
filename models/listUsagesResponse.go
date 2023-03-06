package models

// Response model for listing the usages from a subscription item
type ListUsagesResponse struct {
    Data   *[]GetUsageResponse `json:"data"`
    Paging *PagingResponse     `json:"paging"`
}
