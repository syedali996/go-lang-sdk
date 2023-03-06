package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting an Address
type GetAddressResponse struct {
    Id           *string              `json:"id"`
    Street       *string              `json:"street"`
    Number       *string              `json:"number"`
    Complement   *string              `json:"complement"`
    ZipCode      *string              `json:"zip_code"`
    Neighborhood *string              `json:"neighborhood"`
    City         *string              `json:"city"`
    State        *string              `json:"state"`
    Country      *string              `json:"country"`
    Status       *string              `json:"status"`
    CreatedAt    *time.Time           `json:"created_at"`
    UpdatedAt    *time.Time           `json:"updated_at"`
    Customer     *GetCustomerResponse `json:"customer"`
    Metadata     map[string]*string   `json:"metadata"`
    Line1        *string              `json:"line_1"`
    Line2        *string              `json:"line_2"`
    DeletedAt    *time.Time           `json:"deleted_at"`
}

func (g *GetAddressResponse) MarshalJSON() (
    []byte,
    error) {
    var CreatedAtVal *string
    if g.CreatedAt != nil {
        str := g.CreatedAt.Format(time.RFC3339)
        CreatedAtVal = &str
    } else {
        CreatedAtVal = nil
    }
    var UpdatedAtVal *string
    if g.UpdatedAt != nil {
        str := g.UpdatedAt.Format(time.RFC3339)
        UpdatedAtVal = &str
    } else {
        UpdatedAtVal = nil
    }
    var DeletedAtVal *string
    if g.DeletedAt != nil {
        str := g.DeletedAt.Format(time.RFC3339)
        DeletedAtVal = &str
    } else {
        DeletedAtVal = nil
    }
    type marshallable GetAddressResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"` 
        UpdatedAt    *string `json:"updatedAt"` 
        DeletedAt    *string `json:"deletedAt"` 
    }{
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        DeletedAt:    DeletedAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetAddressResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id           *string              
        Street       *string              
        Number       *string              
        Complement   *string              
        ZipCode      *string              
        Neighborhood *string              
        City         *string              
        State        *string              
        Country      *string              
        Status       *string              
        CreatedAt    *string              
        UpdatedAt    *string              
        Customer     *GetCustomerResponse 
        Metadata     map[string]*string   
        Line1        *string              
        Line2        *string              
        DeletedAt    *string              
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Street = temp.Street
    g.Number = temp.Number
    g.Complement = temp.Complement
    g.ZipCode = temp.ZipCode
    g.Neighborhood = temp.Neighborhood
    g.City = temp.City
    g.State = temp.State
    g.Country = temp.Country
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
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        g.UpdatedAt = &UpdatedAtVal
    } else {
        g.UpdatedAt = nil
    }
    g.Customer = temp.Customer
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    g.Line1 = temp.Line1
    g.Line2 = temp.Line2
    if temp.DeletedAt != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, *temp.DeletedAt)
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        g.DeletedAt = &DeletedAtVal
    } else {
        g.DeletedAt = nil
    }
    return nil
}
