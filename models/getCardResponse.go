package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting a credit card
type GetCardResponse struct {
    Id             *string                    `json:"id"`
    LastFourDigits *string                    `json:"last_four_digits"`
    Brand          *string                    `json:"brand"`
    HolderName     *string                    `json:"holder_name"`
    ExpMonth       *int                       `json:"exp_month"`
    ExpYear        *int                       `json:"exp_year"`
    Status         *string                    `json:"status"`
    CreatedAt      *time.Time                 `json:"created_at"`
    UpdatedAt      *time.Time                 `json:"updated_at"`
    BillingAddress *GetBillingAddressResponse `json:"billing_address"`
    Customer       *GetCustomerResponse       `json:"customer"`
    Metadata       map[string]*string         `json:"metadata"`
    Type           *string                    `json:"type"`
    HolderDocument *string                    `json:"holder_document"`
    DeletedAt      *time.Time                 `json:"deleted_at"`
    FirstSixDigits *string                    `json:"first_six_digits"`
    Label          *string                    `json:"label"`
}

func (g *GetCardResponse) MarshalJSON() (
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
    type marshallable GetCardResponse
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

func (g *GetCardResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id             *string                    
        LastFourDigits *string                    
        Brand          *string                    
        HolderName     *string                    
        ExpMonth       *int                       
        ExpYear        *int                       
        Status         *string                    
        CreatedAt      *string                    
        UpdatedAt      *string                    
        BillingAddress *GetBillingAddressResponse 
        Customer       *GetCustomerResponse       
        Metadata       map[string]*string         
        Type           *string                    
        HolderDocument *string                    
        DeletedAt      *string                    
        FirstSixDigits *string                    
        Label          *string                    
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.LastFourDigits = temp.LastFourDigits
    g.Brand = temp.Brand
    g.HolderName = temp.HolderName
    g.ExpMonth = temp.ExpMonth
    g.ExpYear = temp.ExpYear
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
    g.BillingAddress = temp.BillingAddress
    g.Customer = temp.Customer
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    g.Type = temp.Type
    g.HolderDocument = temp.HolderDocument
    if temp.DeletedAt != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, *temp.DeletedAt)
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt = &DeletedAtVal
    } else {
        g.DeletedAt = nil
    }
    g.FirstSixDigits = temp.FirstSixDigits
    g.Label = temp.Label
    return nil
}
