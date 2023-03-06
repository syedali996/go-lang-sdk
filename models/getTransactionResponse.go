package models

import (
    "encoding/json"
    "log"
    "time"
)

// Generic response object for getting a transaction.
type GetTransactionResponse struct {
    GatewayId           *string                     `json:"gateway_id"`
    Amount              *int                        `json:"amount"`
    Status              *string                     `json:"status"`
    Success             *bool                       `json:"success"`
    CreatedAt           *time.Time                  `json:"created_at"`
    UpdatedAt           *time.Time                  `json:"updated_at"`
    AttemptCount        *int                        `json:"attempt_count"`
    MaxAttempts         *int                        `json:"max_attempts"`
    Splits              *[]GetSplitResponse         `json:"splits"`
    NextAttempt         *time.Time                  `json:"next_attempt"`
    TransactionType     string                      `json:"transaction_type,omitempty"`
    Id                  *string                     `json:"id"`
    GatewayResponse     *GetGatewayResponseResponse `json:"gateway_response"`
    AntifraudResponse   *GetAntifraudResponse       `json:"antifraud_response"`
    Metadata            map[string]*string          `json:"metadata"`
    Split               *[]GetSplitResponse         `json:"split"`
    Interest            *GetInterestResponse        `json:"interest"`
    Fine                *GetFineResponse            `json:"fine"`
    MaxDaysToPayPastDue *int                        `json:"max_days_to_pay_past_due"`
}

// Getter for gateway_id.
func (g *GetTransactionResponse) GetGatewayId() *string {
    return g.GatewayId
}

// Setter for gateway_id.
func (g *GetTransactionResponse) SetGatewayId(gatewayId *string) {
    g.GatewayId = gatewayId
}

// Getter for amount.
func (g *GetTransactionResponse) GetAmount() *int {
    return g.Amount
}

// Setter for amount.
func (g *GetTransactionResponse) SetAmount(amount *int) {
    g.Amount = amount
}

// Getter for status.
func (g *GetTransactionResponse) GetStatus() *string {
    return g.Status
}

// Setter for status.
func (g *GetTransactionResponse) SetStatus(status *string) {
    g.Status = status
}

// Getter for success.
func (g *GetTransactionResponse) GetSuccess() *bool {
    return g.Success
}

// Setter for success.
func (g *GetTransactionResponse) SetSuccess(success *bool) {
    g.Success = success
}

// Getter for created_at.
func (g *GetTransactionResponse) GetCreatedAt() *time.Time {
    return g.CreatedAt
}

// Setter for created_at.
func (g *GetTransactionResponse) SetCreatedAt(createdAt *time.Time) {
    g.CreatedAt = createdAt
}

// Getter for updated_at.
func (g *GetTransactionResponse) GetUpdatedAt() *time.Time {
    return g.UpdatedAt
}

// Setter for updated_at.
func (g *GetTransactionResponse) SetUpdatedAt(updatedAt *time.Time) {
    g.UpdatedAt = updatedAt
}

// Getter for attempt_count.
func (g *GetTransactionResponse) GetAttemptCount() *int {
    return g.AttemptCount
}

// Setter for attempt_count.
func (g *GetTransactionResponse) SetAttemptCount(attemptCount *int) {
    g.AttemptCount = attemptCount
}

// Getter for max_attempts.
func (g *GetTransactionResponse) GetMaxAttempts() *int {
    return g.MaxAttempts
}

// Setter for max_attempts.
func (g *GetTransactionResponse) SetMaxAttempts(maxAttempts *int) {
    g.MaxAttempts = maxAttempts
}

// Getter for splits.
func (g *GetTransactionResponse) GetSplits() *[]GetSplitResponse {
    return g.Splits
}

// Setter for splits.
func (g *GetTransactionResponse) SetSplits(splits *[]GetSplitResponse) {
    g.Splits = splits
}

// Getter for next_attempt.
func (g *GetTransactionResponse) GetNextAttempt() *time.Time {
    return g.NextAttempt
}

// Setter for next_attempt.
func (g *GetTransactionResponse) SetNextAttempt(nextAttempt *time.Time) {
    g.NextAttempt = nextAttempt
}

// Getter for transaction_type.
func (g *GetTransactionResponse) GetTransactionType() string {
    return g.TransactionType
}

// Setter for transaction_type.
func (g *GetTransactionResponse) SetTransactionType(transactionType string) {
    g.TransactionType = transactionType
}

// Getter for id.
func (g *GetTransactionResponse) GetId() *string {
    return g.Id
}

// Setter for id.
func (g *GetTransactionResponse) SetId(id *string) {
    g.Id = id
}

// Getter for gateway_response.
func (g *GetTransactionResponse) GetGatewayResponse() *GetGatewayResponseResponse {
    return g.GatewayResponse
}

// Setter for gateway_response.
func (g *GetTransactionResponse) SetGatewayResponse(gatewayResponse *GetGatewayResponseResponse) {
    g.GatewayResponse = gatewayResponse
}

// Getter for antifraud_response.
func (g *GetTransactionResponse) GetAntifraudResponse() *GetAntifraudResponse {
    return g.AntifraudResponse
}

// Setter for antifraud_response.
func (g *GetTransactionResponse) SetAntifraudResponse(antifraudResponse *GetAntifraudResponse) {
    g.AntifraudResponse = antifraudResponse
}

// Getter for metadata.
func (g *GetTransactionResponse) GetMetadata() map[string]*string {
    return g.Metadata
}

// Setter for metadata.
func (g *GetTransactionResponse) SetMetadata(metadata map[string]*string) {
    g.Metadata = metadata
}

// Getter for split.
func (g *GetTransactionResponse) GetSplit() *[]GetSplitResponse {
    return g.Split
}

// Setter for split.
func (g *GetTransactionResponse) SetSplit(split *[]GetSplitResponse) {
    g.Split = split
}

// Getter for interest.
func (g *GetTransactionResponse) GetInterest() *GetInterestResponse {
    return g.Interest
}

// Setter for interest.
func (g *GetTransactionResponse) SetInterest(interest *GetInterestResponse) {
    g.Interest = interest
}

// Getter for fine.
func (g *GetTransactionResponse) GetFine() *GetFineResponse {
    return g.Fine
}

// Setter for fine.
func (g *GetTransactionResponse) SetFine(fine *GetFineResponse) {
    g.Fine = fine
}

// Getter for max_days_to_pay_past_due.
func (g *GetTransactionResponse) GetMaxDaysToPayPastDue() *int {
    return g.MaxDaysToPayPastDue
}

// Setter for max_days_to_pay_past_due.
func (g *GetTransactionResponse) SetMaxDaysToPayPastDue(maxDaysToPayPastDue *int) {
    g.MaxDaysToPayPastDue = maxDaysToPayPastDue
}

func (g *GetTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    g.TransactionType = "transaction"
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
    type marshallable GetTransactionResponse
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

func (g *GetTransactionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
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
type GetTransactionResponseSlice struct {
    Slice []GetTransactionResponseInterface 
}

func (ga *GetTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetTransactionResponseInterface{}
    	for _, getTransactionResponse := range temp {
    		if getTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetTransactionResponse
    		err := json.Unmarshal(getTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}

func UnmarshalGetTransactionResponse(input []byte) (
    GetTransactionResponseInterface,
    error) {
    discrim := &struct {
        TransactionType *string
    }{}
    
    err := json.Unmarshal(input, &discrim)
    if err != nil {
        return nil, err
    }
    if discrim.TransactionType == nil {
        transactionType := "transaction"
        discrim.TransactionType = &transactionType
    }
    
    var res GetTransactionResponseInterface
    switch *discrim.TransactionType {
    case "bank_transfer":
        res = &GetBankTransferTransactionResponse{}
    case "safetypay":
        res = &GetSafetyPayTransactionResponse{}
    case "voucher":
        res = &GetVoucherTransactionResponse{}
    case "boleto":
        res = &GetBoletoTransactionResponse{}
    case "debit_card":
        res = &GetDebitCardTransactionResponse{}
    case "private_label":
        res = &GetPrivateLabelTransactionResponse{}
    case "cash":
        res = &GetCashTransactionResponse{}
    case "credit_card":
        res = &GetCreditCardTransactionResponse{}
    case "pix":
        res = &GetPixTransactionResponse{}
    default:
        res = &GetTransactionResponse{}
    }
    json.Unmarshal(input, res)
    return res, nil
}

func ToGetTransactionResponseArray(getTransactionResponse []GetTransactionResponseField) []GetTransactionResponseInterface {
    var items []GetTransactionResponseInterface
    for _, temp := range getTransactionResponse {
        items = append(items, temp.GetTransactionResponseInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetTransactionResponseField struct {
    GetTransactionResponseInterface
}

func (g *GetTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetTransactionResponseInterface, err = UnmarshalGetTransactionResponse(input)
    return err
}
