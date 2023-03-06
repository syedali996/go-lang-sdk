package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting an invoice
type GetInvoiceResponse struct {
    Id             *string                    `json:"id"`
    Code           *string                    `json:"code"`
    Url            *string                    `json:"url"`
    Amount         *int                       `json:"amount"`
    Status         *string                    `json:"status"`
    PaymentMethod  *string                    `json:"payment_method"`
    CreatedAt      *time.Time                 `json:"created_at"`
    Items          *[]GetInvoiceItemResponse  `json:"items"`
    Customer       *GetCustomerResponse       `json:"customer"`
    Charge         *GetChargeResponse         `json:"charge"`
    Installments   *int                       `json:"installments"`
    BillingAddress *GetBillingAddressResponse `json:"billing_address"`
    Subscription   *GetSubscriptionResponse   `json:"subscription"`
    Cycle          *GetPeriodResponse         `json:"cycle"`
    Shipping       *GetShippingResponse       `json:"shipping"`
    Metadata       map[string]*string         `json:"metadata"`
    DueAt          *time.Time                 `json:"due_at"`
    CanceledAt     *time.Time                 `json:"canceled_at"`
    BillingAt      *time.Time                 `json:"billing_at"`
    SeenAt         *time.Time                 `json:"seen_at"`
    TotalDiscount  *int                       `json:"total_discount"`
    TotalIncrement *int                       `json:"total_increment"`
    SubscriptionId *string                    `json:"subscription_id"`
}

func (g *GetInvoiceResponse) MarshalJSON() (
    []byte,
    error) {
    var CreatedAtVal *string
    if g.CreatedAt != nil {
        str := g.CreatedAt.Format(time.RFC3339)
        CreatedAtVal = &str
    } else {
        CreatedAtVal = nil
    }
    var DueAtVal *string
    if g.DueAt != nil {
        str := g.DueAt.Format(time.RFC3339)
        DueAtVal = &str
    } else {
        DueAtVal = nil
    }
    var CanceledAtVal *string
    if g.CanceledAt != nil {
        str := g.CanceledAt.Format(time.RFC3339)
        CanceledAtVal = &str
    } else {
        CanceledAtVal = nil
    }
    var BillingAtVal *string
    if g.BillingAt != nil {
        str := g.BillingAt.Format(time.RFC3339)
        BillingAtVal = &str
    } else {
        BillingAtVal = nil
    }
    var SeenAtVal *string
    if g.SeenAt != nil {
        str := g.SeenAt.Format(time.RFC3339)
        SeenAtVal = &str
    } else {
        SeenAtVal = nil
    }
    type marshallable GetInvoiceResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"`  
        DueAt        *string `json:"dueAt"`      
        CanceledAt   *string `json:"canceledAt"` 
        BillingAt    *string `json:"billingAt"`  
        SeenAt       *string `json:"seenAt"`     
    }{
        CreatedAt:    CreatedAtVal,     
        DueAt:        DueAtVal,         
        CanceledAt:   CanceledAtVal,    
        BillingAt:    BillingAtVal,     
        SeenAt:       SeenAtVal,        
        marshallable: marshallable(*g), 
    })
}

func (g *GetInvoiceResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id             *string                    
        Code           *string                    
        Url            *string                    
        Amount         *int                       
        Status         *string                    
        PaymentMethod  *string                    
        CreatedAt      *string                    
        Items          *[]GetInvoiceItemResponse  
        Customer       *GetCustomerResponse       
        Charge         *GetChargeResponse         
        Installments   *int                       
        BillingAddress *GetBillingAddressResponse 
        Subscription   *GetSubscriptionResponse   
        Cycle          *GetPeriodResponse         
        Shipping       *GetShippingResponse       
        Metadata       map[string]*string         
        DueAt          *string                    
        CanceledAt     *string                    
        BillingAt      *string                    
        SeenAt         *string                    
        TotalDiscount  *int                       
        TotalIncrement *int                       
        SubscriptionId *string                    
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Code = temp.Code
    g.Url = temp.Url
    g.Amount = temp.Amount
    g.Status = temp.Status
    g.PaymentMethod = temp.PaymentMethod
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    } else {
        g.CreatedAt = nil
    }
    if temp.Items == nil {
        g.Items = nil
    } else {
        g.Items = temp.Items
    }
    g.Customer = temp.Customer
    g.Charge = temp.Charge
    g.Installments = temp.Installments
    g.BillingAddress = temp.BillingAddress
    g.Subscription = temp.Subscription
    g.Cycle = temp.Cycle
    g.Shipping = temp.Shipping
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    if temp.DueAt != nil {
        DueAtVal, err := time.Parse(time.RFC3339, *temp.DueAt)
        if err != nil {
            log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
        }
        g.DueAt = &DueAtVal
    } else {
        g.DueAt = nil
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
    if temp.BillingAt != nil {
        BillingAtVal, err := time.Parse(time.RFC3339, *temp.BillingAt)
        if err != nil {
            log.Fatalf("Cannot Parse billing_at as % s format.", time.RFC3339)
        }
        g.BillingAt = &BillingAtVal
    } else {
        g.BillingAt = nil
    }
    if temp.SeenAt != nil {
        SeenAtVal, err := time.Parse(time.RFC3339, *temp.SeenAt)
        if err != nil {
            log.Fatalf("Cannot Parse seen_at as % s format.", time.RFC3339)
        }
        g.SeenAt = &SeenAtVal
    } else {
        g.SeenAt = nil
    }
    g.TotalDiscount = temp.TotalDiscount
    g.TotalIncrement = temp.TotalIncrement
    g.SubscriptionId = temp.SubscriptionId
    return nil
}
