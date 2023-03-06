package models

// Request for creating a bank account
type CreateBankAccountRequest struct {
    HolderName        string            `json:"holder_name"`
    HolderType        string            `json:"holder_type"`
    HolderDocument    string            `json:"holder_document"`
    Bank              string            `json:"bank"`
    BranchNumber      string            `json:"branch_number"`
    BranchCheckDigit  string            `json:"branch_check_digit"`
    AccountNumber     string            `json:"account_number"`
    AccountCheckDigit string            `json:"account_check_digit"`
    Type              string            `json:"type"`
    Metadata          map[string]string `json:"metadata"`
    PixKey            string            `json:"pix_key"`
}
