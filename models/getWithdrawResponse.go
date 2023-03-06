package models

import (
    "encoding/json"
    "log"
    "time"
)

type GetWithdrawResponse struct {
    Id                   *string                    `json:"id"`
    GatewayId            *string                    `json:"gateway_id"`
    Amount               *int                       `json:"amount"`
    Status               *string                    `json:"status"`
    CreatedAt            *time.Time                 `json:"created_at"`
    UpdatedAt            *time.Time                 `json:"updated_at"`
    Metadata             *[]string                  `json:"metadata"`
    Fee                  *int                       `json:"fee"`
    FundingDate          *time.Time                 `json:"funding_date"`
    FundingEstimatedDate *time.Time                 `json:"funding_estimated_date"`
    Type                 *string                    `json:"type"`
    Source               *GetWithdrawSourceResponse `json:"source"`
    Target               *GetWithdrawTargetResponse `json:"target"`
}

func (g *GetWithdrawResponse) MarshalJSON() (
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
    var FundingDateVal *string
    if g.FundingDate != nil {
        str := g.FundingDate.Format(time.RFC3339)
        FundingDateVal = &str
    } else {
        FundingDateVal = nil
    }
    var FundingEstimatedDateVal *string
    if g.FundingEstimatedDate != nil {
        str := g.FundingEstimatedDate.Format(time.RFC3339)
        FundingEstimatedDateVal = &str
    } else {
        FundingEstimatedDateVal = nil
    }
    type marshallable GetWithdrawResponse
    return json.Marshal(struct {
        marshallable         
        CreatedAt            *string `json:"createdAt"`            
        UpdatedAt            *string `json:"updatedAt"`            
        FundingDate          *string `json:"fundingDate"`          
        FundingEstimatedDate *string `json:"fundingEstimatedDate"` 
    }{
        CreatedAt:            CreatedAtVal,            
        UpdatedAt:            UpdatedAtVal,            
        FundingDate:          FundingDateVal,          
        FundingEstimatedDate: FundingEstimatedDateVal, 
        marshallable:         marshallable(*g),        
    })
}

func (g *GetWithdrawResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                   *string                    
        GatewayId            *string                    
        Amount               *int                       
        Status               *string                    
        CreatedAt            *string                    
        UpdatedAt            *string                    
        Metadata             *[]string                  
        Fee                  *int                       
        FundingDate          *string                    
        FundingEstimatedDate *string                    
        Type                 *string                    
        Source               *GetWithdrawSourceResponse 
        Target               *GetWithdrawTargetResponse 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.GatewayId = temp.GatewayId
    g.Amount = temp.Amount
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
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    g.Fee = temp.Fee
    if temp.FundingDate != nil {
        FundingDateVal, err := time.Parse(time.RFC3339, *temp.FundingDate)
        if err != nil {
            log.Fatalf("Cannot Parse funding_date as % s format.", time.RFC3339)
        }
        g.FundingDate = &FundingDateVal
    } else {
        g.FundingDate = nil
    }
    if temp.FundingEstimatedDate != nil {
        FundingEstimatedDateVal, err := time.Parse(time.RFC3339, *temp.FundingEstimatedDate)
        if err != nil {
            log.Fatalf("Cannot Parse funding_estimated_date as % s format.", time.RFC3339)
        }
        g.FundingEstimatedDate = &FundingEstimatedDateVal
    } else {
        g.FundingEstimatedDate = nil
    }
    g.Type = temp.Type
    g.Source = temp.Source
    g.Target = temp.Target
    return nil
}
