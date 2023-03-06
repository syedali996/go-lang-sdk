package models

// Request for creating a new Address
type CreateAddressRequest struct {
    Street       string            `json:"street"`
    Number       string            `json:"number"`
    ZipCode      string            `json:"zip_code"`
    Neighborhood string            `json:"neighborhood"`
    City         string            `json:"city"`
    State        string            `json:"state"`
    Country      string            `json:"country"`
    Complement   string            `json:"complement"`
    Metadata     map[string]string `json:"metadata"`
    Line1        string            `json:"line_1"`
    Line2        string            `json:"line_2"`
}
