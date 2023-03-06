package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Response for voucher transactions
type GetVoucherTransactionResponse struct {
    GetTransactionResponse
    StatementDescriptor     *string          `json:"statement_descriptor"`
    AcquirerName            *string          `json:"acquirer_name"`
    AcquirerAffiliationCode *string          `json:"acquirer_affiliation_code"`
    AcquirerTid             *string          `json:"acquirer_tid"`
    AcquirerNsu             *string          `json:"acquirer_nsu"`
    AcquirerAuthCode        *string          `json:"acquirer_auth_code"`
    AcquirerMessage         *string          `json:"acquirer_message"`
    AcquirerReturnCode      *string          `json:"acquirer_return_code"`
    OperationType           *string          `json:"operation_type"`
    Card                    *GetCardResponse `json:"card"`
}

// Getter for statement_descriptor.
func (g *GetVoucherTransactionResponse) GetStatementDescriptor() *string {
    return g.StatementDescriptor
}

// Setter for statement_descriptor.
func (g *GetVoucherTransactionResponse) SetStatementDescriptor(statementDescriptor *string) {
    g.StatementDescriptor = statementDescriptor
}

// Getter for acquirer_name.
func (g *GetVoucherTransactionResponse) GetAcquirerName() *string {
    return g.AcquirerName
}

// Setter for acquirer_name.
func (g *GetVoucherTransactionResponse) SetAcquirerName(acquirerName *string) {
    g.AcquirerName = acquirerName
}

// Getter for acquirer_affiliation_code.
func (g *GetVoucherTransactionResponse) GetAcquirerAffiliationCode() *string {
    return g.AcquirerAffiliationCode
}

// Setter for acquirer_affiliation_code.
func (g *GetVoucherTransactionResponse) SetAcquirerAffiliationCode(acquirerAffiliationCode *string) {
    g.AcquirerAffiliationCode = acquirerAffiliationCode
}

// Getter for acquirer_tid.
func (g *GetVoucherTransactionResponse) GetAcquirerTid() *string {
    return g.AcquirerTid
}

// Setter for acquirer_tid.
func (g *GetVoucherTransactionResponse) SetAcquirerTid(acquirerTid *string) {
    g.AcquirerTid = acquirerTid
}

// Getter for acquirer_nsu.
func (g *GetVoucherTransactionResponse) GetAcquirerNsu() *string {
    return g.AcquirerNsu
}

// Setter for acquirer_nsu.
func (g *GetVoucherTransactionResponse) SetAcquirerNsu(acquirerNsu *string) {
    g.AcquirerNsu = acquirerNsu
}

// Getter for acquirer_auth_code.
func (g *GetVoucherTransactionResponse) GetAcquirerAuthCode() *string {
    return g.AcquirerAuthCode
}

// Setter for acquirer_auth_code.
func (g *GetVoucherTransactionResponse) SetAcquirerAuthCode(acquirerAuthCode *string) {
    g.AcquirerAuthCode = acquirerAuthCode
}

// Getter for acquirer_message.
func (g *GetVoucherTransactionResponse) GetAcquirerMessage() *string {
    return g.AcquirerMessage
}

// Setter for acquirer_message.
func (g *GetVoucherTransactionResponse) SetAcquirerMessage(acquirerMessage *string) {
    g.AcquirerMessage = acquirerMessage
}

// Getter for acquirer_return_code.
func (g *GetVoucherTransactionResponse) GetAcquirerReturnCode() *string {
    return g.AcquirerReturnCode
}

// Setter for acquirer_return_code.
func (g *GetVoucherTransactionResponse) SetAcquirerReturnCode(acquirerReturnCode *string) {
    g.AcquirerReturnCode = acquirerReturnCode
}

// Getter for operation_type.
func (g *GetVoucherTransactionResponse) GetOperationType() *string {
    return g.OperationType
}

// Setter for operation_type.
func (g *GetVoucherTransactionResponse) SetOperationType(operationType *string) {
    g.OperationType = operationType
}

// Getter for card.
func (g *GetVoucherTransactionResponse) GetCard() *GetCardResponse {
    return g.Card
}

// Setter for card.
func (g *GetVoucherTransactionResponse) SetCard(card *GetCardResponse) {
    g.Card = card
}

func (g *GetVoucherTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    g.TransactionType = "voucher"
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
    type marshallable GetVoucherTransactionResponse
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

func (g *GetVoucherTransactionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        StatementDescriptor     *string                     
        AcquirerName            *string                     
        AcquirerAffiliationCode *string                     
        AcquirerTid             *string                     
        AcquirerNsu             *string                     
        AcquirerAuthCode        *string                     
        AcquirerMessage         *string                     
        AcquirerReturnCode      *string                     
        OperationType           *string                     
        Card                    *GetCardResponse            
        GatewayId               *string                     
        Amount                  *int                        
        Status                  *string                     
        Success                 *bool                       
        CreatedAt               *string                     
        UpdatedAt               *string                     
        AttemptCount            *int                        
        MaxAttempts             *int                        
        Splits                  *[]GetSplitResponse         
        NextAttempt             *string                     
        TransactionType         string                      
        Id                      *string                     
        GatewayResponse         *GetGatewayResponseResponse 
        AntifraudResponse       *GetAntifraudResponse       
        Metadata                map[string]*string          
        Split                   *[]GetSplitResponse         
        Interest                *GetInterestResponse        
        Fine                    *GetFineResponse            
        MaxDaysToPayPastDue     *int                        
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.StatementDescriptor = temp.StatementDescriptor
    g.AcquirerName = temp.AcquirerName
    g.AcquirerAffiliationCode = temp.AcquirerAffiliationCode
    g.AcquirerTid = temp.AcquirerTid
    g.AcquirerNsu = temp.AcquirerNsu
    g.AcquirerAuthCode = temp.AcquirerAuthCode
    g.AcquirerMessage = temp.AcquirerMessage
    g.AcquirerReturnCode = temp.AcquirerReturnCode
    g.OperationType = temp.OperationType
    g.Card = temp.Card
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
type GetVoucherTransactionResponseSlice struct {
    Slice []GetVoucherTransactionResponseInterface 
}

func (ga *GetVoucherTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetVoucherTransactionResponseInterface{}
    	for _, getVoucherTransactionResponse := range temp {
    		if getVoucherTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetVoucherTransactionResponse
    		err := json.Unmarshal(getVoucherTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}

func UnmarshalGetVoucherTransactionResponse(input []byte) (
    GetVoucherTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getVoucherTransactionResponse, ok := getTransactionResponse.(GetVoucherTransactionResponseInterface); ok {
        return getVoucherTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getVoucherTransactionResponse %v", getTransactionResponse)
}

func ToGetVoucherTransactionResponseArray(getVoucherTransactionResponse []GetVoucherTransactionResponseField) []GetVoucherTransactionResponseInterface {
    var items []GetVoucherTransactionResponseInterface
    for _, temp := range getVoucherTransactionResponse {
        items = append(items, temp.GetVoucherTransactionResponseInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetVoucherTransactionResponseField struct {
    GetVoucherTransactionResponseInterface
}

func (g *GetVoucherTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetVoucherTransactionResponseInterface, err = UnmarshalGetVoucherTransactionResponse(input)
    return err
}
