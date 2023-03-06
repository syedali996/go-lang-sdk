package models

// Card data
type CreateCardRequest struct {
    Number           string                   `json:"number"`
    HolderName       string                   `json:"holder_name"`
    ExpMonth         int                      `json:"exp_month"`
    ExpYear          int                      `json:"exp_year"`
    Cvv              string                   `json:"cvv"`
    BillingAddress   CreateAddressRequest     `json:"billing_address"`
    Brand            string                   `json:"brand"`
    BillingAddressId string                   `json:"billing_address_id"`
    Metadata         map[string]string        `json:"metadata"`
    Type             string                   `json:"type"`
    Options          CreateCardOptionsRequest `json:"options"`
    HolderDocument   string                   `json:"holder_document,omitempty"`
    PrivateLabel     bool                     `json:"private_label"`
    Label            string                   `json:"label"`
    Id               string                   `json:"id,omitempty"`
    Token            string                   `json:"token,omitempty"`
}
