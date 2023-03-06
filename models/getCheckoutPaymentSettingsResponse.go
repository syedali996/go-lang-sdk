package models

// Checkout Payment Settings Response
type GetCheckoutPaymentSettingsResponse struct {
    SuccessUrl             *string              `json:"success_url"`
    PaymentUrl             *string              `json:"payment_url"`
    AcceptedPaymentMethods *[]string            `json:"accepted_payment_methods"`
    Status                 *string              `json:"status"`
    Customer               *GetCustomerResponse `json:"customer"`
    Amount                 *int                 `json:"amount"`
    DefaultPaymentMethod   *string              `json:"default_payment_method"`
    GatewayAffiliationId   *string              `json:"gateway_affiliation_id"`
}
