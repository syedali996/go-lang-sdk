package models

// Response object for listing addresses
type ListAddressesResponse struct {
    Data   *[]GetAddressResponse `json:"data"`
    Paging *PagingResponse       `json:"paging"`
}
