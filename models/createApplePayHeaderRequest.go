package models

// The ApplePay header request
type CreateApplePayHeaderRequest struct {
    PublicKeyHash      string `json:"public_key_hash,omitempty"`
    EphemeralPublicKey string `json:"ephemeral_public_key"`
    TransactionId      string `json:"transaction_id,omitempty"`
}
