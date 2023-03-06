package models

import (
    "encoding/json"
    "log"
    "time"
)

// Request for creating a new charge
type CreateChargeRequest struct {
    Code       string                 `json:"code"`
    Amount     int                    `json:"amount"`
    CustomerId string                 `json:"customer_id"`
    Customer   CreateCustomerRequest  `json:"customer"`
    Payment    CreatePaymentRequest   `json:"payment"`
    Metadata   map[string]string      `json:"metadata"`
    DueAt      time.Time              `json:"due_at,omitempty"`
    Antifraud  CreateAntifraudRequest `json:"antifraud"`
    OrderId    string                 `json:"order_id"`
}

func (c *CreateChargeRequest) MarshalJSON() (
    []byte,
    error) {
    type marshallable CreateChargeRequest
    return json.Marshal(struct {
        marshallable 
        DueAt        string `json:"dueAt"` 
    }{
        DueAt:        c.DueAt.Format(time.RFC3339), 
        marshallable: marshallable(*c),             
    })
}

func (c *CreateChargeRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Code       string                 
        Amount     int                    
        CustomerId string                 
        Customer   CreateCustomerRequest  
        Payment    CreatePaymentRequest   
        Metadata   map[string]string      
        Antifraud  CreateAntifraudRequest 
        OrderId    string                 
        DueAt      string                 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    c.Code = temp.Code
    c.Amount = temp.Amount
    c.CustomerId = temp.CustomerId
    c.Customer = temp.Customer
    c.Payment = temp.Payment
    c.Metadata = temp.Metadata
    c.Antifraud = temp.Antifraud
    c.OrderId = temp.OrderId
    c.DueAt, err = time.Parse(time.RFC3339, temp.DueAt)
    if err != nil {
        log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
    }
    return nil
}
