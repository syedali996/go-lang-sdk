package models

// Request for updating a Recipient
type UpdateRecipientRequest struct {
    Name        string            `json:"name"`
    Email       string            `json:"email"`
    Description string            `json:"description"`
    Type        string            `json:"type"`
    Status      string            `json:"status"`
    Metadata    map[string]string `json:"metadata"`
}
