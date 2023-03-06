# Recipients

```go
recipientsController := client.RecipientsController
```

## Class Name

`RecipientsController`

## Methods

* [Update Recipient](../../doc/controllers/recipients.md#update-recipient)
* [Create Anticipation](../../doc/controllers/recipients.md#create-anticipation)
* [Get Anticipation Limits](../../doc/controllers/recipients.md#get-anticipation-limits)
* [Get Recipients](../../doc/controllers/recipients.md#get-recipients)
* [Get Withdraw by Id](../../doc/controllers/recipients.md#get-withdraw-by-id)
* [Update Recipient Default Bank Account](../../doc/controllers/recipients.md#update-recipient-default-bank-account)
* [Update Recipient Metadata](../../doc/controllers/recipients.md#update-recipient-metadata)
* [Get Transfers](../../doc/controllers/recipients.md#get-transfers)
* [Get Transfer](../../doc/controllers/recipients.md#get-transfer)
* [Create Withdraw](../../doc/controllers/recipients.md#create-withdraw)
* [Update Automatic Anticipation Settings](../../doc/controllers/recipients.md#update-automatic-anticipation-settings)
* [Get Anticipation](../../doc/controllers/recipients.md#get-anticipation)
* [Update Recipient Transfer Settings](../../doc/controllers/recipients.md#update-recipient-transfer-settings)
* [Get Anticipations](../../doc/controllers/recipients.md#get-anticipations)
* [Get Recipient](../../doc/controllers/recipients.md#get-recipient)
* [Get Balance](../../doc/controllers/recipients.md#get-balance)
* [Get Withdrawals](../../doc/controllers/recipients.md#get-withdrawals)
* [Create Transfer](../../doc/controllers/recipients.md#create-transfer)
* [Create Recipient](../../doc/controllers/recipients.md#create-recipient)
* [Get Recipient by Code](../../doc/controllers/recipients.md#get-recipient-by-code)
* [Get Default Recipient](../../doc/controllers/recipients.md#get-default-recipient)


# Update Recipient

Updates a recipient

```go
UpdateRecipient(
    recipientId string,
    request models.UpdateRecipientRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `request` | [`models.UpdateRecipientRequest`](../../doc/models/update-recipient-request.md) | Body, Required | Recipient data |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
requestMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
request := models.UpdateRecipientRequest { 
    Name: "name6",
    Email: "email0",
    Description: "description6",
    Type: "type4",
    Status: "status8",
    Metadata: requestMetadata,
}

apiResponse, err := recipientsController.UpdateRecipient(recipientId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Anticipation

Creates an anticipation

```go
CreateAnticipation(
    recipientId string,
    request models.CreateAnticipationRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetAnticipationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `request` | [`models.CreateAnticipationRequest`](../../doc/models/create-anticipation-request.md) | Body, Required | Anticipation data |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetAnticipationResponse`](../../doc/models/get-anticipation-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
requestPaymentDate, err := time.Parse(time.RFC3339, "2016-03-13T12:52:32.123Z")
if err != nil {
    log.Fatalln(err)
}
request := models.CreateAnticipationRequest { 
    Amount: 242,
    Timeframe: "timeframe8",
    PaymentDate: requestPaymentDate,
}

apiResponse, err := recipientsController.CreateAnticipation(recipientId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Anticipation Limits

Gets the anticipation limits for a recipient

```go
GetAnticipationLimits(
    recipientId string,
    timeframe string,
    paymentDate time.Time) (
    https.ApiResponse[models.GetAnticipationLimitResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `timeframe` | `string` | Query, Required | Timeframe |
| `paymentDate` | `time.Time` | Query, Required | Anticipation payment date |

## Response Type

[`models.GetAnticipationLimitResponse`](../../doc/models/get-anticipation-limit-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
timeframe := "timeframe2"
paymentDate, err := time.Parse(time.RFC3339, "2016-03-13T12:52:32.123Z")
if err != nil {
    log.Fatalln(err)
}
apiResponse, err := recipientsController.GetAnticipationLimits(recipientId, timeframe, paymentDate)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Recipients

Retrieves paginated recipients information

```go
GetRecipients(
    page int,
    size int) (
    https.ApiResponse[models.ListRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `int` | Query, Optional | Page number |
| `size` | `int` | Query, Optional | Page size |

## Response Type

[`models.ListRecipientResponse`](../../doc/models/list-recipient-response.md)

## Example Usage

```go
apiResponse, err := recipientsController.GetRecipients(page, size)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Withdraw by Id

```go
GetWithdrawById(
    recipientId string,
    withdrawalId string) (
    https.ApiResponse[models.GetWithdrawResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | - |
| `withdrawalId` | `string` | Template, Required | - |

## Response Type

[`models.GetWithdrawResponse`](../../doc/models/get-withdraw-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
withdrawalId := "withdrawal_id2"
apiResponse, err := recipientsController.GetWithdrawById(recipientId, withdrawalId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Recipient Default Bank Account

Updates the default bank account from a recipient

```go
UpdateRecipientDefaultBankAccount(
    recipientId string,
    request models.UpdateRecipientBankAccountRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `request` | [`models.UpdateRecipientBankAccountRequest`](../../doc/models/update-recipient-bank-account-request.md) | Body, Required | Bank account data |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
requestBankAccountMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'.", "Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
requestBankAccount := models.CreateBankAccountRequest { 
    HolderName: "holder_name6",
    HolderType: "holder_type2",
    HolderDocument: "holder_document4",
    Bank: "bank8",
    BranchNumber: "branch_number6",
    BranchCheckDigit: "branch_check_digit6",
    AccountNumber: "account_number0",
    AccountCheckDigit: "account_check_digit6",
    Type: "type0",
    Metadata: requestBankAccountMetadata,
    PixKey: "pix_key4",
}

request := models.UpdateRecipientBankAccountRequest { 
    BankAccount: requestBankAccount,
    PaymentMode: "bank_transfer",
}

apiResponse, err := recipientsController.UpdateRecipientDefaultBankAccount(recipientId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Recipient Metadata

Updates recipient metadata

```go
UpdateRecipientMetadata(
    recipientId string,
    request models.UpdateMetadataRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Metadata |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
requestMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
request := models.UpdateMetadataRequest { 
    Metadata: requestMetadata,
}

apiResponse, err := recipientsController.UpdateRecipientMetadata(recipientId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Transfers

Gets a paginated list of transfers for the recipient

```go
GetTransfers(
    recipientId string,
    page int,
    size int,
    status string,
    createdSince time.Time,
    createdUntil time.Time) (
    https.ApiResponse[models.ListTransferResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `page` | `int` | Query, Optional | Page number |
| `size` | `int` | Query, Optional | Page size |
| `status` | `string` | Query, Optional | Filter for transfer status |
| `createdSince` | `time.Time` | Query, Optional | Filter for start range of transfer creation date |
| `createdUntil` | `time.Time` | Query, Optional | Filter for end range of transfer creation date |

## Response Type

[`models.ListTransferResponse`](../../doc/models/list-transfer-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
apiResponse, err := recipientsController.GetTransfers(recipientId, page, size, status, createdSince, createdUntil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Transfer

Gets a transfer

```go
GetTransfer(
    recipientId string,
    transferId string) (
    https.ApiResponse[models.GetTransferResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `transferId` | `string` | Template, Required | Transfer id |

## Response Type

[`models.GetTransferResponse`](../../doc/models/get-transfer-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
transferId := "transfer_id6"
apiResponse, err := recipientsController.GetTransfer(recipientId, transferId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Withdraw

```go
CreateWithdraw(
    recipientId string,
    request models.CreateWithdrawRequest) (
    https.ApiResponse[models.GetWithdrawResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | - |
| `request` | [`models.CreateWithdrawRequest`](../../doc/models/create-withdraw-request.md) | Body, Required | - |

## Response Type

[`models.GetWithdrawResponse`](../../doc/models/get-withdraw-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
request := models.CreateWithdrawRequest { 
    Amount: 242,
}

apiResponse, err := recipientsController.CreateWithdraw(recipientId, request)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Automatic Anticipation Settings

Updates recipient metadata

```go
UpdateAutomaticAnticipationSettings(
    recipientId string,
    request models.UpdateAutomaticAnticipationSettingsRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `request` | [`models.UpdateAutomaticAnticipationSettingsRequest`](../../doc/models/update-automatic-anticipation-settings-request.md) | Body, Required | Metadata |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
request := models.UpdateAutomaticAnticipationSettingsRequest { }

apiResponse, err := recipientsController.UpdateAutomaticAnticipationSettings(recipientId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Anticipation

Gets an anticipation

```go
GetAnticipation(
    recipientId string,
    anticipationId string) (
    https.ApiResponse[models.GetAnticipationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `anticipationId` | `string` | Template, Required | Anticipation id |

## Response Type

[`models.GetAnticipationResponse`](../../doc/models/get-anticipation-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
anticipationId := "anticipation_id0"
apiResponse, err := recipientsController.GetAnticipation(recipientId, anticipationId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Recipient Transfer Settings

```go
UpdateRecipientTransferSettings(
    recipientId string,
    request models.UpdateTransferSettingsRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient Identificator |
| `request` | [`models.UpdateTransferSettingsRequest`](../../doc/models/update-transfer-settings-request.md) | Body, Required | - |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
request := models.UpdateTransferSettingsRequest { 
    TransferEnabled: "transfer_enabled2",
    TransferInterval: "transfer_interval6",
    TransferDay: "transfer_day6",
}

apiResponse, err := recipientsController.UpdateRecipientTransferSettings(recipientId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Anticipations

Retrieves a paginated list of anticipations from a recipient

```go
GetAnticipations(
    recipientId string,
    page int,
    size int,
    status string,
    timeframe string,
    paymentDateSince time.Time,
    paymentDateUntil time.Time,
    createdSince time.Time,
    createdUntil time.Time) (
    https.ApiResponse[models.ListAnticipationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `page` | `int` | Query, Optional | Page number |
| `size` | `int` | Query, Optional | Page size |
| `status` | `string` | Query, Optional | Filter for anticipation status |
| `timeframe` | `string` | Query, Optional | Filter for anticipation timeframe |
| `paymentDateSince` | `time.Time` | Query, Optional | Filter for start range for anticipation payment date |
| `paymentDateUntil` | `time.Time` | Query, Optional | Filter for end range for anticipation payment date |
| `createdSince` | `time.Time` | Query, Optional | Filter for start range for anticipation creation date |
| `createdUntil` | `time.Time` | Query, Optional | Filter for end range for anticipation creation date |

## Response Type

[`models.ListAnticipationResponse`](../../doc/models/list-anticipation-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
apiResponse, err := recipientsController.GetAnticipations(recipientId, page, size, status, timeframe, paymentDateSince, paymentDateUntil, createdSince, createdUntil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Recipient

Retrieves recipient information

```go
GetRecipient(
    recipientId string) (
    https.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipiend id |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
apiResponse, err := recipientsController.GetRecipient(recipientId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Balance

Get balance information for a recipient

```go
GetBalance(
    recipientId string) (
    https.ApiResponse[models.GetBalanceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |

## Response Type

[`models.GetBalanceResponse`](../../doc/models/get-balance-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
apiResponse, err := recipientsController.GetBalance(recipientId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Withdrawals

Gets a paginated list of transfers for the recipient

```go
GetWithdrawals(
    recipientId string,
    page int,
    size int,
    status string,
    createdSince time.Time,
    createdUntil time.Time) (
    https.ApiResponse[models.ListWithdrawals],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | - |
| `page` | `int` | Query, Optional | - |
| `size` | `int` | Query, Optional | - |
| `status` | `string` | Query, Optional | - |
| `createdSince` | `time.Time` | Query, Optional | - |
| `createdUntil` | `time.Time` | Query, Optional | - |

## Response Type

[`models.ListWithdrawals`](../../doc/models/list-withdrawals.md)

## Example Usage

```go
recipientId := "recipient_id0"
apiResponse, err := recipientsController.GetWithdrawals(recipientId, page, size, status, createdSince, createdUntil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Transfer

Creates a transfer for a recipient

```go
CreateTransfer(
    recipientId string,
    request models.CreateTransferRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetTransferResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient Id |
| `request` | [`models.CreateTransferRequest`](../../doc/models/create-transfer-request.md) | Body, Required | Transfer data |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetTransferResponse`](../../doc/models/get-transfer-response.md)

## Example Usage

```go
recipientId := "recipient_id0"
requestMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
request := models.CreateTransferRequest { 
    Amount: 242,
    Metadata: requestMetadata,
}

apiResponse, err := recipientsController.CreateTransfer(recipientId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Recipient

Creates a new recipient

```go
CreateRecipient(
    request models.CreateRecipientRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `request` | [`models.CreateRecipientRequest`](../../doc/models/create-recipient-request.md) | Body, Required | Recipient data |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
requestDefaultBankAccountMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
requestDefaultBankAccount := models.CreateBankAccountRequest { 
    HolderName: "holder_name0",
    HolderType: "holder_type6",
    HolderDocument: "holder_document8",
    Bank: "bank2",
    BranchNumber: "branch_number0",
    BranchCheckDigit: "branch_check_digit0",
    AccountNumber: "account_number4",
    AccountCheckDigit: "account_check_digit0",
    Type: "type4",
    Metadata: requestDefaultBankAccountMetadata,
    PixKey: "pix_key8",
}

requestMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
request := models.CreateRecipientRequest { 
    Name: "name6",
    Email: "email0",
    Description: "description6",
    Document: "document0",
    Type: "type4",
    DefaultBankAccount: requestDefaultBankAccount,
    Metadata: requestMetadata,
    Code: "code4",
    PaymentMode: "bank_transfer",
}

apiResponse, err := recipientsController.CreateRecipient(request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Recipient by Code

Retrieves recipient information

```go
GetRecipientByCode(
    code string) (
    https.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `code` | `string` | Template, Required | Recipient code |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
code := "code8"
apiResponse, err := recipientsController.GetRecipientByCode(code)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Default Recipient

```go
GetDefaultRecipient() (
    https.ApiResponse[models.GetRecipientResponse],
    error)
```

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
apiResponse, err := recipientsController.GetDefaultRecipient()
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

