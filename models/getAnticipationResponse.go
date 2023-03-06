package models

import (
    "encoding/json"
    "log"
    "time"
)

// Anticipation
type GetAnticipationResponse struct {
    Id              *string               `json:"id"`
    RequestedAmount *int                  `json:"requested_amount"`
    ApprovedAmount  *int                  `json:"approved_amount"`
    Recipient       *GetRecipientResponse `json:"recipient"`
    Pgid            *string               `json:"pgid"`
    CreatedAt       *time.Time            `json:"created_at"`
    UpdatedAt       *time.Time            `json:"updated_at"`
    PaymentDate     *time.Time            `json:"payment_date"`
    Status          *string               `json:"status"`
    Timeframe       *string               `json:"timeframe"`
}

func (g *GetAnticipationResponse) MarshalJSON() (
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
    var PaymentDateVal *string
    if g.PaymentDate != nil {
        str := g.PaymentDate.Format(time.RFC3339)
        PaymentDateVal = &str
    } else {
        PaymentDateVal = nil
    }
    type marshallable GetAnticipationResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"`   
        UpdatedAt    *string `json:"updatedAt"`   
        PaymentDate  *string `json:"paymentDate"` 
    }{
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        PaymentDate:  PaymentDateVal,   
        marshallable: marshallable(*g), 
    })
}

func (g *GetAnticipationResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id              *string               
        RequestedAmount *int                  
        ApprovedAmount  *int                  
        Recipient       *GetRecipientResponse 
        Pgid            *string               
        CreatedAt       *string               
        UpdatedAt       *string               
        PaymentDate     *string               
        Status          *string               
        Timeframe       *string               
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.RequestedAmount = temp.RequestedAmount
    g.ApprovedAmount = temp.ApprovedAmount
    g.Recipient = temp.Recipient
    g.Pgid = temp.Pgid
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
    if temp.PaymentDate != nil {
        PaymentDateVal, err := time.Parse(time.RFC3339, *temp.PaymentDate)
        if err != nil {
            log.Fatalf("Cannot Parse payment_date as % s format.", time.RFC3339)
        }
        g.PaymentDate = &PaymentDateVal
    } else {
        g.PaymentDate = nil
    }
    g.Status = temp.Status
    g.Timeframe = temp.Timeframe
    return nil
}
