package models

// Object used for returning lists of objects with pagination
type PagingResponse struct {
    Total    *int    `json:"total"`
    Previous *string `json:"previous"`
    Next     *string `json:"next"`
}
