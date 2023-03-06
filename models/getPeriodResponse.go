package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting a period
type GetPeriodResponse struct {
    StartAt      *time.Time               `json:"start_at"`
    EndAt        *time.Time               `json:"end_at"`
    Id           *string                  `json:"id"`
    BillingAt    *time.Time               `json:"billing_at"`
    Subscription *GetSubscriptionResponse `json:"subscription"`
    Status       *string                  `json:"status"`
    Duration     *int                     `json:"duration"`
    CreatedAt    *string                  `json:"created_at"`
    UpdatedAt    *string                  `json:"updated_at"`
    Cycle        *int                     `json:"cycle"`
}

func (g *GetPeriodResponse) MarshalJSON() (
    []byte,
    error) {
    var StartAtVal *string
    if g.StartAt != nil {
        str := g.StartAt.Format(time.RFC3339)
        StartAtVal = &str
    } else {
        StartAtVal = nil
    }
    var EndAtVal *string
    if g.EndAt != nil {
        str := g.EndAt.Format(time.RFC3339)
        EndAtVal = &str
    } else {
        EndAtVal = nil
    }
    var BillingAtVal *string
    if g.BillingAt != nil {
        str := g.BillingAt.Format(time.RFC3339)
        BillingAtVal = &str
    } else {
        BillingAtVal = nil
    }
    type marshallable GetPeriodResponse
    return json.Marshal(struct {
        marshallable 
        StartAt      *string `json:"startAt"`   
        EndAt        *string `json:"endAt"`     
        BillingAt    *string `json:"billingAt"` 
    }{
        StartAt:      StartAtVal,       
        EndAt:        EndAtVal,         
        BillingAt:    BillingAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetPeriodResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        StartAt      *string                  
        EndAt        *string                  
        Id           *string                  
        BillingAt    *string                  
        Subscription *GetSubscriptionResponse 
        Status       *string                  
        Duration     *int                     
        CreatedAt    *string                  
        UpdatedAt    *string                  
        Cycle        *int                     
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    if temp.StartAt != nil {
        StartAtVal, err := time.Parse(time.RFC3339, *temp.StartAt)
        if err != nil {
            log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
        }
        g.StartAt = &StartAtVal
    } else {
        g.StartAt = nil
    }
    if temp.EndAt != nil {
        EndAtVal, err := time.Parse(time.RFC3339, *temp.EndAt)
        if err != nil {
            log.Fatalf("Cannot Parse end_at as % s format.", time.RFC3339)
        }
        g.EndAt = &EndAtVal
    } else {
        g.EndAt = nil
    }
    g.Id = temp.Id
    if temp.BillingAt != nil {
        BillingAtVal, err := time.Parse(time.RFC3339, *temp.BillingAt)
        if err != nil {
            log.Fatalf("Cannot Parse billing_at as % s format.", time.RFC3339)
        }
        g.BillingAt = &BillingAtVal
    } else {
        g.BillingAt = nil
    }
    g.Subscription = temp.Subscription
    g.Status = temp.Status
    g.Duration = temp.Duration
    g.CreatedAt = temp.CreatedAt
    g.UpdatedAt = temp.UpdatedAt
    g.Cycle = temp.Cycle
    return nil
}
