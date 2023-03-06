package models

import (
    "time"
)

// Response object for listing transactions
type ListTransactionsResponseInterface interface {
    GetData() *[]GetTransactionResponseInterface
    SetData(data *[]GetTransactionResponseInterface)
    GetPaging() *PagingResponse
    SetPaging(paging *PagingResponse)
}

// Response object for getting a charge
type GetChargeResponseInterface interface {
    GetId() *string
    SetId(id *string)
    GetCode() *string
    SetCode(code *string)
    GetGatewayId() *string
    SetGatewayId(gatewayId *string)
    GetAmount() *int
    SetAmount(amount *int)
    GetStatus() *string
    SetStatus(status *string)
    GetCurrency() *string
    SetCurrency(currency *string)
    GetPaymentMethod() *string
    SetPaymentMethod(paymentMethod *string)
    GetDueAt() *time.Time
    SetDueAt(dueAt *time.Time)
    GetCreatedAt() *time.Time
    SetCreatedAt(createdAt *time.Time)
    GetUpdatedAt() *time.Time
    SetUpdatedAt(updatedAt *time.Time)
    GetLastTransaction() GetTransactionResponseInterface
    SetLastTransaction(lastTransaction GetTransactionResponseInterface)
    GetInvoice() *GetInvoiceResponse
    SetInvoice(invoice *GetInvoiceResponse)
    GetOrder() *GetOrderResponse
    SetOrder(order *GetOrderResponse)
    GetCustomer() *GetCustomerResponse
    SetCustomer(customer *GetCustomerResponse)
    GetMetadata() map[string]*string
    SetMetadata(metadata map[string]*string)
    GetPaidAt() *time.Time
    SetPaidAt(paidAt *time.Time)
    GetCanceledAt() *time.Time
    SetCanceledAt(canceledAt *time.Time)
    GetCanceledAmount() *int
    SetCanceledAmount(canceledAmount *int)
    GetPaidAmount() *int
    SetPaidAmount(paidAmount *int)
    GetInterestAndFinePaid() *int
    SetInterestAndFinePaid(interestAndFinePaid *int)
    GetRecurrencyCycle() *string
    SetRecurrencyCycle(recurrencyCycle *string)
}

// Response object for getting a bank transfer transaction
type GetBankTransferTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetUrl() string
    SetUrl(url string)
    GetBankTid() string
    SetBankTid(bankTid string)
    GetBank() string
    SetBank(bank string)
    GetPaidAt() time.Time
    SetPaidAt(paidAt time.Time)
    GetPaidAmount() int
    SetPaidAmount(paidAmount int)
}

// Response object for getting a safety pay transaction
type GetSafetyPayTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetUrl() *string
    SetUrl(url *string)
    GetBankTid() *string
    SetBankTid(bankTid *string)
    GetPaidAt() *time.Time
    SetPaidAt(paidAt *time.Time)
    GetPaidAmount() *int
    SetPaidAmount(paidAmount *int)
}

// Response for voucher transactions
type GetVoucherTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetStatementDescriptor() *string
    SetStatementDescriptor(statementDescriptor *string)
    GetAcquirerName() *string
    SetAcquirerName(acquirerName *string)
    GetAcquirerAffiliationCode() *string
    SetAcquirerAffiliationCode(acquirerAffiliationCode *string)
    GetAcquirerTid() *string
    SetAcquirerTid(acquirerTid *string)
    GetAcquirerNsu() *string
    SetAcquirerNsu(acquirerNsu *string)
    GetAcquirerAuthCode() *string
    SetAcquirerAuthCode(acquirerAuthCode *string)
    GetAcquirerMessage() *string
    SetAcquirerMessage(acquirerMessage *string)
    GetAcquirerReturnCode() *string
    SetAcquirerReturnCode(acquirerReturnCode *string)
    GetOperationType() *string
    SetOperationType(operationType *string)
    GetCard() *GetCardResponse
    SetCard(card *GetCardResponse)
}

// Response object for getting a boleto transaction
type GetBoletoTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetUrl() *string
    SetUrl(url *string)
    GetBarcode() *string
    SetBarcode(barcode *string)
    GetNossoNumero() *string
    SetNossoNumero(nossoNumero *string)
    GetBank() *string
    SetBank(bank *string)
    GetDocumentNumber() *string
    SetDocumentNumber(documentNumber *string)
    GetInstructions() *string
    SetInstructions(instructions *string)
    GetBillingAddress() *GetBillingAddressResponse
    SetBillingAddress(billingAddress *GetBillingAddressResponse)
    GetDueAt() *time.Time
    SetDueAt(dueAt *time.Time)
    GetQrCode() *string
    SetQrCode(qrCode *string)
    GetLine() *string
    SetLine(line *string)
    GetPdfPassword() *string
    SetPdfPassword(pdfPassword *string)
    GetPdf() *string
    SetPdf(pdf *string)
    GetPaidAt() *time.Time
    SetPaidAt(paidAt *time.Time)
    GetPaidAmount() *string
    SetPaidAmount(paidAmount *string)
    GetType() *string
    SetType(mType *string)
    GetCreditAt() *time.Time
    SetCreditAt(creditAt *time.Time)
    GetStatementDescriptor() *string
    SetStatementDescriptor(statementDescriptor *string)
}

// Response object for getting a debit card transaction
type GetDebitCardTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetStatementDescriptor() *string
    SetStatementDescriptor(statementDescriptor *string)
    GetAcquirerName() *string
    SetAcquirerName(acquirerName *string)
    GetAcquirerAffiliationCode() *string
    SetAcquirerAffiliationCode(acquirerAffiliationCode *string)
    GetAcquirerTid() *string
    SetAcquirerTid(acquirerTid *string)
    GetAcquirerNsu() *string
    SetAcquirerNsu(acquirerNsu *string)
    GetAcquirerAuthCode() *string
    SetAcquirerAuthCode(acquirerAuthCode *string)
    GetOperationType() *string
    SetOperationType(operationType *string)
    GetCard() *GetCardResponse
    SetCard(card *GetCardResponse)
    GetAcquirerMessage() *string
    SetAcquirerMessage(acquirerMessage *string)
    GetAcquirerReturnCode() *string
    SetAcquirerReturnCode(acquirerReturnCode *string)
    GetMpi() *string
    SetMpi(mpi *string)
    GetEci() *string
    SetEci(eci *string)
    GetAuthenticationType() *string
    SetAuthenticationType(authenticationType *string)
    GetThreedAuthenticationUrl() *string
    SetThreedAuthenticationUrl(threedAuthenticationUrl *string)
}

// Generic response object for getting a transaction.
type GetTransactionResponseInterface interface {
    GetGatewayId() *string
    SetGatewayId(gatewayId *string)
    GetAmount() *int
    SetAmount(amount *int)
    GetStatus() *string
    SetStatus(status *string)
    GetSuccess() *bool
    SetSuccess(success *bool)
    GetCreatedAt() *time.Time
    SetCreatedAt(createdAt *time.Time)
    GetUpdatedAt() *time.Time
    SetUpdatedAt(updatedAt *time.Time)
    GetAttemptCount() *int
    SetAttemptCount(attemptCount *int)
    GetMaxAttempts() *int
    SetMaxAttempts(maxAttempts *int)
    GetSplits() *[]GetSplitResponse
    SetSplits(splits *[]GetSplitResponse)
    GetNextAttempt() *time.Time
    SetNextAttempt(nextAttempt *time.Time)
    GetTransactionType() string
    SetTransactionType(transactionType string)
    GetId() *string
    SetId(id *string)
    GetGatewayResponse() *GetGatewayResponseResponse
    SetGatewayResponse(gatewayResponse *GetGatewayResponseResponse)
    GetAntifraudResponse() *GetAntifraudResponse
    SetAntifraudResponse(antifraudResponse *GetAntifraudResponse)
    GetMetadata() map[string]*string
    SetMetadata(metadata map[string]*string)
    GetSplit() *[]GetSplitResponse
    SetSplit(split *[]GetSplitResponse)
    GetInterest() *GetInterestResponse
    SetInterest(interest *GetInterestResponse)
    GetFine() *GetFineResponse
    SetFine(fine *GetFineResponse)
    GetMaxDaysToPayPastDue() *int
    SetMaxDaysToPayPastDue(maxDaysToPayPastDue *int)
}

// Response object for getting a private label transaction
type GetPrivateLabelTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetStatementDescriptor() *string
    SetStatementDescriptor(statementDescriptor *string)
    GetAcquirerName() *string
    SetAcquirerName(acquirerName *string)
    GetAcquirerAffiliationCode() *string
    SetAcquirerAffiliationCode(acquirerAffiliationCode *string)
    GetAcquirerTid() *string
    SetAcquirerTid(acquirerTid *string)
    GetAcquirerNsu() *string
    SetAcquirerNsu(acquirerNsu *string)
    GetAcquirerAuthCode() *string
    SetAcquirerAuthCode(acquirerAuthCode *string)
    GetOperationType() *string
    SetOperationType(operationType *string)
    GetCard() *GetCardResponse
    SetCard(card *GetCardResponse)
    GetAcquirerMessage() *string
    SetAcquirerMessage(acquirerMessage *string)
    GetAcquirerReturnCode() *string
    SetAcquirerReturnCode(acquirerReturnCode *string)
    GetInstallments() *int
    SetInstallments(installments *int)
}

// Response object for getting a cash transaction
type GetCashTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetDescription() *string
    SetDescription(description *string)
}

// Response object for getting a credit card transaction
type GetCreditCardTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetStatementDescriptor() *string
    SetStatementDescriptor(statementDescriptor *string)
    GetAcquirerName() string
    SetAcquirerName(acquirerName string)
    GetAcquirerAffiliationCode() *string
    SetAcquirerAffiliationCode(acquirerAffiliationCode *string)
    GetAcquirerTid() string
    SetAcquirerTid(acquirerTid string)
    GetAcquirerNsu() string
    SetAcquirerNsu(acquirerNsu string)
    GetAcquirerAuthCode() *string
    SetAcquirerAuthCode(acquirerAuthCode *string)
    GetOperationType() *string
    SetOperationType(operationType *string)
    GetCard() *GetCardResponse
    SetCard(card *GetCardResponse)
    GetAcquirerMessage() *string
    SetAcquirerMessage(acquirerMessage *string)
    GetAcquirerReturnCode() *string
    SetAcquirerReturnCode(acquirerReturnCode *string)
    GetInstallments() *int
    SetInstallments(installments *int)
    GetThreedAuthenticationUrl() *string
    SetThreedAuthenticationUrl(threedAuthenticationUrl *string)
}

// Response object for listing charge transactions
type ListChargeTransactionsResponseInterface interface {
    GetData() *[]GetTransactionResponseInterface
    SetData(data *[]GetTransactionResponseInterface)
    GetPaging() *PagingResponse
    SetPaging(paging *PagingResponse)
}

// Response object when getting a pix transaction
type GetPixTransactionResponseInterface interface {
    GetTransactionResponseInterface
    GetQrCode() *string
    SetQrCode(qrCode *string)
    GetQrCodeUrl() *string
    SetQrCodeUrl(qrCodeUrl *string)
    GetExpiresAt() *time.Time
    SetExpiresAt(expiresAt *time.Time)
    GetAdditionalInformation() *[]PixAdditionalInformation
    SetAdditionalInformation(additionalInformation *[]PixAdditionalInformation)
    GetEndToEndId() *string
    SetEndToEndId(endToEndId *string)
    GetPayer() *GetPixPayerResponse
    SetPayer(payer *GetPixPayerResponse)
}
