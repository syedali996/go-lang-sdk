package models

import (
    "encoding/json"
    "log"
    "time"
)

// Shipping data
type CreateShippingRequest struct {
    Amount                int                  `json:"amount"`
    Description           string               `json:"description"`
    RecipientName         string               `json:"recipient_name"`
    RecipientPhone        string               `json:"recipient_phone"`
    AddressId             string               `json:"address_id"`
    Address               CreateAddressRequest `json:"address"`
    MaxDeliveryDate       time.Time            `json:"max_delivery_date,omitempty"`
    EstimatedDeliveryDate time.Time            `json:"estimated_delivery_date,omitempty"`
    Type                  string               `json:"type"`
}

func (c *CreateShippingRequest) MarshalJSON() (
    []byte,
    error) {
    type marshallable CreateShippingRequest
    return json.Marshal(struct {
        marshallable          
        MaxDeliveryDate       string `json:"maxDeliveryDate"`       
        EstimatedDeliveryDate string `json:"estimatedDeliveryDate"` 
    }{
        MaxDeliveryDate:       c.MaxDeliveryDate.Format(time.RFC3339),       
        EstimatedDeliveryDate: c.EstimatedDeliveryDate.Format(time.RFC3339), 
        marshallable:          marshallable(*c),                             
    })
}

func (c *CreateShippingRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Amount                int                  
        Description           string               
        RecipientName         string               
        RecipientPhone        string               
        AddressId             string               
        Address               CreateAddressRequest 
        Type                  string               
        MaxDeliveryDate       string               
        EstimatedDeliveryDate string               
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    c.Amount = temp.Amount
    c.Description = temp.Description
    c.RecipientName = temp.RecipientName
    c.RecipientPhone = temp.RecipientPhone
    c.AddressId = temp.AddressId
    c.Address = temp.Address
    c.Type = temp.Type
    c.MaxDeliveryDate, err = time.Parse(time.RFC3339, temp.MaxDeliveryDate)
    if err != nil {
        log.Fatalf("Cannot Parse max_delivery_date as % s format.", time.RFC3339)
    }
    c.EstimatedDeliveryDate, err = time.Parse(time.RFC3339, temp.EstimatedDeliveryDate)
    if err != nil {
        log.Fatalf("Cannot Parse estimated_delivery_date as % s format.", time.RFC3339)
    }
    return nil
}
