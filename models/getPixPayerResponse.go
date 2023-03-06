package models

// Pix payer data.
type GetPixPayerResponse struct {
    Name         *string                    `json:"name"`
    Document     *string                    `json:"document"`
    DocumentType *string                    `json:"document_type"`
    BankAccount  *GetPixBankAccountResponse `json:"bank_account"`
}
