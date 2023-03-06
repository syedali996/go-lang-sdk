package models

import (
    "encoding/json"
    "log"
    "time"
)

// Request for creating a subcription
type CreateSubscriptionRequest struct {
    Customer             CreateCustomerRequest           `json:"customer"`
    Card                 CreateCardRequest               `json:"card"`
    Code                 string                          `json:"code"`
    PaymentMethod        string                          `json:"payment_method"`
    BillingType          string                          `json:"billing_type"`
    StatementDescriptor  string                          `json:"statement_descriptor"`
    Description          string                          `json:"description"`
    Currency             string                          `json:"currency"`
    Interval             string                          `json:"interval"`
    IntervalCount        int                             `json:"interval_count"`
    PricingScheme        CreatePricingSchemeRequest      `json:"pricing_scheme"`
    Items                []CreateSubscriptionItemRequest `json:"items"`
    Shipping             CreateShippingRequest           `json:"shipping"`
    Discounts            []CreateDiscountRequest         `json:"discounts"`
    Metadata             map[string]string               `json:"metadata"`
    Setup                CreateSetupRequest              `json:"setup,omitempty"`
    PlanId               string                          `json:"plan_id,omitempty"`
    CustomerId           string                          `json:"customer_id,omitempty"`
    CardId               string                          `json:"card_id,omitempty"`
    BillingDay           int                             `json:"billing_day,omitempty"`
    Installments         int                             `json:"installments,omitempty"`
    StartAt              time.Time                       `json:"start_at,omitempty"`
    MinimumPrice         int                             `json:"minimum_price,omitempty"`
    Cycles               int                             `json:"cycles,omitempty"`
    CardToken            string                          `json:"card_token,omitempty"`
    GatewayAffiliationId string                          `json:"gateway_affiliation_id,omitempty"`
    Quantity             int                             `json:"quantity,omitempty"`
    BoletoDueDays        int                             `json:"boleto_due_days,omitempty"`
    Increments           []CreateIncrementRequest        `json:"increments"`
    Period               CreatePeriodRequest             `json:"period,omitempty"`
    Submerchant          CreateSubMerchantRequest        `json:"submerchant,omitempty"`
    Split                CreateSubscriptionSplitRequest  `json:"split,omitempty"`
    Boleto               CreateSubscriptionBoletoRequest `json:"boleto,omitempty"`
}

func (c *CreateSubscriptionRequest) MarshalJSON() (
    []byte,
    error) {
    type marshallable CreateSubscriptionRequest
    return json.Marshal(struct {
        marshallable 
        StartAt      string `json:"startAt"` 
    }{
        StartAt:      c.StartAt.Format(time.RFC3339), 
        marshallable: marshallable(*c),               
    })
}

func (c *CreateSubscriptionRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Customer             CreateCustomerRequest           
        Card                 CreateCardRequest               
        Code                 string                          
        PaymentMethod        string                          
        BillingType          string                          
        StatementDescriptor  string                          
        Description          string                          
        Currency             string                          
        Interval             string                          
        IntervalCount        int                             
        PricingScheme        CreatePricingSchemeRequest      
        Items                []CreateSubscriptionItemRequest 
        Shipping             CreateShippingRequest           
        Discounts            []CreateDiscountRequest         
        Metadata             map[string]string               
        Increments           []CreateIncrementRequest        
        Setup                CreateSetupRequest              
        PlanId               string                          
        CustomerId           string                          
        CardId               string                          
        BillingDay           int                             
        Installments         int                             
        StartAt              string                          
        MinimumPrice         int                             
        Cycles               int                             
        CardToken            string                          
        GatewayAffiliationId string                          
        Quantity             int                             
        BoletoDueDays        int                             
        Period               CreatePeriodRequest             
        Submerchant          CreateSubMerchantRequest        
        Split                CreateSubscriptionSplitRequest  
        Boleto               CreateSubscriptionBoletoRequest 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    c.Customer = temp.Customer
    c.Card = temp.Card
    c.Code = temp.Code
    c.PaymentMethod = temp.PaymentMethod
    c.BillingType = temp.BillingType
    c.StatementDescriptor = temp.StatementDescriptor
    c.Description = temp.Description
    c.Currency = temp.Currency
    c.Interval = temp.Interval
    c.IntervalCount = temp.IntervalCount
    c.PricingScheme = temp.PricingScheme
    c.Items = temp.Items
    c.Shipping = temp.Shipping
    c.Discounts = temp.Discounts
    c.Metadata = temp.Metadata
    c.Increments = temp.Increments
    c.Setup = temp.Setup
    c.PlanId = temp.PlanId
    c.CustomerId = temp.CustomerId
    c.CardId = temp.CardId
    c.BillingDay = temp.BillingDay
    c.Installments = temp.Installments
    c.StartAt, err = time.Parse(time.RFC3339, temp.StartAt)
    if err != nil {
        log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
    }
    c.MinimumPrice = temp.MinimumPrice
    c.Cycles = temp.Cycles
    c.CardToken = temp.CardToken
    c.GatewayAffiliationId = temp.GatewayAffiliationId
    c.Quantity = temp.Quantity
    c.BoletoDueDays = temp.BoletoDueDays
    c.Period = temp.Period
    c.Submerchant = temp.Submerchant
    c.Split = temp.Split
    c.Boleto = temp.Boleto
    return nil
}
