package models

// The payment authentication request
type CreatePaymentAuthenticationRequest struct {
    Type         string                    `json:"type"`
    ThreedSecure CreateThreeDSecureRequest `json:"threed_secure"`
}
