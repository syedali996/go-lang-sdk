package models

import (
    "encoding/json"
    "log"
    "time"
)

// Recipient response
type GetRecipientResponse struct {
    Id                            *string                           `json:"id"`
    Name                          *string                           `json:"name"`
    Email                         *string                           `json:"email"`
    Document                      *string                           `json:"document"`
    Description                   *string                           `json:"description"`
    Type                          *string                           `json:"type"`
    Status                        *string                           `json:"status"`
    CreatedAt                     *time.Time                        `json:"created_at"`
    UpdatedAt                     *time.Time                        `json:"updated_at"`
    DeletedAt                     *time.Time                        `json:"deleted_at"`
    DefaultBankAccount            *GetBankAccountResponse           `json:"default_bank_account"`
    GatewayRecipients             *[]GetGatewayRecipientResponse    `json:"gateway_recipients"`
    Metadata                      map[string]*string                `json:"metadata"`
    AutomaticAnticipationSettings *GetAutomaticAnticipationResponse `json:"automatic_anticipation_settings"`
    TransferSettings              *GetTransferSettingsResponse      `json:"transfer_settings"`
    Code                          *string                           `json:"code"`
    PaymentMode                   *string                           `json:"payment_mode"`
}

func (g *GetRecipientResponse) MarshalJSON() (
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
    type marshallable GetRecipientResponse
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

func (g *GetRecipientResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                            *string                           
        Name                          *string                           
        Email                         *string                           
        Document                      *string                           
        Description                   *string                           
        Type                          *string                           
        Status                        *string                           
        CreatedAt                     *string                           
        UpdatedAt                     *string                           
        DeletedAt                     *string                           
        DefaultBankAccount            *GetBankAccountResponse           
        GatewayRecipients             *[]GetGatewayRecipientResponse    
        Metadata                      map[string]*string                
        AutomaticAnticipationSettings *GetAutomaticAnticipationResponse 
        TransferSettings              *GetTransferSettingsResponse      
        Code                          *string                           
        PaymentMode                   *string                           
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Name = temp.Name
    g.Email = temp.Email
    g.Document = temp.Document
    g.Description = temp.Description
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
    g.DefaultBankAccount = temp.DefaultBankAccount
    if temp.GatewayRecipients == nil {
        g.GatewayRecipients = nil
    } else {
        g.GatewayRecipients = temp.GatewayRecipients
    }
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    g.AutomaticAnticipationSettings = temp.AutomaticAnticipationSettings
    g.TransferSettings = temp.TransferSettings
    g.Code = temp.Code
    g.PaymentMode = temp.PaymentMode
    return nil
}
