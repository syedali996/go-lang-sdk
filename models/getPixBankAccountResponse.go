package models

// Payer's bank details.
type GetPixBankAccountResponse struct {
    BankName      *string `json:"bank_name"`
    Ispb          *string `json:"ispb"`
    BranchCode    *string `json:"branch_code"`
    AccountNumber *string `json:"account_number"`
}
