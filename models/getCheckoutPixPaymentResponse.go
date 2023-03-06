package models

import (
    "encoding/json"
    "log"
    "time"
)

// Checkout pix payment response
type GetCheckoutPixPaymentResponse struct {
    ExpiresAt             *time.Time                  `json:"expires_at"`
    AdditionalInformation *[]PixAdditionalInformation `json:"additional_information"`
}

func (g *GetCheckoutPixPaymentResponse) MarshalJSON() (
    []byte,
    error) {
    var ExpiresAtVal *string
    if g.ExpiresAt != nil {
        str := g.ExpiresAt.Format(time.RFC3339)
        ExpiresAtVal = &str
    } else {
        ExpiresAtVal = nil
    }
    type marshallable GetCheckoutPixPaymentResponse
    return json.Marshal(struct {
        marshallable 
        ExpiresAt    *string `json:"expiresAt"` 
    }{
        ExpiresAt:    ExpiresAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetCheckoutPixPaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        ExpiresAt             *string                     
        AdditionalInformation *[]PixAdditionalInformation 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        g.ExpiresAt = &ExpiresAtVal
    } else {
        g.ExpiresAt = nil
    }
    if temp.AdditionalInformation == nil {
        g.AdditionalInformation = nil
    } else {
        g.AdditionalInformation = temp.AdditionalInformation
    }
    return nil
}
