package models

// Request for creating a new Invoice
type CreateInvoiceRequest struct {
    Metadata map[string]string `json:"metadata"`
}
