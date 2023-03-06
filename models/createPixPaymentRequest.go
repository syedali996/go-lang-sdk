package models

import (
    "encoding/json"
    "log"
    "time"
)

// Contains information to create a pix payment
type CreatePixPaymentRequest struct {
    ExpiresAt             time.Time                  `json:"expires_at,omitempty"`
    ExpiresIn             int                        `json:"expires_in,omitempty"`
    AdditionalInformation []PixAdditionalInformation `json:"additional_information,omitempty"`
}

func (c *CreatePixPaymentRequest) MarshalJSON() (
    []byte,
    error) {
    type marshallable CreatePixPaymentRequest
    return json.Marshal(struct {
        marshallable 
        ExpiresAt    string `json:"expiresAt"` 
    }{
        ExpiresAt:    c.ExpiresAt.Format(time.RFC3339), 
        marshallable: marshallable(*c),                 
    })
}

func (c *CreatePixPaymentRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        ExpiresAt             string                     
        ExpiresIn             int                        
        AdditionalInformation []PixAdditionalInformation 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    c.ExpiresAt, err = time.Parse(time.RFC3339, temp.ExpiresAt)
    if err != nil {
        log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
    }
    c.ExpiresIn = temp.ExpiresIn
    c.AdditionalInformation = temp.AdditionalInformation
    return nil
}
