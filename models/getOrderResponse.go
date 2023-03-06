package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting an Order
type GetOrderResponse struct {
    Id         *string                       `json:"id"`
    Code       *string                       `json:"code"`
    Currency   *string                       `json:"currency"`
    Items      *[]GetOrderItemResponse       `json:"items"`
    Customer   *GetCustomerResponse          `json:"customer"`
    Status     *string                       `json:"status"`
    CreatedAt  *time.Time                    `json:"created_at"`
    UpdatedAt  *time.Time                    `json:"updated_at"`
    Charges    *[]GetChargeResponse          `json:"charges"`
    InvoiceUrl *string                       `json:"invoice_url"`
    Shipping   *GetShippingResponse          `json:"shipping"`
    Metadata   map[string]*string            `json:"metadata"`
    Checkouts  *[]GetCheckoutPaymentResponse `json:"checkouts"`
    Ip         *string                       `json:"ip"`
    SessionId  *string                       `json:"session_id"`
    Location   *GetLocationResponse          `json:"location"`
    Device     *GetDeviceResponse            `json:"device"`
    Closed     *bool                         `json:"closed"`
}

func (g *GetOrderResponse) MarshalJSON() (
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
    type marshallable GetOrderResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"` 
        UpdatedAt    *string `json:"updatedAt"` 
    }{
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetOrderResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id         *string                       
        Code       *string                       
        Currency   *string                       
        Items      *[]GetOrderItemResponse       
        Customer   *GetCustomerResponse          
        Status     *string                       
        CreatedAt  *string                       
        UpdatedAt  *string                       
        Charges    *[]GetChargeResponse          
        InvoiceUrl *string                       
        Shipping   *GetShippingResponse          
        Metadata   map[string]*string            
        Checkouts  *[]GetCheckoutPaymentResponse 
        Ip         *string                       
        SessionId  *string                       
        Location   *GetLocationResponse          
        Device     *GetDeviceResponse            
        Closed     *bool                         
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Code = temp.Code
    g.Currency = temp.Currency
    if temp.Items == nil {
        g.Items = nil
    } else {
        g.Items = temp.Items
    }
    g.Customer = temp.Customer
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
    if temp.Charges == nil {
        g.Charges = nil
    } else {
        g.Charges = temp.Charges
    }
    g.InvoiceUrl = temp.InvoiceUrl
    g.Shipping = temp.Shipping
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    if temp.Checkouts == nil {
        g.Checkouts = nil
    } else {
        g.Checkouts = temp.Checkouts
    }
    g.Ip = temp.Ip
    g.SessionId = temp.SessionId
    g.Location = temp.Location
    g.Device = temp.Device
    g.Closed = temp.Closed
    return nil
}
