package models

// Request for creating a recipient
type CreateRecipientRequest struct {
    Name               string                        `json:"name"`
    Email              string                        `json:"email"`
    Description        string                        `json:"description"`
    Document           string                        `json:"document"`
    Type               string                        `json:"type"`
    DefaultBankAccount CreateBankAccountRequest      `json:"default_bank_account"`
    Metadata           map[string]string             `json:"metadata"`
    TransferSettings   CreateTransferSettingsRequest `json:"transfer_settings,omitempty"`
    Code               string                        `json:"code"`
    PaymentMode        string                        `json:"payment_mode"`
}
