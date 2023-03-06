package models

// List of paginated transfer objects
type ListTransferResponse struct {
    Data   *[]GetTransferResponse `json:"data"`
    Paging *PagingResponse        `json:"paging"`
}
