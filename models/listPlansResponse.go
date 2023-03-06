package models

// Response object for listing plans
type ListPlansResponse struct {
    Data   *[]GetPlanResponse `json:"data"`
    Paging *PagingResponse    `json:"paging"`
}
