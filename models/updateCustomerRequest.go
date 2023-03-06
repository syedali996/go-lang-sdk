package models

// Request for updating a customer
type UpdateCustomerRequest struct {
    Name         string               `json:"name,omitempty"`
    Email        string               `json:"email,omitempty"`
    Document     string               `json:"document,omitempty"`
    Type         string               `json:"type,omitempty"`
    Address      CreateAddressRequest `json:"address,omitempty"`
    Metadata     map[string]string    `json:"metadata,omitempty"`
    Phones       CreatePhonesRequest  `json:"phones,omitempty"`
    Code         string               `json:"code,omitempty"`
    Gender       string               `json:"gender,omitempty"`
    DocumentType string               `json:"document_type,omitempty"`
}
