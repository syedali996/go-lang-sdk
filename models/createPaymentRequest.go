package models

// Payment data
type CreatePaymentRequest struct {
    PaymentMethod        string                           `json:"payment_method"`
    CreditCard           CreateCreditCardPaymentRequest   `json:"credit_card,omitempty"`
    DebitCard            CreateDebitCardPaymentRequest    `json:"debit_card,omitempty"`
    Boleto               CreateBoletoPaymentRequest       `json:"boleto,omitempty"`
    Currency             string                           `json:"currency,omitempty"`
    Voucher              CreateVoucherPaymentRequest      `json:"voucher,omitempty"`
    Split                []CreateSplitRequest             `json:"split,omitempty"`
    BankTransfer         CreateBankTransferPaymentRequest `json:"bank_transfer,omitempty"`
    GatewayAffiliationId string                           `json:"gateway_affiliation_id,omitempty"`
    Amount               int                              `json:"amount,omitempty"`
    Checkout             CreateCheckoutPaymentRequest     `json:"checkout,omitempty"`
    CustomerId           string                           `json:"customer_id,omitempty"`
    Customer             CreateCustomerRequest            `json:"customer,omitempty"`
    Metadata             map[string]string                `json:"metadata,omitempty"`
    Cash                 CreateCashPaymentRequest         `json:"cash,omitempty"`
    PrivateLabel         CreatePrivateLabelPaymentRequest `json:"private_label,omitempty"`
    Pix                  CreatePixPaymentRequest          `json:"pix,omitempty"`
}
