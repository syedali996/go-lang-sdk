package models

// The ApplePay Token Payment Request
type CreateApplePayRequest struct {
    Version            string                      `json:"version"`
    Data               string                      `json:"data"`
    Header             CreateApplePayHeaderRequest `json:"header"`
    Signature          string                      `json:"signature"`
    MerchantIdentifier string                      `json:"merchant_identifier"`
}
