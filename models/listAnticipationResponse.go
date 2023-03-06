package models

// Anticipations
type ListAnticipationResponse struct {
    Data   *[]GetAnticipationResponse `json:"data"`
    Paging *PagingResponse            `json:"paging"`
}
