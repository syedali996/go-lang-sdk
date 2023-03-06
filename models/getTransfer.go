package models

import (
    "encoding/json"
    "log"
    "time"
)

type GetTransfer struct {
    Id                   string                    `json:"id"`
    GatewayId            string                    `json:"gateway_id"`
    Amount               int                       `json:"amount"`
    Status               string                    `json:"status"`
    CreatedAt            time.Time                 `json:"created_at"`
    UpdatedAt            time.Time                 `json:"updated_at"`
    Metadata             map[string]string         `json:"metadata,omitempty"`
    Fee                  int                       `json:"fee,omitempty"`
    FundingDate          time.Time                 `json:"funding_date,omitempty"`
    FundingEstimatedDate time.Time                 `json:"funding_estimated_date,omitempty"`
    Type                 string                    `json:"type"`
    Source               GetTransferSourceResponse `json:"source"`
    Target               GetTransferTargetResponse `json:"target"`
}

func (g *GetTransfer) MarshalJSON() (
    []byte,
    error) {
    type marshallable GetTransfer
    return json.Marshal(struct {
        marshallable         
        CreatedAt            string `json:"createdAt"`            
        UpdatedAt            string `json:"updatedAt"`            
        FundingDate          string `json:"fundingDate"`          
        FundingEstimatedDate string `json:"fundingEstimatedDate"` 
    }{
        CreatedAt:            g.CreatedAt.Format(time.RFC3339),            
        UpdatedAt:            g.UpdatedAt.Format(time.RFC3339),            
        FundingDate:          g.FundingDate.Format(time.RFC3339),          
        FundingEstimatedDate: g.FundingEstimatedDate.Format(time.RFC3339), 
        marshallable:         marshallable(*g),                            
    })
}

func (g *GetTransfer) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                   string                    
        GatewayId            string                    
        Amount               int                       
        Status               string                    
        CreatedAt            string                    
        UpdatedAt            string                    
        Type                 string                    
        Source               GetTransferSourceResponse 
        Target               GetTransferTargetResponse 
        Metadata             map[string]string         
        Fee                  int                       
        FundingDate          string                    
        FundingEstimatedDate string                    
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.GatewayId = temp.GatewayId
    g.Amount = temp.Amount
    g.Status = temp.Status
    g.CreatedAt, err = time.Parse(time.RFC3339, temp.CreatedAt)
    if err != nil {
        log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
    }
    g.UpdatedAt, err = time.Parse(time.RFC3339, temp.UpdatedAt)
    if err != nil {
        log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
    }
    g.Type = temp.Type
    g.Source = temp.Source
    g.Target = temp.Target
    g.Metadata = temp.Metadata
    g.Fee = temp.Fee
    g.FundingDate, err = time.Parse(time.RFC3339, temp.FundingDate)
    if err != nil {
        log.Fatalf("Cannot Parse funding_date as % s format.", time.RFC3339)
    }
    g.FundingEstimatedDate, err = time.Parse(time.RFC3339, temp.FundingEstimatedDate)
    if err != nil {
        log.Fatalf("Cannot Parse funding_estimated_date as % s format.", time.RFC3339)
    }
    return nil
}
