package models

// Response for the listing recipient method
type ListRecipientResponse struct {
    Data   *[]GetRecipientResponse `json:"data"`
    Paging *PagingResponse         `json:"paging"`
}
