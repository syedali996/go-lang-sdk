package models

import (
    "encoding/json"
    "log"
    "time"
)

type GetCheckoutBoletoPaymentResponse struct {
    DueAt        *time.Time `json:"due_at"`
    Instructions *string    `json:"instructions"`
}

func (g *GetCheckoutBoletoPaymentResponse) MarshalJSON() (
    []byte,
    error) {
    var DueAtVal *string
    if g.DueAt != nil {
        str := g.DueAt.Format(time.RFC3339)
        DueAtVal = &str
    } else {
        DueAtVal = nil
    }
    type marshallable GetCheckoutBoletoPaymentResponse
    return json.Marshal(struct {
        marshallable 
        DueAt        *string `json:"dueAt"` 
    }{
        DueAt:        DueAtVal,         
        marshallable: marshallable(*g), 
    })
}

func (g *GetCheckoutBoletoPaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        DueAt        *string 
        Instructions *string 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
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
    g.Instructions = temp.Instructions
    return nil
}
