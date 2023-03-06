package models

// Request for updating a card
type UpdateCardRequest struct {
    HolderName       string               `json:"holder_name"`
    ExpMonth         int                  `json:"exp_month"`
    ExpYear          int                  `json:"exp_year"`
    BillingAddressId string               `json:"billing_address_id"`
    BillingAddress   CreateAddressRequest `json:"billing_address"`
    Metadata         map[string]string    `json:"metadata"`
    Label            string               `json:"label"`
}
