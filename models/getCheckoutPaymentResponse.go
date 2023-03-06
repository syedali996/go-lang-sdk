package models

import (
    "encoding/json"
    "log"
    "time"
)

// Resposta das configurações de pagamento do checkout
type GetCheckoutPaymentResponse struct {
    Id                      *string                                 `json:"id"`
    Amount                  *int                                    `json:"amount"`
    DefaultPaymentMethod    *string                                 `json:"default_payment_method"`
    SuccessUrl              *string                                 `json:"success_url"`
    PaymentUrl              *string                                 `json:"payment_url"`
    GatewayAffiliationId    *string                                 `json:"gateway_affiliation_id"`
    AcceptedPaymentMethods  *[]string                               `json:"accepted_payment_methods"`
    Status                  *string                                 `json:"status"`
    SkipCheckoutSuccessPage *bool                                   `json:"skip_checkout_success_page"`
    CreatedAt               *time.Time                              `json:"created_at"`
    UpdatedAt               *time.Time                              `json:"updated_at"`
    CanceledAt              *time.Time                              `json:"canceled_at"`
    CustomerEditable        *bool                                   `json:"customer_editable"`
    Customer                *GetCustomerResponse                    `json:"customer"`
    Billingaddress          *GetAddressResponse                     `json:"billingaddress"`
    CreditCard              *GetCheckoutCreditCardPaymentResponse   `json:"credit_card"`
    Boleto                  *GetCheckoutBoletoPaymentResponse       `json:"boleto"`
    BillingAddressEditable  *bool                                   `json:"billing_address_editable"`
    Shipping                *GetShippingResponse                    `json:"shipping"`
    Shippable               *bool                                   `json:"shippable"`
    ClosedAt                *time.Time                              `json:"closed_at"`
    ExpiresAt               *time.Time                              `json:"expires_at"`
    Currency                *string                                 `json:"currency"`
    DebitCard               *GetCheckoutDebitCardPaymentResponse    `json:"debit_card"`
    BankTransfer            *GetCheckoutBankTransferPaymentResponse `json:"bank_transfer"`
    AcceptedBrands          *[]string                               `json:"accepted_brands"`
    Pix                     *GetCheckoutPixPaymentResponse          `json:"pix"`
}

func (g *GetCheckoutPaymentResponse) MarshalJSON() (
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
    var CanceledAtVal *string
    if g.CanceledAt != nil {
        str := g.CanceledAt.Format(time.RFC3339)
        CanceledAtVal = &str
    } else {
        CanceledAtVal = nil
    }
    var ClosedAtVal *string
    if g.ClosedAt != nil {
        str := g.ClosedAt.Format(time.RFC3339)
        ClosedAtVal = &str
    } else {
        ClosedAtVal = nil
    }
    var ExpiresAtVal *string
    if g.ExpiresAt != nil {
        str := g.ExpiresAt.Format(time.RFC3339)
        ExpiresAtVal = &str
    } else {
        ExpiresAtVal = nil
    }
    type marshallable GetCheckoutPaymentResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"`  
        UpdatedAt    *string `json:"updatedAt"`  
        CanceledAt   *string `json:"canceledAt"` 
        ClosedAt     *string `json:"closedAt"`   
        ExpiresAt    *string `json:"expiresAt"`  
    }{
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        CanceledAt:   CanceledAtVal,    
        ClosedAt:     ClosedAtVal,      
        ExpiresAt:    ExpiresAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetCheckoutPaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                      *string                                 
        Amount                  *int                                    
        DefaultPaymentMethod    *string                                 
        SuccessUrl              *string                                 
        PaymentUrl              *string                                 
        GatewayAffiliationId    *string                                 
        AcceptedPaymentMethods  *[]string                               
        Status                  *string                                 
        SkipCheckoutSuccessPage *bool                                   
        CreatedAt               *string                                 
        UpdatedAt               *string                                 
        CanceledAt              *string                                 
        CustomerEditable        *bool                                   
        Customer                *GetCustomerResponse                    
        Billingaddress          *GetAddressResponse                     
        CreditCard              *GetCheckoutCreditCardPaymentResponse   
        Boleto                  *GetCheckoutBoletoPaymentResponse       
        BillingAddressEditable  *bool                                   
        Shipping                *GetShippingResponse                    
        Shippable               *bool                                   
        ClosedAt                *string                                 
        ExpiresAt               *string                                 
        Currency                *string                                 
        DebitCard               *GetCheckoutDebitCardPaymentResponse    
        BankTransfer            *GetCheckoutBankTransferPaymentResponse 
        AcceptedBrands          *[]string                               
        Pix                     *GetCheckoutPixPaymentResponse          
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Amount = temp.Amount
    g.DefaultPaymentMethod = temp.DefaultPaymentMethod
    g.SuccessUrl = temp.SuccessUrl
    g.PaymentUrl = temp.PaymentUrl
    g.GatewayAffiliationId = temp.GatewayAffiliationId
    if temp.AcceptedPaymentMethods == nil {
        g.AcceptedPaymentMethods = nil
    } else {
        g.AcceptedPaymentMethods = temp.AcceptedPaymentMethods
    }
    g.Status = temp.Status
    g.SkipCheckoutSuccessPage = temp.SkipCheckoutSuccessPage
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
    if temp.CanceledAt != nil {
        CanceledAtVal, err := time.Parse(time.RFC3339, *temp.CanceledAt)
        if err != nil {
            log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
        }
        g.CanceledAt = &CanceledAtVal
    } else {
        g.CanceledAt = nil
    }
    g.CustomerEditable = temp.CustomerEditable
    g.Customer = temp.Customer
    g.Billingaddress = temp.Billingaddress
    g.CreditCard = temp.CreditCard
    g.Boleto = temp.Boleto
    g.BillingAddressEditable = temp.BillingAddressEditable
    g.Shipping = temp.Shipping
    g.Shippable = temp.Shippable
    if temp.ClosedAt != nil {
        ClosedAtVal, err := time.Parse(time.RFC3339, *temp.ClosedAt)
        if err != nil {
            log.Fatalf("Cannot Parse closed_at as % s format.", time.RFC3339)
        }
        g.ClosedAt = &ClosedAtVal
    } else {
        g.ClosedAt = nil
    }
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        g.ExpiresAt = &ExpiresAtVal
    } else {
        g.ExpiresAt = nil
    }
    g.Currency = temp.Currency
    g.DebitCard = temp.DebitCard
    g.BankTransfer = temp.BankTransfer
    if temp.AcceptedBrands == nil {
        g.AcceptedBrands = nil
    } else {
        g.AcceptedBrands = temp.AcceptedBrands
    }
    g.Pix = temp.Pix
    return nil
}
