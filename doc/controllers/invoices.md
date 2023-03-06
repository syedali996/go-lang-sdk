# Invoices

```go
invoicesController := client.InvoicesController
```

## Class Name

`InvoicesController`

## Methods

* [Update Invoice Metadata](../../doc/controllers/invoices.md#update-invoice-metadata)
* [Get Partial Invoice](../../doc/controllers/invoices.md#get-partial-invoice)
* [Cancel Invoice](../../doc/controllers/invoices.md#cancel-invoice)
* [Create Invoice](../../doc/controllers/invoices.md#create-invoice)
* [Get Invoices](../../doc/controllers/invoices.md#get-invoices)
* [Get Invoice](../../doc/controllers/invoices.md#get-invoice)
* [Update Invoice Status](../../doc/controllers/invoices.md#update-invoice-status)


# Update Invoice Metadata

Updates the metadata from an invoice

```go
UpdateInvoiceMetadata(
    invoiceId string,
    request models.UpdateMetadataRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `invoiceId` | `string` | Template, Required | The invoice id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the invoice metadata |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
invoiceId := "invoice_id0"
requestMetadata := map[string]string{"Liquid error: Object of type 'System.Collections.Generic.KeyValuePair`2[System.String,System.Object]' cannot be converted to type 'System.String'."}
request := models.UpdateMetadataRequest { 
    Metadata: requestMetadata,
}

apiResponse, err := invoicesController.UpdateInvoiceMetadata(invoiceId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Partial Invoice

```go
GetPartialInvoice(
    subscriptionId string) (
    https.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
apiResponse, err := invoicesController.GetPartialInvoice(subscriptionId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Cancel Invoice

Cancels an invoice

```go
CancelInvoice(
    invoiceId string,
    idempotencyKey string) (
    https.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `invoiceId` | `string` | Template, Required | Invoice id |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
invoiceId := "invoice_id0"
apiResponse, err := invoicesController.CancelInvoice(invoiceId, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Invoice

Create an Invoice

```go
CreateInvoice(
    subscriptionId string,
    cycleId string,
    request models.CreateInvoiceRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `cycleId` | `string` | Template, Required | Cycle Id |
| `request` | [`models.CreateInvoiceRequest`](../../doc/models/create-invoice-request.md) | Body, Optional | - |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
subscriptionId := "subscription_id0"
cycleId := "cycle_id6"
apiResponse, err := invoicesController.CreateInvoice(subscriptionId, cycleId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Invoices

Gets all invoices

```go
GetInvoices(
    page int,
    size int,
    code string,
    customerId string,
    subscriptionId string,
    createdSince time.Time,
    createdUntil time.Time,
    status string,
    dueSince time.Time,
    dueUntil time.Time,
    customerDocument string) (
    https.ApiResponse[models.ListInvoicesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `int` | Query, Optional | Page number |
| `size` | `int` | Query, Optional | Page size |
| `code` | `string` | Query, Optional | Filter for Invoice's code |
| `customerId` | `string` | Query, Optional | Filter for Invoice's customer id |
| `subscriptionId` | `string` | Query, Optional | Filter for Invoice's subscription id |
| `createdSince` | `time.Time` | Query, Optional | Filter for Invoice's creation date start range |
| `createdUntil` | `time.Time` | Query, Optional | Filter for Invoices creation date end range |
| `status` | `string` | Query, Optional | Filter for Invoice's status |
| `dueSince` | `time.Time` | Query, Optional | Filter for Invoice's due date start range |
| `dueUntil` | `time.Time` | Query, Optional | Filter for Invoice's due date end range |
| `customerDocument` | `string` | Query, Optional | - |

## Response Type

[`models.ListInvoicesResponse`](../../doc/models/list-invoices-response.md)

## Example Usage

```go
apiResponse, err := invoicesController.GetInvoices(page, size, code, customerId, subscriptionId, createdSince, createdUntil, status, dueSince, dueUntil, customerDocument)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Invoice

Gets an invoice

```go
GetInvoice(
    invoiceId string) (
    https.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `invoiceId` | `string` | Template, Required | Invoice Id |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
invoiceId := "invoice_id0"
apiResponse, err := invoicesController.GetInvoice(invoiceId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Invoice Status

Updates the status from an invoice

```go
UpdateInvoiceStatus(
    invoiceId string,
    request models.UpdateInvoiceStatusRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `invoiceId` | `string` | Template, Required | Invoice Id |
| `request` | [`models.UpdateInvoiceStatusRequest`](../../doc/models/update-invoice-status-request.md) | Body, Required | Request for updating an invoice's status |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
invoiceId := "invoice_id0"
request := models.UpdateInvoiceStatusRequest { 
    Status: "status8",
}

apiResponse, err := invoicesController.UpdateInvoiceStatus(invoiceId, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

