package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting a access token
type GetAccessTokenResponse struct {
    Id        *string              `json:"id"`
    Code      *string              `json:"code"`
    Status    *string              `json:"status"`
    CreatedAt *time.Time           `json:"created_at"`
    Customer  *GetCustomerResponse `json:"customer"`
}

func (g *GetAccessTokenResponse) MarshalJSON() (
    []byte,
    error) {
    var CreatedAtVal *string
    if g.CreatedAt != nil {
        str := g.CreatedAt.Format(time.RFC3339)
        CreatedAtVal = &str
    } else {
        CreatedAtVal = nil
    }
    type marshallable GetAccessTokenResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"` 
    }{
        CreatedAt:    CreatedAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetAccessTokenResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id        *string              
        Code      *string              
        Status    *string              
        CreatedAt *string              
        Customer  *GetCustomerResponse 
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Code = temp.Code
    g.Status = temp.Status
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    } else {
        g.CreatedAt = nil
    }
    g.Customer = temp.Customer
    return nil
}
