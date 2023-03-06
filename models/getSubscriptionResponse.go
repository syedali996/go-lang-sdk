package models

import (
    "encoding/json"
    "log"
    "time"
)

type GetSubscriptionResponse struct {
    Id                   *string                        `json:"id"`
    Code                 *string                        `json:"code"`
    StartAt              *time.Time                     `json:"start_at"`
    Interval             *string                        `json:"interval"`
    IntervalCount        *int                           `json:"interval_count"`
    BillingType          *string                        `json:"billing_type"`
    CurrentCycle         *GetPeriodResponse             `json:"current_cycle"`
    PaymentMethod        *string                        `json:"payment_method"`
    Currency             *string                        `json:"currency"`
    Installments         *int                           `json:"installments"`
    Status               *string                        `json:"status"`
    CreatedAt            *time.Time                     `json:"created_at"`
    UpdatedAt            *time.Time                     `json:"updated_at"`
    Customer             *GetCustomerResponse           `json:"customer"`
    Card                 *GetCardResponse               `json:"card"`
    Items                *[]GetSubscriptionItemResponse `json:"items"`
    StatementDescriptor  *string                        `json:"statement_descriptor"`
    Metadata             map[string]*string             `json:"metadata"`
    Setup                *GetSetupResponse              `json:"setup"`
    GatewayAffiliationId *string                        `json:"gateway_affiliation_id"`
    NextBillingAt        *time.Time                     `json:"next_billing_at"`
    BillingDay           *int                           `json:"billing_day"`
    MinimumPrice         *int                           `json:"minimum_price"`
    CanceledAt           *time.Time                     `json:"canceled_at"`
    Discounts            *[]GetDiscountResponse         `json:"discounts"`
    Increments           *[]GetIncrementResponse        `json:"increments"`
    BoletoDueDays        *int                           `json:"boleto_due_days"`
    Split                *GetSubscriptionSplitResponse  `json:"split"`
    Boleto               *GetSubscriptionBoletoResponse `json:"boleto"`
}

func (g *GetSubscriptionResponse) MarshalJSON() (
    []byte,
    error) {
    var StartAtVal *string
    if g.StartAt != nil {
        str := g.StartAt.Format(time.RFC3339)
        StartAtVal = &str
    } else {
        StartAtVal = nil
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
    var NextBillingAtVal *string
    if g.NextBillingAt != nil {
        str := g.NextBillingAt.Format(time.RFC3339)
        NextBillingAtVal = &str
    } else {
        NextBillingAtVal = nil
    }
    var CanceledAtVal *string
    if g.CanceledAt != nil {
        str := g.CanceledAt.Format(time.RFC3339)
        CanceledAtVal = &str
    } else {
        CanceledAtVal = nil
    }
    type marshallable GetSubscriptionResponse
    return json.Marshal(struct {
        marshallable  
        StartAt       *string `json:"startAt"`       
        CreatedAt     *string `json:"createdAt"`     
        UpdatedAt     *string `json:"updatedAt"`     
        NextBillingAt *string `json:"nextBillingAt"` 
        CanceledAt    *string `json:"canceledAt"`    
    }{
        StartAt:       StartAtVal,       
        CreatedAt:     CreatedAtVal,     
        UpdatedAt:     UpdatedAtVal,     
        NextBillingAt: NextBillingAtVal, 
        CanceledAt:    CanceledAtVal,    
        marshallable:  marshallable(*g), 
    })
}

func (g *GetSubscriptionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                   *string                        
        Code                 *string                        
        StartAt              *string                        
        Interval             *string                        
        IntervalCount        *int                           
        BillingType          *string                        
        CurrentCycle         *GetPeriodResponse             
        PaymentMethod        *string                        
        Currency             *string                        
        Installments         *int                           
        Status               *string                        
        CreatedAt            *string                        
        UpdatedAt            *string                        
        Customer             *GetCustomerResponse           
        Card                 *GetCardResponse               
        Items                *[]GetSubscriptionItemResponse 
        StatementDescriptor  *string                        
        Metadata             map[string]*string             
        Setup                *GetSetupResponse              
        GatewayAffiliationId *string                        
        NextBillingAt        *string                        
        BillingDay           *int                           
        MinimumPrice         *int                           
        CanceledAt           *string                        
        Discounts            *[]GetDiscountResponse         
        Increments           *[]GetIncrementResponse        
        BoletoDueDays        *int                           
        Split                *GetSubscriptionSplitResponse  
        Boleto               *GetSubscriptionBoletoResponse 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Code = temp.Code
    if temp.StartAt != nil {
        StartAtVal, err := time.Parse(time.RFC3339, *temp.StartAt)
        if err != nil {
            log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
        }
        g.StartAt = &StartAtVal
    } else {
        g.StartAt = nil
    }
    g.Interval = temp.Interval
    g.IntervalCount = temp.IntervalCount
    g.BillingType = temp.BillingType
    g.CurrentCycle = temp.CurrentCycle
    g.PaymentMethod = temp.PaymentMethod
    g.Currency = temp.Currency
    g.Installments = temp.Installments
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
    g.Customer = temp.Customer
    g.Card = temp.Card
    if temp.Items == nil {
        g.Items = nil
    } else {
        g.Items = temp.Items
    }
    g.StatementDescriptor = temp.StatementDescriptor
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    g.Setup = temp.Setup
    g.GatewayAffiliationId = temp.GatewayAffiliationId
    if temp.NextBillingAt != nil {
        NextBillingAtVal, err := time.Parse(time.RFC3339, *temp.NextBillingAt)
        if err != nil {
            log.Fatalf("Cannot Parse next_billing_at as % s format.", time.RFC3339)
        }
        g.NextBillingAt = &NextBillingAtVal
    } else {
        g.NextBillingAt = nil
    }
    g.BillingDay = temp.BillingDay
    g.MinimumPrice = temp.MinimumPrice
    if temp.CanceledAt != nil {
        CanceledAtVal, err := time.Parse(time.RFC3339, *temp.CanceledAt)
        if err != nil {
            log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
        }
        g.CanceledAt = &CanceledAtVal
    } else {
        g.CanceledAt = nil
    }
    if temp.Discounts == nil {
        g.Discounts = nil
    } else {
        g.Discounts = temp.Discounts
    }
    if temp.Increments == nil {
        g.Increments = nil
    } else {
        g.Increments = temp.Increments
    }
    g.BoletoDueDays = temp.BoletoDueDays
    g.Split = temp.Split
    g.Boleto = temp.Boleto
    return nil
}
