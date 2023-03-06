package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Response object for getting a credit card transaction
type GetCreditCardTransactionResponse struct {
    GetTransactionResponse
    StatementDescriptor     *string          `json:"statement_descriptor"`
    AcquirerName            string           `json:"acquirer_name"`
    AcquirerAffiliationCode *string          `json:"acquirer_affiliation_code"`
    AcquirerTid             string           `json:"acquirer_tid"`
    AcquirerNsu             string           `json:"acquirer_nsu"`
    AcquirerAuthCode        *string          `json:"acquirer_auth_code"`
    OperationType           *string          `json:"operation_type"`
    Card                    *GetCardResponse `json:"card"`
    AcquirerMessage         *string          `json:"acquirer_message"`
    AcquirerReturnCode      *string          `json:"acquirer_return_code"`
    Installments            *int             `json:"installments"`
    ThreedAuthenticationUrl *string          `json:"threed_authentication_url"`
}

// Getter for statement_descriptor.
func (g *GetCreditCardTransactionResponse) GetStatementDescriptor() *string {
    return g.StatementDescriptor
}

// Setter for statement_descriptor.
func (g *GetCreditCardTransactionResponse) SetStatementDescriptor(statementDescriptor *string) {
    g.StatementDescriptor = statementDescriptor
}

// Getter for acquirer_name.
func (g *GetCreditCardTransactionResponse) GetAcquirerName() string {
    return g.AcquirerName
}

// Setter for acquirer_name.
func (g *GetCreditCardTransactionResponse) SetAcquirerName(acquirerName string) {
    g.AcquirerName = acquirerName
}

// Getter for acquirer_affiliation_code.
func (g *GetCreditCardTransactionResponse) GetAcquirerAffiliationCode() *string {
    return g.AcquirerAffiliationCode
}

// Setter for acquirer_affiliation_code.
func (g *GetCreditCardTransactionResponse) SetAcquirerAffiliationCode(acquirerAffiliationCode *string) {
    g.AcquirerAffiliationCode = acquirerAffiliationCode
}

// Getter for acquirer_tid.
func (g *GetCreditCardTransactionResponse) GetAcquirerTid() string {
    return g.AcquirerTid
}

// Setter for acquirer_tid.
func (g *GetCreditCardTransactionResponse) SetAcquirerTid(acquirerTid string) {
    g.AcquirerTid = acquirerTid
}

// Getter for acquirer_nsu.
func (g *GetCreditCardTransactionResponse) GetAcquirerNsu() string {
    return g.AcquirerNsu
}

// Setter for acquirer_nsu.
func (g *GetCreditCardTransactionResponse) SetAcquirerNsu(acquirerNsu string) {
    g.AcquirerNsu = acquirerNsu
}

// Getter for acquirer_auth_code.
func (g *GetCreditCardTransactionResponse) GetAcquirerAuthCode() *string {
    return g.AcquirerAuthCode
}

// Setter for acquirer_auth_code.
func (g *GetCreditCardTransactionResponse) SetAcquirerAuthCode(acquirerAuthCode *string) {
    g.AcquirerAuthCode = acquirerAuthCode
}

// Getter for operation_type.
func (g *GetCreditCardTransactionResponse) GetOperationType() *string {
    return g.OperationType
}

// Setter for operation_type.
func (g *GetCreditCardTransactionResponse) SetOperationType(operationType *string) {
    g.OperationType = operationType
}

// Getter for card.
func (g *GetCreditCardTransactionResponse) GetCard() *GetCardResponse {
    return g.Card
}

// Setter for card.
func (g *GetCreditCardTransactionResponse) SetCard(card *GetCardResponse) {
    g.Card = card
}

// Getter for acquirer_message.
func (g *GetCreditCardTransactionResponse) GetAcquirerMessage() *string {
    return g.AcquirerMessage
}

// Setter for acquirer_message.
func (g *GetCreditCardTransactionResponse) SetAcquirerMessage(acquirerMessage *string) {
    g.AcquirerMessage = acquirerMessage
}

// Getter for acquirer_return_code.
func (g *GetCreditCardTransactionResponse) GetAcquirerReturnCode() *string {
    return g.AcquirerReturnCode
}

// Setter for acquirer_return_code.
func (g *GetCreditCardTransactionResponse) SetAcquirerReturnCode(acquirerReturnCode *string) {
    g.AcquirerReturnCode = acquirerReturnCode
}

// Getter for installments.
func (g *GetCreditCardTransactionResponse) GetInstallments() *int {
    return g.Installments
}

// Setter for installments.
func (g *GetCreditCardTransactionResponse) SetInstallments(installments *int) {
    g.Installments = installments
}

// Getter for threed_authentication_url.
func (g *GetCreditCardTransactionResponse) GetThreedAuthenticationUrl() *string {
    return g.ThreedAuthenticationUrl
}

// Setter for threed_authentication_url.
func (g *GetCreditCardTransactionResponse) SetThreedAuthenticationUrl(threedAuthenticationUrl *string) {
    g.ThreedAuthenticationUrl = threedAuthenticationUrl
}

func (g *GetCreditCardTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    g.TransactionType = "credit_card"
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
    type marshallable GetCreditCardTransactionResponse
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

func (g *GetCreditCardTransactionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        AcquirerName            string                      
        AcquirerTid             string                      
        AcquirerNsu             string                      
        StatementDescriptor     *string                     
        AcquirerAffiliationCode *string                     
        AcquirerAuthCode        *string                     
        OperationType           *string                     
        Card                    *GetCardResponse            
        AcquirerMessage         *string                     
        AcquirerReturnCode      *string                     
        Installments            *int                        
        ThreedAuthenticationUrl *string                     
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
    
    g.AcquirerName = temp.AcquirerName
    g.AcquirerTid = temp.AcquirerTid
    g.AcquirerNsu = temp.AcquirerNsu
    g.StatementDescriptor = temp.StatementDescriptor
    g.AcquirerAffiliationCode = temp.AcquirerAffiliationCode
    g.AcquirerAuthCode = temp.AcquirerAuthCode
    g.OperationType = temp.OperationType
    g.Card = temp.Card
    g.AcquirerMessage = temp.AcquirerMessage
    g.AcquirerReturnCode = temp.AcquirerReturnCode
    g.Installments = temp.Installments
    g.ThreedAuthenticationUrl = temp.ThreedAuthenticationUrl
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
type GetCreditCardTransactionResponseSlice struct {
    Slice []GetCreditCardTransactionResponseInterface 
}

func (ga *GetCreditCardTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetCreditCardTransactionResponseInterface{}
    	for _, getCreditCardTransactionResponse := range temp {
    		if getCreditCardTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetCreditCardTransactionResponse
    		err := json.Unmarshal(getCreditCardTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}

func UnmarshalGetCreditCardTransactionResponse(input []byte) (
    GetCreditCardTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getCreditCardTransactionResponse, ok := getTransactionResponse.(GetCreditCardTransactionResponseInterface); ok {
        return getCreditCardTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getCreditCardTransactionResponse %v", getTransactionResponse)
}

func ToGetCreditCardTransactionResponseArray(getCreditCardTransactionResponse []GetCreditCardTransactionResponseField) []GetCreditCardTransactionResponseInterface {
    var items []GetCreditCardTransactionResponseInterface
    for _, temp := range getCreditCardTransactionResponse {
        items = append(items, temp.GetCreditCardTransactionResponseInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetCreditCardTransactionResponseField struct {
    GetCreditCardTransactionResponseInterface
}

func (g *GetCreditCardTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetCreditCardTransactionResponseInterface, err = UnmarshalGetCreditCardTransactionResponse(input)
    return err
}
