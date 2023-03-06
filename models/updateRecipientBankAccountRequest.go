package models

// Updates the default bank account for a recipient
type UpdateRecipientBankAccountRequest struct {
    BankAccount CreateBankAccountRequest `json:"bank_account"`
    PaymentMode string                   `json:"payment_mode"`
}
