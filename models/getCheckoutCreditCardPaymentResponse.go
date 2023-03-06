package models

type GetCheckoutCreditCardPaymentResponse struct {
    StatementDescriptor *string                                      `json:"statementDescriptor"`
    Installments        *[]GetCheckoutCardInstallmentOptionsResponse `json:"installments"`
    Authentication      *GetPaymentAuthenticationResponse            `json:"authentication"`
}
