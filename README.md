
# Getting Started with PagarmeApiSDK

## Introduction

Pagarme API

## Installation

The following section explains how to use the pagarmeapisdk library in a new project.

### 1. Build the SDK

To make this SDK buildable to your local development environment run the following command:

```bash
$ go get ./...
```

This command will retrieve all the latest version of the package available in mod file.

### 2. Install the Package

To install this SDK as a package to your local development environment run the following command:

```bash
$ go get github_link_to_repo
```

This command will retrieve the latest version of the package available.

### 3. Use the Installed Package

To use the installed package add it to the import block of your package file where you need it as shown below:

```bash
import (
    "fmt"

    "github_link_to_repo"
)
```

Once the package is imported, you can easily initialize a client and make calls to specific endpoints.

### 4. Test the SDK

To test this SDK in your local development environment run the following command:

```bash
$ go test ./...
```

This command will run all test cases that are generated for this SDK.

## Initialize the API Client

**_Note:_** Documentation for the client can be found [here.](doc/client.md)

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `httpConfiguration` | `https.HttpConfiguration` | Configurable http client options.<br>*Default*: `DefaultHttpConfiguration()` |
| `basicAuthUserName` | `string` | The username to use with basic authentication |
| `basicAuthPassword` | `string` | The password to use with basic authentication |

The API client can be initialized as follows:

```go
config := pagarmeapisdk.ConfigurationFactory(
    pagarmeapisdk.WithBasicAuthUserName("BasicAuthUserName"), 
    pagarmeapisdk.WithBasicAuthPassword("BasicAuthPassword"),
)
config.LoadFromEnvironment()

client := pagarmeapisdk.NewClient(config)
```

## Authorization

This API uses `Basic Authentication`.

## API Errors

Here is the list of errors that the API might throw.

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Invalid request | [`ErrorException`](doc/models/error-exception.md) |
| 401 | Invalid API key | [`ErrorException`](doc/models/error-exception.md) |
| 404 | An informed resource was not found | [`ErrorException`](doc/models/error-exception.md) |
| 412 | Business validation error | [`ErrorException`](doc/models/error-exception.md) |
| 422 | Contract validation error | [`ErrorException`](doc/models/error-exception.md) |
| 500 | Internal server error | [`ErrorException`](doc/models/error-exception.md) |

## List of APIs

* [Subscriptions](doc/controllers/subscriptions.md)
* [Orders](doc/controllers/orders.md)
* [Plans](doc/controllers/plans.md)
* [Invoices](doc/controllers/invoices.md)
* [Customers](doc/controllers/customers.md)
* [Charges](doc/controllers/charges.md)
* [Recipients](doc/controllers/recipients.md)
* [Tokens](doc/controllers/tokens.md)
* [Transactions](doc/controllers/transactions.md)
* [Transfers](doc/controllers/transfers.md)

