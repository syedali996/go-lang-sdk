package models

// Response object for listing cards
type ListCardsResponse struct {
    Data   *[]GetCardResponse `json:"data"`
    Paging *PagingResponse    `json:"paging"`
}
