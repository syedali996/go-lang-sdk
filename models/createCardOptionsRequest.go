package models

// Options for creating the card
type CreateCardOptionsRequest struct {
    VerifyCard bool `json:"verify_card"`
}
