package models

// The card payment contactless request
type CreateCardPaymentContactlessRequest struct {
    Type      string                  `json:"type"`
    ApplePay  CreateApplePayRequest   `json:"apple_pay,omitempty"`
    GooglePay CreateGooglePayRequest  `json:"google_pay,omitempty"`
    Emv       CreateEmvDecryptRequest `json:"emv,omitempty"`
}
