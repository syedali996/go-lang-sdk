package models

import (
    "encoding/json"
    "log"
    "time"
)

// Contains the settings for creating a boleto payment
type CreateBoletoPaymentRequest struct {
    Retries             int                   `json:"retries"`
    Bank                string                `json:"bank"`
    Instructions        string                `json:"instructions"`
    DueAt               time.Time             `json:"due_at,omitempty"`
    BillingAddress      CreateAddressRequest  `json:"billing_address"`
    BillingAddressId    string                `json:"billing_address_id"`
    NossoNumero         string                `json:"nosso_numero,omitempty"`
    DocumentNumber      string                `json:"document_number"`
    StatementDescriptor string                `json:"statement_descriptor"`
    Interest            CreateInterestRequest `json:"interest,omitempty"`
    Fine                CreateFineRequest     `json:"fine,omitempty"`
    MaxDaysToPayPastDue int                   `json:"max_days_to_pay_past_due,omitempty"`
}

func (c *CreateBoletoPaymentRequest) MarshalJSON() (
    []byte,
    error) {
    type marshallable CreateBoletoPaymentRequest
    return json.Marshal(struct {
        marshallable 
        DueAt        string `json:"dueAt"` 
    }{
        DueAt:        c.DueAt.Format(time.RFC3339), 
        marshallable: marshallable(*c),             
    })
}

func (c *CreateBoletoPaymentRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Retries             int                   
        Bank                string                
        Instructions        string                
        BillingAddress      CreateAddressRequest  
        BillingAddressId    string                
        DocumentNumber      string                
        StatementDescriptor string                
        DueAt               string                
        NossoNumero         string                
        Interest            CreateInterestRequest 
        Fine                CreateFineRequest     
        MaxDaysToPayPastDue int                   
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    c.Retries = temp.Retries
    c.Bank = temp.Bank
    c.Instructions = temp.Instructions
    c.BillingAddress = temp.BillingAddress
    c.BillingAddressId = temp.BillingAddressId
    c.DocumentNumber = temp.DocumentNumber
    c.StatementDescriptor = temp.StatementDescriptor
    c.DueAt, err = time.Parse(time.RFC3339, temp.DueAt)
    if err != nil {
        log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
    }
    c.NossoNumero = temp.NossoNumero
    c.Interest = temp.Interest
    c.Fine = temp.Fine
    c.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
    return nil
}
