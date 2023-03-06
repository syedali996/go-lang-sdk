package models

// Checkout payment request
type CreateCheckoutPaymentRequest struct {
    AcceptedPaymentMethods      []string                               `json:"accepted_payment_methods"`
    AcceptedMultiPaymentMethods []map[string]interface{}               `json:"accepted_multi_payment_methods"`
    SuccessUrl                  string                                 `json:"success_url"`
    DefaultPaymentMethod        string                                 `json:"default_payment_method,omitempty"`
    GatewayAffiliationId        string                                 `json:"gateway_affiliation_id,omitempty"`
    CreditCard                  CreateCheckoutCreditCardPaymentRequest `json:"credit_card,omitempty"`
    DebitCard                   CreateCheckoutDebitCardPaymentRequest  `json:"debit_card,omitempty"`
    Boleto                      CreateCheckoutBoletoPaymentRequest     `json:"boleto,omitempty"`
    CustomerEditable            bool                                   `json:"customer_editable,omitempty"`
    ExpiresIn                   int                                    `json:"expires_in,omitempty"`
    SkipCheckoutSuccessPage     bool                                   `json:"skip_checkout_success_page"`
    BillingAddressEditable      bool                                   `json:"billing_address_editable"`
    BillingAddress              CreateAddressRequest                   `json:"billing_address"`
    BankTransfer                CreateCheckoutBankTransferRequest      `json:"bank_transfer,omitempty"`
    AcceptedBrands              []string                               `json:"accepted_brands"`
    Pix                         CreateCheckoutPixPaymentRequest        `json:"pix,omitempty"`
}
