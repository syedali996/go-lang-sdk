package models

// Response object for listing subscription cycles
type ListCyclesResponse struct {
    Data   *[]GetPeriodResponse `json:"data"`
    Paging *PagingResponse      `json:"paging"`
}
