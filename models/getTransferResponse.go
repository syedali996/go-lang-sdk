package models

import (
    "encoding/json"
    "log"
    "time"
)

// Transfer response
type GetTransferResponse struct {
    Id          *string                 `json:"id"`
    Amount      *int                    `json:"amount"`
    Status      *string                 `json:"status"`
    CreatedAt   *time.Time              `json:"created_at"`
    UpdatedAt   *time.Time              `json:"updated_at"`
    BankAccount *GetBankAccountResponse `json:"bank_account"`
    Metadata    map[string]*string      `json:"metadata"`
}

func (g *GetTransferResponse) MarshalJSON() (
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
    type marshallable GetTransferResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"` 
        UpdatedAt    *string `json:"updatedAt"` 
    }{
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetTransferResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id          *string                 
        Amount      *int                    
        Status      *string                 
        CreatedAt   *string                 
        UpdatedAt   *string                 
        BankAccount *GetBankAccountResponse 
        Metadata    map[string]*string      
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Amount = temp.Amount
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
    g.BankAccount = temp.BankAccount
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    return nil
}
