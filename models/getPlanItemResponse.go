package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting a plan item
type GetPlanItemResponse struct {
    Id            *string                   `json:"id"`
    Name          *string                   `json:"name"`
    Status        *string                   `json:"status"`
    CreatedAt     *time.Time                `json:"created_at"`
    UpdatedAt     *time.Time                `json:"updated_at"`
    PricingScheme *GetPricingSchemeResponse `json:"pricing_scheme"`
    Description   *string                   `json:"description"`
    Plan          *GetPlanResponse          `json:"plan"`
    Quantity      *int                      `json:"quantity"`
    Cycles        *int                      `json:"cycles"`
    DeletedAt     *time.Time                `json:"deleted_at"`
}

func (g *GetPlanItemResponse) MarshalJSON() (
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
    type marshallable GetPlanItemResponse
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

func (g *GetPlanItemResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id            *string                   
        Name          *string                   
        Status        *string                   
        CreatedAt     *string                   
        UpdatedAt     *string                   
        PricingScheme *GetPricingSchemeResponse 
        Description   *string                   
        Plan          *GetPlanResponse          
        Quantity      *int                      
        Cycles        *int                      
        DeletedAt     *string                   
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Name = temp.Name
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
    g.Description = temp.Description
    g.Plan = temp.Plan
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
