package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting a plan
type GetPlanResponse struct {
    Id                  *string                `json:"id"`
    Name                *string                `json:"name"`
    Description         *string                `json:"description"`
    Url                 *string                `json:"url"`
    StatementDescriptor *string                `json:"statement_descriptor"`
    Interval            *string                `json:"interval"`
    IntervalCount       *int                   `json:"interval_count"`
    BillingType         *string                `json:"billing_type"`
    PaymentMethods      *[]string              `json:"payment_methods"`
    Installments        *[]int                 `json:"installments"`
    Status              *string                `json:"status"`
    Currency            *string                `json:"currency"`
    CreatedAt           *time.Time             `json:"created_at"`
    UpdatedAt           *time.Time             `json:"updated_at"`
    Items               *[]GetPlanItemResponse `json:"items"`
    BillingDays         *[]int                 `json:"billing_days"`
    Shippable           *bool                  `json:"shippable"`
    Metadata            map[string]*string     `json:"metadata"`
    TrialPeriodDays     *int                   `json:"trial_period_days"`
    MinimumPrice        *int                   `json:"minimum_price"`
    DeletedAt           *time.Time             `json:"deleted_at"`
}

func (g *GetPlanResponse) MarshalJSON() (
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
    type marshallable GetPlanResponse
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

func (g *GetPlanResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                  *string                
        Name                *string                
        Description         *string                
        Url                 *string                
        StatementDescriptor *string                
        Interval            *string                
        IntervalCount       *int                   
        BillingType         *string                
        PaymentMethods      *[]string              
        Installments        *[]int                 
        Status              *string                
        Currency            *string                
        CreatedAt           *string                
        UpdatedAt           *string                
        Items               *[]GetPlanItemResponse 
        BillingDays         *[]int                 
        Shippable           *bool                  
        Metadata            map[string]*string     
        TrialPeriodDays     *int                   
        MinimumPrice        *int                   
        DeletedAt           *string                
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Name = temp.Name
    g.Description = temp.Description
    g.Url = temp.Url
    g.StatementDescriptor = temp.StatementDescriptor
    g.Interval = temp.Interval
    g.IntervalCount = temp.IntervalCount
    g.BillingType = temp.BillingType
    if temp.PaymentMethods == nil {
        g.PaymentMethods = nil
    } else {
        g.PaymentMethods = temp.PaymentMethods
    }
    if temp.Installments == nil {
        g.Installments = nil
    } else {
        g.Installments = temp.Installments
    }
    g.Status = temp.Status
    g.Currency = temp.Currency
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
    if temp.Items == nil {
        g.Items = nil
    } else {
        g.Items = temp.Items
    }
    if temp.BillingDays == nil {
        g.BillingDays = nil
    } else {
        g.BillingDays = temp.BillingDays
    }
    g.Shippable = temp.Shippable
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    g.TrialPeriodDays = temp.TrialPeriodDays
    g.MinimumPrice = temp.MinimumPrice
    if temp.DeletedAt != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, *temp.DeletedAt)
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt = &DeletedAtVal
    } else {
        g.DeletedAt = nil
    }
    return nil
}
