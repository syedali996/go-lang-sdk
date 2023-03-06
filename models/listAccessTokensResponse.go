package models

// Response object for listing access tokens
type ListAccessTokensResponse struct {
    Data   *[]GetAccessTokenResponse `json:"data"`
    Paging *PagingResponse           `json:"paging"`
}
