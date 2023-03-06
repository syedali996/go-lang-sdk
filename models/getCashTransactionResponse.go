package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Response object for getting a cash transaction
type GetCashTransactionResponse struct {
    GetTransactionResponse
    Description            *string `json:"description"`
}

// Getter for description.
func (g *GetCashTransactionResponse) GetDescription() *string {
    return g.Description
}

// Setter for description.
func (g *GetCashTransactionResponse) SetDescription(description *string) {
    g.Description = description
}

func (g *GetCashTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    g.TransactionType = "cash"
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
    type marshallable GetCashTransactionResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"`   
        UpdatedAt    *string `json:"updatedAt"`   
        NextAttempt  *string `json:"nextAttempt"` 
    }{
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        NextAttempt:  NextAttemptVal,   
        marshallable: marshallable(*g), 
    })
}

func (g *GetCashTransactionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Description         *string                     
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
    
    g.Description = temp.Description
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
type GetCashTransactionResponseSlice struct {
    Slice []GetCashTransactionResponseInterface 
}

func (ga *GetCashTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetCashTransactionResponseInterface{}
    	for _, getCashTransactionResponse := range temp {
    		if getCashTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetCashTransactionResponse
    		err := json.Unmarshal(getCashTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}

func UnmarshalGetCashTransactionResponse(input []byte) (
    GetCashTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getCashTransactionResponse, ok := getTransactionResponse.(GetCashTransactionResponseInterface); ok {
        return getCashTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getCashTransactionResponse %v", getTransactionResponse)
}

func ToGetCashTransactionResponseArray(getCashTransactionResponse []GetCashTransactionResponseField) []GetCashTransactionResponseInterface {
    var items []GetCashTransactionResponseInterface
    for _, temp := range getCashTransactionResponse {
        items = append(items, temp.GetCashTransactionResponseInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetCashTransactionResponseField struct {
    GetCashTransactionResponseInterface
}

func (g *GetCashTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetCashTransactionResponseInterface, err = UnmarshalGetCashTransactionResponse(input)
    return err
}
