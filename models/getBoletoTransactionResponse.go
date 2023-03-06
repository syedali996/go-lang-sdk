package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Response object for getting a boleto transaction
type GetBoletoTransactionResponse struct {
    GetTransactionResponse
    Url                    *string                    `json:"url"`
    Barcode                *string                    `json:"barcode"`
    NossoNumero            *string                    `json:"nosso_numero"`
    Bank                   *string                    `json:"bank"`
    DocumentNumber         *string                    `json:"document_number"`
    Instructions           *string                    `json:"instructions"`
    BillingAddress         *GetBillingAddressResponse `json:"billing_address"`
    DueAt                  *time.Time                 `json:"due_at"`
    QrCode                 *string                    `json:"qr_code"`
    Line                   *string                    `json:"line"`
    PdfPassword            *string                    `json:"pdf_password"`
    Pdf                    *string                    `json:"pdf"`
    PaidAt                 *time.Time                 `json:"paid_at"`
    PaidAmount             *string                    `json:"paid_amount"`
    Type                   *string                    `json:"type"`
    CreditAt               *time.Time                 `json:"credit_at"`
    StatementDescriptor    *string                    `json:"statement_descriptor"`
}

// Getter for url.
func (g *GetBoletoTransactionResponse) GetUrl() *string {
    return g.Url
}

// Setter for url.
func (g *GetBoletoTransactionResponse) SetUrl(url *string) {
    g.Url = url
}

// Getter for barcode.
func (g *GetBoletoTransactionResponse) GetBarcode() *string {
    return g.Barcode
}

// Setter for barcode.
func (g *GetBoletoTransactionResponse) SetBarcode(barcode *string) {
    g.Barcode = barcode
}

// Getter for nosso_numero.
func (g *GetBoletoTransactionResponse) GetNossoNumero() *string {
    return g.NossoNumero
}

// Setter for nosso_numero.
func (g *GetBoletoTransactionResponse) SetNossoNumero(nossoNumero *string) {
    g.NossoNumero = nossoNumero
}

// Getter for bank.
func (g *GetBoletoTransactionResponse) GetBank() *string {
    return g.Bank
}

// Setter for bank.
func (g *GetBoletoTransactionResponse) SetBank(bank *string) {
    g.Bank = bank
}

// Getter for document_number.
func (g *GetBoletoTransactionResponse) GetDocumentNumber() *string {
    return g.DocumentNumber
}

// Setter for document_number.
func (g *GetBoletoTransactionResponse) SetDocumentNumber(documentNumber *string) {
    g.DocumentNumber = documentNumber
}

// Getter for instructions.
func (g *GetBoletoTransactionResponse) GetInstructions() *string {
    return g.Instructions
}

// Setter for instructions.
func (g *GetBoletoTransactionResponse) SetInstructions(instructions *string) {
    g.Instructions = instructions
}

// Getter for billing_address.
func (g *GetBoletoTransactionResponse) GetBillingAddress() *GetBillingAddressResponse {
    return g.BillingAddress
}

// Setter for billing_address.
func (g *GetBoletoTransactionResponse) SetBillingAddress(billingAddress *GetBillingAddressResponse) {
    g.BillingAddress = billingAddress
}

// Getter for due_at.
func (g *GetBoletoTransactionResponse) GetDueAt() *time.Time {
    return g.DueAt
}

// Setter for due_at.
func (g *GetBoletoTransactionResponse) SetDueAt(dueAt *time.Time) {
    g.DueAt = dueAt
}

// Getter for qr_code.
func (g *GetBoletoTransactionResponse) GetQrCode() *string {
    return g.QrCode
}

// Setter for qr_code.
func (g *GetBoletoTransactionResponse) SetQrCode(qrCode *string) {
    g.QrCode = qrCode
}

// Getter for line.
func (g *GetBoletoTransactionResponse) GetLine() *string {
    return g.Line
}

// Setter for line.
func (g *GetBoletoTransactionResponse) SetLine(line *string) {
    g.Line = line
}

// Getter for pdf_password.
func (g *GetBoletoTransactionResponse) GetPdfPassword() *string {
    return g.PdfPassword
}

// Setter for pdf_password.
func (g *GetBoletoTransactionResponse) SetPdfPassword(pdfPassword *string) {
    g.PdfPassword = pdfPassword
}

// Getter for pdf.
func (g *GetBoletoTransactionResponse) GetPdf() *string {
    return g.Pdf
}

// Setter for pdf.
func (g *GetBoletoTransactionResponse) SetPdf(pdf *string) {
    g.Pdf = pdf
}

// Getter for paid_at.
func (g *GetBoletoTransactionResponse) GetPaidAt() *time.Time {
    return g.PaidAt
}

// Setter for paid_at.
func (g *GetBoletoTransactionResponse) SetPaidAt(paidAt *time.Time) {
    g.PaidAt = paidAt
}

// Getter for paid_amount.
func (g *GetBoletoTransactionResponse) GetPaidAmount() *string {
    return g.PaidAmount
}

// Setter for paid_amount.
func (g *GetBoletoTransactionResponse) SetPaidAmount(paidAmount *string) {
    g.PaidAmount = paidAmount
}

// Getter for type.
func (g *GetBoletoTransactionResponse) GetType() *string {
    return g.Type
}

// Setter for type.
func (g *GetBoletoTransactionResponse) SetType(mType *string) {
    g.Type = mType
}

// Getter for credit_at.
func (g *GetBoletoTransactionResponse) GetCreditAt() *time.Time {
    return g.CreditAt
}

// Setter for credit_at.
func (g *GetBoletoTransactionResponse) SetCreditAt(creditAt *time.Time) {
    g.CreditAt = creditAt
}

// Getter for statement_descriptor.
func (g *GetBoletoTransactionResponse) GetStatementDescriptor() *string {
    return g.StatementDescriptor
}

// Setter for statement_descriptor.
func (g *GetBoletoTransactionResponse) SetStatementDescriptor(statementDescriptor *string) {
    g.StatementDescriptor = statementDescriptor
}

func (g *GetBoletoTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    g.TransactionType = "boleto"
    var DueAtVal *string
    if g.DueAt != nil {
        str := g.DueAt.Format(time.RFC3339)
        DueAtVal = &str
    } else {
        DueAtVal = nil
    }
    var PaidAtVal *string
    if g.PaidAt != nil {
        str := g.PaidAt.Format(time.RFC3339)
        PaidAtVal = &str
    } else {
        PaidAtVal = nil
    }
    var CreditAtVal *string
    if g.CreditAt != nil {
        str := g.CreditAt.Format(time.RFC3339)
        CreditAtVal = &str
    } else {
        CreditAtVal = nil
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
    type marshallable GetBoletoTransactionResponse
    return json.Marshal(struct {
        marshallable 
        DueAt        *string `json:"dueAt"`       
        PaidAt       *string `json:"paidAt"`      
        CreditAt     *string `json:"creditAt"`    
        CreatedAt    *string `json:"createdAt"`   
        UpdatedAt    *string `json:"updatedAt"`   
        NextAttempt  *string `json:"nextAttempt"` 
    }{
        DueAt:        DueAtVal,         
        PaidAt:       PaidAtVal,        
        CreditAt:     CreditAtVal,      
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        NextAttempt:  NextAttemptVal,   
        marshallable: marshallable(*g), 
    })
}

func (g *GetBoletoTransactionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Url                 *string                     
        Barcode             *string                     
        NossoNumero         *string                     
        Bank                *string                     
        DocumentNumber      *string                     
        Instructions        *string                     
        BillingAddress      *GetBillingAddressResponse  
        DueAt               *string                     
        QrCode              *string                     
        Line                *string                     
        PdfPassword         *string                     
        Pdf                 *string                     
        PaidAt              *string                     
        PaidAmount          *string                     
        Type                *string                     
        CreditAt            *string                     
        StatementDescriptor *string                     
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
    g.Barcode = temp.Barcode
    g.NossoNumero = temp.NossoNumero
    g.Bank = temp.Bank
    g.DocumentNumber = temp.DocumentNumber
    g.Instructions = temp.Instructions
    g.BillingAddress = temp.BillingAddress
    if temp.DueAt != nil {
        DueAtVal, err := time.Parse(time.RFC3339, *temp.DueAt)
        if err != nil {
            log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
        }
        g.DueAt = &DueAtVal
    } else {
        g.DueAt = nil
    }
    g.QrCode = temp.QrCode
    g.Line = temp.Line
    g.PdfPassword = temp.PdfPassword
    g.Pdf = temp.Pdf
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
    g.Type = temp.Type
    if temp.CreditAt != nil {
        CreditAtVal, err := time.Parse(time.RFC3339, *temp.CreditAt)
        if err != nil {
            log.Fatalf("Cannot Parse credit_at as % s format.", time.RFC3339)
        }
        g.CreditAt = &CreditAtVal
    } else {
        g.CreditAt = nil
    }
    g.StatementDescriptor = temp.StatementDescriptor
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
type GetBoletoTransactionResponseSlice struct {
    Slice []GetBoletoTransactionResponseInterface 
}

func (ga *GetBoletoTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
    	ga.Slice = []GetBoletoTransactionResponseInterface{}
    	for _, getBoletoTransactionResponse := range temp {
    		if getBoletoTransactionResponse == nil {
    			ga.Slice = append(ga.Slice, nil)
    		}
    		var obj GetBoletoTransactionResponse
    		err := json.Unmarshal(getBoletoTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.Slice = append(ga.Slice, &obj)
    	}
    }
    return nil
}

func UnmarshalGetBoletoTransactionResponse(input []byte) (
    GetBoletoTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getBoletoTransactionResponse, ok := getTransactionResponse.(GetBoletoTransactionResponseInterface); ok {
        return getBoletoTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getBoletoTransactionResponse %v", getTransactionResponse)
}

func ToGetBoletoTransactionResponseArray(getBoletoTransactionResponse []GetBoletoTransactionResponseField) []GetBoletoTransactionResponseInterface {
    var items []GetBoletoTransactionResponseInterface
    for _, temp := range getBoletoTransactionResponse {
        items = append(items, temp.GetBoletoTransactionResponseInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetBoletoTransactionResponseField struct {
    GetBoletoTransactionResponseInterface
}

func (g *GetBoletoTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetBoletoTransactionResponseInterface, err = UnmarshalGetBoletoTransactionResponse(input)
    return err
}
