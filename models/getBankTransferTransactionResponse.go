package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Response object for getting a bank transfer transaction
type GetBankTransferTransactionResponse struct {
    GetTransactionResponse
    Url                    string    `json:"url"`
    BankTid                string    `json:"bank_tid"`
    Bank                   string    `json:"bank"`
    PaidAt                 time.Time `json:"paid_at,omitempty"`
    PaidAmount             int       `json:"paid_amount,omitempty"`
}

// Getter for url.
func (g *GetBankTransferTransactionResponse) GetUrl() string {
    return g.Url
}

// Setter for url.
func (g *GetBankTransferTransactionResponse) SetUrl(url string) {
    g.Url = url
}

// Getter for bank_tid.
func (g *GetBankTransferTransactionResponse) GetBankTid() string {
    return g.BankTid
}

// Setter for bank_tid.
func (g *GetBankTransferTransactionResponse) SetBankTid(bankTid string) {
    g.BankTid = bankTid
}

// Getter for bank.
func (g *GetBankTransferTransactionResponse) GetBank() string {
    return g.Bank
}

// Setter for bank.
func (g *GetBankTransferTransactionResponse) SetBank(bank string) {
    g.Bank = bank
}

// Getter for paid_at.
func (g *GetBankTransferTransactionResponse) GetPaidAt() time.Time {
    return g.PaidAt
}

// Setter for paid_at.
func (g *GetBankTransferTransactionResponse) SetPaidAt(paidAt time.Time) {
    g.PaidAt = paidAt
}

// Getter for paid_amount.
func (g *GetBankTransferTransactionResponse) GetPaidAmount() int {
    return g.PaidAmount
}

// Setter for paid_amount.
func (g *GetBankTransferTransactionResponse) SetPaidAmount(paidAmount int) {
    g.PaidAmount = paidAmount
}

func (g *GetBankTransferTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    g.TransactionType = "bank_transfer"
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
    type marshallable GetBankTransferTransactionResponse
    return json.Marshal(struct {
        marshallable 
        PaidAt       string  `json:"paidAt"`      
        CreatedAt    *string `json:"createdAt"`   
        UpdatedAt    *string `json:"updatedAt"`   
        NextAttempt  *string `json:"nextAttempt"` 
    }{
        PaidAt:       g.PaidAt.Format(time.RFC3339), 
        CreatedAt:    CreatedAtVal,                  
        UpdatedAt:    UpdatedAtVal,                  
        NextAttempt:  NextAttemptVal,                
        marshallable: marshallable(*g),              
    })
}

func (g *GetBankTransferTransactionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Url                 string                      
        BankTid             string                      
        Bank                string                      
        PaidAt              string                      
        PaidAmount          int                         
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
    g.Bank = temp.Bank
    g.PaidAt, err = time.Parse(time.RFC3339, temp.PaidAt)
    if err != nil {
        log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
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
type GetBankTransferTransactionResponseSlice struct {
    Slice []GetBankTransferTransactionResponseInterface 
}

func (ga *GetBankTransferTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetBankTransferTransactionResponseInterface{}
    	for _, getBankTransferTransactionResponse := range temp {
    		if getBankTransferTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetBankTransferTransactionResponse
    		err := json.Unmarshal(getBankTransferTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}

func UnmarshalGetBankTransferTransactionResponse(input []byte) (
    GetBankTransferTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getBankTransferTransactionResponse, ok := getTransactionResponse.(GetBankTransferTransactionResponseInterface); ok {
        return getBankTransferTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getBankTransferTransactionResponse %v", getTransactionResponse)
}

func ToGetBankTransferTransactionResponseArray(getBankTransferTransactionResponse []GetBankTransferTransactionResponseField) []GetBankTransferTransactionResponseInterface {
    var items []GetBankTransferTransactionResponseInterface
    for _, temp := range getBankTransferTransactionResponse {
        items = append(items, temp.GetBankTransferTransactionResponseInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetBankTransferTransactionResponseField struct {
    GetBankTransferTransactionResponseInterface
}

func (g *GetBankTransferTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetBankTransferTransactionResponseInterface, err = UnmarshalGetBankTransferTransactionResponse(input)
    return err
}
