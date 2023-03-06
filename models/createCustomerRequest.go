package models

// Request for creating a new customer
type CreateCustomerRequest struct {
    Name         string               `json:"name"`
    Email        string               `json:"email"`
    Document     string               `json:"document"`
    Type         string               `json:"type"`
    Address      CreateAddressRequest `json:"address"`
    Metadata     map[string]string    `json:"metadata"`
    Phones       CreatePhonesRequest  `json:"phones"`
    Code         string               `json:"code"`
    Gender       string               `json:"gender,omitempty"`
    DocumentType string               `json:"document_type,omitempty"`
}
