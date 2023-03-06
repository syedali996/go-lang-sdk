package models

import (
    "encoding/json"
    "log"
    "time"
)

type GetSubscriptionItemResponse struct {
    Id            *string                   `json:"id"`
    Description   *string                   `json:"description"`
    Status        *string                   `json:"status"`
    CreatedAt     *time.Time                `json:"created_at"`
    UpdatedAt     *time.Time                `json:"updated_at"`
    PricingScheme *GetPricingSchemeResponse `json:"pricing_scheme"`
    Discounts     *[]GetDiscountResponse    `json:"discounts"`
    Increments    *[]GetIncrementResponse   `json:"increments"`
    Subscription  *GetSubscriptionResponse  `json:"subscription"`
    Name          *string                   `json:"name"`
    Quantity      *int                      `json:"quantity"`
    Cycles        *int                      `json:"cycles"`
    DeletedAt     *time.Time                `json:"deleted_at"`
}

func (g *GetSubscriptionItemResponse) MarshalJSON() (
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
    type marshallable GetSubscriptionItemResponse
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

func (g *GetSubscriptionItemResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id            *string                   
        Description   *string                   
        Status        *string                   
        CreatedAt     *string                   
        UpdatedAt     *string                   
        PricingScheme *GetPricingSchemeResponse 
        Discounts     *[]GetDiscountResponse    
        Increments    *[]GetIncrementResponse   
        Subscription  *GetSubscriptionResponse  
        Name          *string                   
        Quantity      *int                      
        Cycles        *int                      
        DeletedAt     *string                   
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Description = temp.Description
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
    g.PricingScheme = temp.PricingScheme
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
    g.Subscription = temp.Subscription
    g.Name = temp.Name
    g.Quantity = temp.Quantity
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
    return nil
}
