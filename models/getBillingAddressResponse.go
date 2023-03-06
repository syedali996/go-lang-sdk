package models

// Response object for getting a billing address
type GetBillingAddressResponse struct {
    Street       *string `json:"street"`
    Number       *string `json:"number"`
    ZipCode      *string `json:"zip_code"`
    Neighborhood *string `json:"neighborhood"`
    City         *string `json:"city"`
    State        *string `json:"state"`
    Country      *string `json:"country"`
    Complement   *string `json:"complement"`
    Line1        *string `json:"line_1"`
    Line2        *string `json:"line_2"`
}
