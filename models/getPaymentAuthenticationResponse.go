package models

// Payment Authentication response
type GetPaymentAuthenticationResponse struct {
    Type         *string                  `json:"type"`
    ThreedSecure *GetThreeDSecureResponse `json:"threed_secure"`
}
