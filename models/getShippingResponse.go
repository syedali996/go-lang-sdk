package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting the shipping data
type GetShippingResponse struct {
    Amount                *int                `json:"amount"`
    Description           *string             `json:"description"`
    RecipientName         *string             `json:"recipient_name"`
    RecipientPhone        *string             `json:"recipient_phone"`
    Address               *GetAddressResponse `json:"address"`
    MaxDeliveryDate       *time.Time          `json:"max_delivery_date"`
    EstimatedDeliveryDate *time.Time          `json:"estimated_delivery_date"`
    Type                  *string             `json:"type"`
}

func (g *GetShippingResponse) MarshalJSON() (
    []byte,
    error) {
    var MaxDeliveryDateVal *string
    if g.MaxDeliveryDate != nil {
        str := g.MaxDeliveryDate.Format(time.RFC3339)
        MaxDeliveryDateVal = &str
    } else {
        MaxDeliveryDateVal = nil
    }
    var EstimatedDeliveryDateVal *string
    if g.EstimatedDeliveryDate != nil {
        str := g.EstimatedDeliveryDate.Format(time.RFC3339)
        EstimatedDeliveryDateVal = &str
    } else {
        EstimatedDeliveryDateVal = nil
    }
    type marshallable GetShippingResponse
    return json.Marshal(struct {
        marshallable          
        MaxDeliveryDate       *string `json:"maxDeliveryDate"`       
        EstimatedDeliveryDate *string `json:"estimatedDeliveryDate"` 
    }{
        MaxDeliveryDate:       MaxDeliveryDateVal,       
        EstimatedDeliveryDate: EstimatedDeliveryDateVal, 
        marshallable:          marshallable(*g),         
    })
}

func (g *GetShippingResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Amount                *int                
        Description           *string             
        RecipientName         *string             
        RecipientPhone        *string             
        Address               *GetAddressResponse 
        MaxDeliveryDate       *string             
        EstimatedDeliveryDate *string             
        Type                  *string             
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Amount = temp.Amount
    g.Description = temp.Description
    g.RecipientName = temp.RecipientName
    g.RecipientPhone = temp.RecipientPhone
    g.Address = temp.Address
    if temp.MaxDeliveryDate != nil {
        MaxDeliveryDateVal, err := time.Parse(time.RFC3339, *temp.MaxDeliveryDate)
        if err != nil {
            log.Fatalf("Cannot Parse max_delivery_date as % s format.", time.RFC3339)
        }
        g.MaxDeliveryDate = &MaxDeliveryDateVal
    } else {
        g.MaxDeliveryDate = nil
    }
    if temp.EstimatedDeliveryDate != nil {
        EstimatedDeliveryDateVal, err := time.Parse(time.RFC3339, *temp.EstimatedDeliveryDate)
        if err != nil {
            log.Fatalf("Cannot Parse estimated_delivery_date as % s format.", time.RFC3339)
        }
        g.EstimatedDeliveryDate = &EstimatedDeliveryDateVal
    } else {
        g.EstimatedDeliveryDate = nil
    }
    g.Type = temp.Type
    return nil
}
