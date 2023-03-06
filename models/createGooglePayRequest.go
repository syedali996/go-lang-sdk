package models

// The GooglePay Token Payment Request
type CreateGooglePayRequest struct {
    Version            string                       `json:"version"`
    Data               string                       `json:"data"`
    Header             CreateGooglePayHeaderRequest `json:"header"`
    Signature          string                       `json:"signature"`
    MerchantIdentifier string                       `json:"merchant_identifier"`
}
