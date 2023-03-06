package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting a increment
type GetIncrementResponse struct {
    Id               *string                      `json:"id"`
    Value            *float64                     `json:"value"`
    IncrementType    *string                      `json:"increment_type"`
    Status           *string                      `json:"status"`
    CreatedAt        *time.Time                   `json:"created_at"`
    Cycles           *int                         `json:"cycles"`
    DeletedAt        *time.Time                   `json:"deleted_at"`
    Description      *string                      `json:"description"`
    Subscription     *GetSubscriptionResponse     `json:"subscription"`
    SubscriptionItem *GetSubscriptionItemResponse `json:"subscription_item"`
}

func (g *GetIncrementResponse) MarshalJSON() (
    []byte,
    error) {
    var CreatedAtVal *string
    if g.CreatedAt != nil {
        str := g.CreatedAt.Format(time.RFC3339)
        CreatedAtVal = &str
    } else {
        CreatedAtVal = nil
    }
    var DeletedAtVal *string
    if g.DeletedAt != nil {
        str := g.DeletedAt.Format(time.RFC3339)
        DeletedAtVal = &str
    } else {
        DeletedAtVal = nil
    }
    type marshallable GetIncrementResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"` 
        DeletedAt    *string `json:"deletedAt"` 
    }{
        CreatedAt:    CreatedAtVal,     
        DeletedAt:    DeletedAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetIncrementResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id               *string                      
        Value            *float64                     
        IncrementType    *string                      
        Status           *string                      
        CreatedAt        *string                      
        Cycles           *int                         
        DeletedAt        *string                      
        Description      *string                      
        Subscription     *GetSubscriptionResponse     
        SubscriptionItem *GetSubscriptionItemResponse 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Value = temp.Value
    g.IncrementType = temp.IncrementType
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
    g.Cycles = temp.Cycles
    if temp.DeletedAt != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, *temp.DeletedAt)
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt = &DeletedAtVal
    } else {
        g.DeletedAt = nil
    }
    g.Description = temp.Description
    g.Subscription = temp.Subscription
    g.SubscriptionItem = temp.SubscriptionItem
    return nil
}
