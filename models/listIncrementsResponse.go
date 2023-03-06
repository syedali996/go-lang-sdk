package models

type ListIncrementsResponse struct {
    Data   *[]GetIncrementResponse `json:"data"`
    Paging *PagingResponse         `json:"paging"`
}
