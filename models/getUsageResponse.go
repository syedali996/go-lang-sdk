package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting a usage
type GetUsageResponse struct {
    Id               *string                      `json:"id"`
    Quantity         *int                         `json:"quantity"`
    Description      *string                      `json:"description"`
    UsedAt           *time.Time                   `json:"used_at"`
    CreatedAt        *time.Time                   `json:"created_at"`
    Status           *string                      `json:"status"`
    DeletedAt        *time.Time                   `json:"deleted_at"`
    SubscriptionItem *GetSubscriptionItemResponse `json:"subscription_item"`
    Code             *string                      `json:"code"`
    Group            *string                      `json:"group"`
    Amount           *int                         `json:"amount"`
}

func (g *GetUsageResponse) MarshalJSON() (
    []byte,
    error) {
    var UsedAtVal *string
    if g.UsedAt != nil {
        str := g.UsedAt.Format(time.RFC3339)
        UsedAtVal = &str
    } else {
        UsedAtVal = nil
    }
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
    type marshallable GetUsageResponse
    return json.Marshal(struct {
        marshallable 
        UsedAt       *string `json:"usedAt"`    
        CreatedAt    *string `json:"createdAt"` 
        DeletedAt    *string `json:"deletedAt"` 
    }{
        UsedAt:       UsedAtVal,        
        CreatedAt:    CreatedAtVal,     
        DeletedAt:    DeletedAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetUsageResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id               *string                      
        Quantity         *int                         
        Description      *string                      
        UsedAt           *string                      
        CreatedAt        *string                      
        Status           *string                      
        DeletedAt        *string                      
        SubscriptionItem *GetSubscriptionItemResponse 
        Code             *string                      
        Group            *string                      
        Amount           *int                         
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Quantity = temp.Quantity
    g.Description = temp.Description
    if temp.UsedAt != nil {
        UsedAtVal, err := time.Parse(time.RFC3339, *temp.UsedAt)
        if err != nil {
            log.Fatalf("Cannot Parse used_at as % s format.", time.RFC3339)
        }
        g.UsedAt = &UsedAtVal
    } else {
        g.UsedAt = nil
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
    g.Status = temp.Status
    if temp.DeletedAt != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, *temp.DeletedAt)
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt = &DeletedAtVal
    } else {
        g.DeletedAt = nil
    }
    g.SubscriptionItem = temp.SubscriptionItem
    g.Code = temp.Code
    g.Group = temp.Group
    g.Amount = temp.Amount
    return nil
}
