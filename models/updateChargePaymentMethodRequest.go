package models

// Request for updating the payment method of a charge
type UpdateChargePaymentMethodRequest struct {
    UpdateSubscription bool                             `json:"update_subscription"`
    PaymentMethod      string                           `json:"payment_method"`
    CreditCard         CreateCreditCardPaymentRequest   `json:"credit_card"`
    DebitCard          CreateDebitCardPaymentRequest    `json:"debit_card"`
    Boleto             CreateBoletoPaymentRequest       `json:"boleto"`
    Voucher            CreateVoucherPaymentRequest      `json:"voucher"`
    Cash               CreateCashPaymentRequest         `json:"cash"`
    BankTransfer       CreateBankTransferPaymentRequest `json:"bank_transfer"`
    PrivateLabel       CreatePrivateLabelPaymentRequest `json:"private_label"`
}
