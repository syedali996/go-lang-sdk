package models

// The GooglePay header request
type CreateGooglePayHeaderRequest struct {
    EphemeralPublicKey string `json:"ephemeral_public_key"`
}
