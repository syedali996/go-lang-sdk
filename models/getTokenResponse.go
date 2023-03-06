package models

import (
    "encoding/json"
    "log"
    "time"
)

// Token data
type GetTokenResponse struct {
    Id        *string               `json:"id"`
    Type      *string               `json:"type"`
    CreatedAt *time.Time            `json:"created_at"`
    ExpiresAt *string               `json:"expires_at"`
    Card      *GetCardTokenResponse `json:"card"`
}

func (g *GetTokenResponse) MarshalJSON() (
    []byte,
    error) {
    var CreatedAtVal *string
    if g.CreatedAt != nil {
        str := g.CreatedAt.Format(time.RFC3339)
        CreatedAtVal = &str
    } else {
        CreatedAtVal = nil
    }
    type marshallable GetTokenResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"` 
    }{
        CreatedAt:    CreatedAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetTokenResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id        *string               
        Type      *string               
        CreatedAt *string               
        ExpiresAt *string               
        Card      *GetCardTokenResponse 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Type = temp.Type
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    } else {
        g.CreatedAt = nil
    }
    g.ExpiresAt = temp.ExpiresAt
    g.Card = temp.Card
    return nil
}
