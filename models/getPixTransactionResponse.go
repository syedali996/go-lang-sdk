package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Response object when getting a pix transaction
type GetPixTransactionResponse struct {
    GetTransactionResponse
    QrCode                 *string                     `json:"qr_code"`
    QrCodeUrl              *string                     `json:"qr_code_url"`
    ExpiresAt              *time.Time                  `json:"expires_at"`
    AdditionalInformation  *[]PixAdditionalInformation `json:"additional_information"`
    EndToEndId             *string                     `json:"end_to_end_id"`
    Payer                  *GetPixPayerResponse        `json:"payer"`
}

// Getter for qr_code.
func (g *GetPixTransactionResponse) GetQrCode() *string {
    return g.QrCode
}

// Setter for qr_code.
func (g *GetPixTransactionResponse) SetQrCode(qrCode *string) {
    g.QrCode = qrCode
}

// Getter for qr_code_url.
func (g *GetPixTransactionResponse) GetQrCodeUrl() *string {
    return g.QrCodeUrl
}

// Setter for qr_code_url.
func (g *GetPixTransactionResponse) SetQrCodeUrl(qrCodeUrl *string) {
    g.QrCodeUrl = qrCodeUrl
}

// Getter for expires_at.
func (g *GetPixTransactionResponse) GetExpiresAt() *time.Time {
    return g.ExpiresAt
}

// Setter for expires_at.
func (g *GetPixTransactionResponse) SetExpiresAt(expiresAt *time.Time) {
    g.ExpiresAt = expiresAt
}

// Getter for additional_information.
func (g *GetPixTransactionResponse) GetAdditionalInformation() *[]PixAdditionalInformation {
    return g.AdditionalInformation
}

// Setter for additional_information.
func (g *GetPixTransactionResponse) SetAdditionalInformation(additionalInformation *[]PixAdditionalInformation) {
    g.AdditionalInformation = additionalInformation
}

// Getter for end_to_end_id.
func (g *GetPixTransactionResponse) GetEndToEndId() *string {
    return g.EndToEndId
}

// Setter for end_to_end_id.
func (g *GetPixTransactionResponse) SetEndToEndId(endToEndId *string) {
    g.EndToEndId = endToEndId
}

// Getter for payer.
func (g *GetPixTransactionResponse) GetPayer() *GetPixPayerResponse {
    return g.Payer
}

// Setter for payer.
func (g *GetPixTransactionResponse) SetPayer(payer *GetPixPayerResponse) {
    g.Payer = payer
}

func (g *GetPixTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    g.TransactionType = "pix"
    var ExpiresAtVal *string
    if g.ExpiresAt != nil {
        str := g.ExpiresAt.Format(time.RFC3339)
        ExpiresAtVal = &str
    } else {
        ExpiresAtVal = nil
    }
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
    var NextAttemptVal *string
    if g.NextAttempt != nil {
        str := g.NextAttempt.Format(time.RFC3339)
        NextAttemptVal = &str
    } else {
        NextAttemptVal = nil
    }
    type marshallable GetPixTransactionResponse
    return json.Marshal(struct {
        marshallable 
        ExpiresAt    *string `json:"expiresAt"`   
        CreatedAt    *string `json:"createdAt"`   
        UpdatedAt    *string `json:"updatedAt"`   
        NextAttempt  *string `json:"nextAttempt"` 
    }{
        ExpiresAt:    ExpiresAtVal,     
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        NextAttempt:  NextAttemptVal,   
        marshallable: marshallable(*g), 
    })
}

func (g *GetPixTransactionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        QrCode                *string                     
        QrCodeUrl             *string                     
        ExpiresAt             *string                     
        AdditionalInformation *[]PixAdditionalInformation 
        EndToEndId            *string                     
        Payer                 *GetPixPayerResponse        
        GatewayId             *string                     
        Amount                *int                        
        Status                *string                     
        Success               *bool                       
        CreatedAt             *string                     
        UpdatedAt             *string                     
        AttemptCount          *int                        
        MaxAttempts           *int                        
        Splits                *[]GetSplitResponse         
        NextAttempt           *string                     
        TransactionType       string                      
        Id                    *string                     
        GatewayResponse       *GetGatewayResponseResponse 
        AntifraudResponse     *GetAntifraudResponse       
        Metadata              map[string]*string          
        Split                 *[]GetSplitResponse         
        Interest              *GetInterestResponse        
        Fine                  *GetFineResponse            
        MaxDaysToPayPastDue   *int                        
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.QrCode = temp.QrCode
    g.QrCodeUrl = temp.QrCodeUrl
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
    g.EndToEndId = temp.EndToEndId
    g.Payer = temp.Payer
    g.GatewayId = temp.GatewayId
    g.Amount = temp.Amount
    g.Status = temp.Status
    g.Success = temp.Success
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
    g.AttemptCount = temp.AttemptCount
    g.MaxAttempts = temp.MaxAttempts
    if temp.Splits == nil {
        g.Splits = nil
    } else {
        g.Splits = temp.Splits
    }
    if temp.NextAttempt != nil {
        NextAttemptVal, err := time.Parse(time.RFC3339, *temp.NextAttempt)
        if err != nil {
            log.Fatalf("Cannot Parse next_attempt as % s format.", time.RFC3339)
        }
        g.NextAttempt = &NextAttemptVal
    } else {
        g.NextAttempt = nil
    }
    g.TransactionType = temp.TransactionType
    g.Id = temp.Id
    g.GatewayResponse = temp.GatewayResponse
    g.AntifraudResponse = temp.AntifraudResponse
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    if temp.Split == nil {
        g.Split = nil
    } else {
        g.Split = temp.Split
    }
    g.Interest = temp.Interest
    g.Fine = temp.Fine
    g.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
    return nil
}

// Utility struct that helps in the unmarshalling of slices.
type GetPixTransactionResponseSlice struct {
    Slice []GetPixTransactionResponseInterface 
}

func (ga *GetPixTransactionResponseSlice) UnmarshalJSON(input []byte) error {
    tempStruct := struct {
    	Slice []json.RawMessage
    }{}
    temp := []json.RawMessage{}
    err := json.Unmarshal(input, &tempStruct)
    if err != nil {
    	err := json.Unmarshal(input, &temp)
    	if err != nil {
    		return err
    	}
    } else {
    	temp = tempStruct.Slice
    }
    
    ga.Slice = nil
    if temp != nil {
    	ga.Slice = []GetPixTransactionResponseInterface{}
    	for _, getPixTransactionResponse := range temp {
    		if getPixTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetPixTransactionResponse
    		err := json.Unmarshal(getPixTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}

func UnmarshalGetPixTransactionResponse(input []byte) (
    GetPixTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getPixTransactionResponse, ok := getTransactionResponse.(GetPixTransactionResponseInterface); ok {
        return getPixTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getPixTransactionResponse %v", getTransactionResponse)
}

func ToGetPixTransactionResponseArray(getPixTransactionResponse []GetPixTransactionResponseField) []GetPixTransactionResponseInterface {
    var items []GetPixTransactionResponseInterface
    for _, temp := range getPixTransactionResponse {
        items = append(items, temp.GetPixTransactionResponseInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetPixTransactionResponseField struct {
    GetPixTransactionResponseInterface
}

func (g *GetPixTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetPixTransactionResponseInterface, err = UnmarshalGetPixTransactionResponse(input)
    return err
}
