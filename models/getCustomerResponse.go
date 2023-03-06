package models

import (
    "encoding/json"
    "log"
    "time"
)

// Response object for getting a customer
type GetCustomerResponse struct {
    Id            *string             `json:"id"`
    Name          *string             `json:"name"`
    Email         *string             `json:"email"`
    Delinquent    *bool               `json:"delinquent"`
    CreatedAt     *time.Time          `json:"created_at"`
    UpdatedAt     *time.Time          `json:"updated_at"`
    Document      *string             `json:"document"`
    Type          *string             `json:"type"`
    FbAccessToken *string             `json:"fb_access_token"`
    Address       *GetAddressResponse `json:"address"`
    Metadata      map[string]*string  `json:"metadata"`
    Phones        *GetPhonesResponse  `json:"phones"`
    FbId          *int64              `json:"fb_id"`
    Code          *string             `json:"code"`
    DocumentType  *string             `json:"document_type"`
}

func (g *GetCustomerResponse) MarshalJSON() (
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
    type marshallable GetCustomerResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"` 
        UpdatedAt    *string `json:"updatedAt"` 
    }{
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        marshallable: marshallable(*g), 
    })
}

func (g *GetCustomerResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id            *string             
        Name          *string             
        Email         *string             
        Delinquent    *bool               
        CreatedAt     *string             
        UpdatedAt     *string             
        Document      *string             
        Type          *string             
        FbAccessToken *string             
        Address       *GetAddressResponse 
        Metadata      map[string]*string  
        Phones        *GetPhonesResponse  
        FbId          *int64              
        Code          *string             
        DocumentType  *string             
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.Name = temp.Name
    g.Email = temp.Email
    g.Delinquent = temp.Delinquent
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
    g.Document = temp.Document
    g.Type = temp.Type
    g.FbAccessToken = temp.FbAccessToken
    g.Address = temp.Address
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    g.Phones = temp.Phones
    g.FbId = temp.FbId
    g.Code = temp.Code
    g.DocumentType = temp.DocumentType
    return nil
}
