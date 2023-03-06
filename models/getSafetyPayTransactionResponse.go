package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Response object for getting a safety pay transaction
type GetSafetyPayTransactionResponse struct {
    GetTransactionResponse
    Url                    *string    `json:"url"`
    BankTid                *string    `json:"bank_tid"`
    PaidAt                 *time.Time `json:"paid_at"`
    PaidAmount             *int       `json:"paid_amount"`
}

// Getter for url.
func (g *GetSafetyPayTransactionResponse) GetUrl() *string {
    return g.Url
}

// Setter for url.
func (g *GetSafetyPayTransactionResponse) SetUrl(url *string) {
    g.Url = url
}

// Getter for bank_tid.
func (g *GetSafetyPayTransactionResponse) GetBankTid() *string {
    return g.BankTid
}

// Setter for bank_tid.
func (g *GetSafetyPayTransactionResponse) SetBankTid(bankTid *string) {
    g.BankTid = bankTid
}

// Getter for paid_at.
func (g *GetSafetyPayTransactionResponse) GetPaidAt() *time.Time {
    return g.PaidAt
}

// Setter for paid_at.
func (g *GetSafetyPayTransactionResponse) SetPaidAt(paidAt *time.Time) {
    g.PaidAt = paidAt
}

// Getter for paid_amount.
func (g *GetSafetyPayTransactionResponse) GetPaidAmount() *int {
    return g.PaidAmount
}

// Setter for paid_amount.
func (g *GetSafetyPayTransactionResponse) SetPaidAmount(paidAmount *int) {
    g.PaidAmount = paidAmount
}

func (g *GetSafetyPayTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    g.TransactionType = "safetypay"
    var PaidAtVal *string
    if g.PaidAt != nil {
        str := g.PaidAt.Format(time.RFC3339)
        PaidAtVal = &str
    } else {
        PaidAtVal = nil
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
    type marshallable GetSafetyPayTransactionResponse
    return json.Marshal(struct {
        marshallable 
        PaidAt       *string `json:"paidAt"`      
        CreatedAt    *string `json:"createdAt"`   
        UpdatedAt    *string `json:"updatedAt"`   
        NextAttempt  *string `json:"nextAttempt"` 
    }{
        PaidAt:       PaidAtVal,        
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        NextAttempt:  NextAttemptVal,   
        marshallable: marshallable(*g), 
    })
}

func (g *GetSafetyPayTransactionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Url                 *string                     
        BankTid             *string                     
        PaidAt              *string                     
        PaidAmount          *int                        
        GatewayId           *string                     
        Amount              *int                        
        Status              *string                     
        Success             *bool                       
        CreatedAt           *string                     
        UpdatedAt           *string                     
        AttemptCount        *int                        
        MaxAttempts         *int                        
        Splits              *[]GetSplitResponse         
        NextAttempt         *string                     
        TransactionType     string                      
        Id                  *string                     
        GatewayResponse     *GetGatewayResponseResponse 
        AntifraudResponse   *GetAntifraudResponse       
        Metadata            map[string]*string          
        Split               *[]GetSplitResponse         
        Interest            *GetInterestResponse        
        Fine                *GetFineResponse            
        MaxDaysToPayPastDue *int                        
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.Url = temp.Url
    g.BankTid = temp.BankTid
    if temp.PaidAt != nil {
        PaidAtVal, err := time.Parse(time.RFC3339, *temp.PaidAt)
        if err != nil {
            log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
        }
        g.PaidAt = &PaidAtVal
    } else {
        g.PaidAt = nil
    }
    g.PaidAmount = temp.PaidAmount
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
type GetSafetyPayTransactionResponseSlice struct {
    Slice []GetSafetyPayTransactionResponseInterface 
}

func (ga *GetSafetyPayTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetSafetyPayTransactionResponseInterface{}
    	for _, getSafetyPayTransactionResponse := range temp {
    		if getSafetyPayTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetSafetyPayTransactionResponse
    		err := json.Unmarshal(getSafetyPayTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}

func UnmarshalGetSafetyPayTransactionResponse(input []byte) (
    GetSafetyPayTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getSafetyPayTransactionResponse, ok := getTransactionResponse.(GetSafetyPayTransactionResponseInterface); ok {
        return getSafetyPayTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getSafetyPayTransactionResponse %v", getTransactionResponse)
}

func ToGetSafetyPayTransactionResponseArray(getSafetyPayTransactionResponse []GetSafetyPayTransactionResponseField) []GetSafetyPayTransactionResponseInterface {
    var items []GetSafetyPayTransactionResponseInterface
    for _, temp := range getSafetyPayTransactionResponse {
        items = append(items, temp.GetSafetyPayTransactionResponseInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetSafetyPayTransactionResponseField struct {
    GetSafetyPayTransactionResponseInterface
}

func (g *GetSafetyPayTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetSafetyPayTransactionResponseInterface, err = UnmarshalGetSafetyPayTransactionResponse(input)
    return err
}
