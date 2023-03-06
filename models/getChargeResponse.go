package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting a charge
type GetChargeResponse struct {
    Id                  *string                         `json:"id"`
    Code                *string                         `json:"code"`
    GatewayId           *string                         `json:"gateway_id"`
    Amount              *int                            `json:"amount"`
    Status              *string                         `json:"status"`
    Currency            *string                         `json:"currency"`
    PaymentMethod       *string                         `json:"payment_method"`
    DueAt               *time.Time                      `json:"due_at"`
    CreatedAt           *time.Time                      `json:"created_at"`
    UpdatedAt           *time.Time                      `json:"updated_at"`
    LastTransaction     GetTransactionResponseInterface `json:"last_transaction"`
    Invoice             *GetInvoiceResponse             `json:"invoice"`
    Order               *GetOrderResponse               `json:"order"`
    Customer            *GetCustomerResponse            `json:"customer"`
    Metadata            map[string]*string              `json:"metadata"`
    PaidAt              *time.Time                      `json:"paid_at"`
    CanceledAt          *time.Time                      `json:"canceled_at"`
    CanceledAmount      *int                            `json:"canceled_amount"`
    PaidAmount          *int                            `json:"paid_amount"`
    InterestAndFinePaid *int                            `json:"interest_and_fine_paid"`
    RecurrencyCycle     *string                         `json:"recurrency_cycle"`
}

func (g *GetChargeResponse) MarshalJSON() (
    []byte,
    error) {
    var DueAtVal *string
    if g.DueAt != nil {
        str := g.DueAt.Format(time.RFC3339)
        DueAtVal = &str
    } else {
        DueAtVal = nil
    }
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
    var PaidAtVal *string
    if g.PaidAt != nil {
        str := g.PaidAt.Format(time.RFC3339)
        PaidAtVal = &str
    } else {
        PaidAtVal = nil
    }
    var CanceledAtVal *string
    if g.CanceledAt != nil {
        str := g.CanceledAt.Format(time.RFC3339)
        CanceledAtVal = &str
    } else {
        CanceledAtVal = nil
    }
    type marshallable GetChargeResponse
    return json.Marshal(struct {
        marshallable 
        DueAt        *string `json:"dueAt"`      
        CreatedAt    *string `json:"createdAt"`  
        UpdatedAt    *string `json:"updatedAt"`  
        PaidAt       *string `json:"paidAt"`     
        CanceledAt   *string `json:"canceledAt"` 
    }{
        DueAt:        DueAtVal,         
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        PaidAt:       PaidAtVal,        
        CanceledAt:   CanceledAtVal,    
        marshallable: marshallable(*g), 
    })
}

func (g *GetChargeResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        LastTransaction     GetTransactionResponseField 
        Id                  *string                     
        Code                *string                     
        GatewayId           *string                     
        Amount              *int                        
        Status              *string                     
        Currency            *string                     
        PaymentMethod       *string                     
        DueAt               *string                     
        CreatedAt           *string                     
        UpdatedAt           *string                     
        Invoice             *GetInvoiceResponse         
        Order               *GetOrderResponse           
        Customer            *GetCustomerResponse        
        Metadata            map[string]*string          
        PaidAt              *string                     
        CanceledAt          *string                     
        CanceledAmount      *int                        
        PaidAmount          *int                        
        InterestAndFinePaid *int                        
        RecurrencyCycle     *string                     
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.LastTransaction = temp.LastTransaction.GetTransactionResponseInterface
    g.Id = temp.Id
    g.Code = temp.Code
    g.GatewayId = temp.GatewayId
    g.Amount = temp.Amount
    g.Status = temp.Status
    g.Currency = temp.Currency
    g.PaymentMethod = temp.PaymentMethod
    if temp.DueAt != nil {
        DueAtVal, err := time.Parse(time.RFC3339, *temp.DueAt)
        if err != nil {
            log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
        }
        g.DueAt = &DueAtVal
    } else {
        g.DueAt = nil
    }
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
    g.Invoice = temp.Invoice
    g.Order = temp.Order
    g.Customer = temp.Customer
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    if temp.PaidAt != nil {
        PaidAtVal, err := time.Parse(time.RFC3339, *temp.PaidAt)
        if err != nil {
            log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
        }
        g.PaidAt = &PaidAtVal
    } else {
        g.PaidAt = nil
    }
    if temp.CanceledAt != nil {
        CanceledAtVal, err := time.Parse(time.RFC3339, *temp.CanceledAt)
        if err != nil {
            log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
        }
        g.CanceledAt = &CanceledAtVal
    } else {
        g.CanceledAt = nil
    }
    g.CanceledAmount = temp.CanceledAmount
    g.PaidAmount = temp.PaidAmount
    g.InterestAndFinePaid = temp.InterestAndFinePaid
    g.RecurrencyCycle = temp.RecurrencyCycle
    return nil
}
