package models

import (
    "encoding/json"
    "log"
    "time"
)

type GetBankAccountResponse struct {
    Id                *string               `json:"id"`
    HolderName        *string               `json:"holder_name"`
    HolderType        *string               `json:"holder_type"`
    Bank              *string               `json:"bank"`
    BranchNumber      *string               `json:"branch_number"`
    BranchCheckDigit  *string               `json:"branch_check_digit"`
    AccountNumber     *string               `json:"account_number"`
    AccountCheckDigit *string               `json:"account_check_digit"`
    Type              *string               `json:"type"`
    Status            *string               `json:"status"`
    CreatedAt         *time.Time            `json:"created_at"`
    UpdatedAt         *time.Time            `json:"updated_at"`
    DeletedAt         *time.Time            `json:"deleted_at"`
    Recipient         *GetRecipientResponse `json:"recipient"`
    Metadata          map[string]*string    `json:"metadata"`
    PixKey            *string               `json:"pix_key"`
}

func (g *GetBankAccountResponse) MarshalJSON() (
    []byte,
    error) {
    var CreatedAtVal *string
    if g.CreatedAt != nil {
        str := g.CreatedAt.Format(time.RFC3339)
        CreatedAtVal = &str
    } else {
        CreatedAtVal = nil
    }
    var UpdatedAtVal *string
    if g.UpdatedAt != nil {
        str := g.UpdatedAt.Format(time.RFC3339)
        UpdatedAtVal = &str
    } else {
        UpdatedAtVal = nil
    }
    var DeletedAtVal *string
    if g.DeletedAt != nil {
        str := g.DeletedAt.Format(time.RFC3339)
        DeletedAtVal = &str
    } else {
        DeletedAtVal = nil
    }
    type marshallable GetBankAccountResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"` 
        UpdatedAt    *string `json:"updatedAt"` 
        DeletedAt    *string `json:"deletedAt"` 
    }{
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        DeletedAt:    DeletedAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetBankAccountResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                *string               
        HolderName        *string               
        HolderType        *string               
        Bank              *string               
        BranchNumber      *string               
        BranchCheckDigit  *string               
        AccountNumber     *string               
        AccountCheckDigit *string               
        Type              *string               
        Status            *string               
        CreatedAt         *string               
        UpdatedAt         *string               
        DeletedAt         *string               
        Recipient         *GetRecipientResponse 
        Metadata          map[string]*string    
        PixKey            *string               
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.HolderName = temp.HolderName
    g.HolderType = temp.HolderType
    g.Bank = temp.Bank
    g.BranchNumber = temp.BranchNumber
    g.BranchCheckDigit = temp.BranchCheckDigit
    g.AccountNumber = temp.AccountNumber
    g.AccountCheckDigit = temp.AccountCheckDigit
    g.Type = temp.Type
    g.Status = temp.Status
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    } else {
        g.CreatedAt = nil
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        g.UpdatedAt = &UpdatedAtVal
    } else {
        g.UpdatedAt = nil
    }
    if temp.DeletedAt != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, *temp.DeletedAt)
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt = &DeletedAtVal
    } else {
        g.DeletedAt = nil
    }
    g.Recipient = temp.Recipient
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    g.PixKey = temp.PixKey
    return nil
}
