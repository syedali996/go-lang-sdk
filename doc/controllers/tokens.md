# Tokens

```go
tokensController := client.TokensController
```

## Class Name

`TokensController`

## Methods

* [Create Token](../../doc/controllers/tokens.md#create-token)
* [Get Token](../../doc/controllers/tokens.md#get-token)


# Create Token

:information_source: **Note** This endpoint does not require authentication.

```go
CreateToken(
    publicKey string,
    request models.CreateTokenRequest,
    idempotencyKey string) (
    https.ApiResponse[models.GetTokenResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `publicKey` | `string` | Template, Required | Public key |
| `request` | [`models.CreateTokenRequest`](../../doc/models/create-token-request.md) | Body, Required | Request for creating a token |
| `idempotencyKey` | `string` | Header, Optional | - |

## Response Type

[`models.GetTokenResponse`](../../doc/models/get-token-response.md)

## Example Usage

```go
publicKey := "public_key6"
requestCard := models.CreateCardTokenRequest { 
    Number: "number2",
    HolderName: "holder_name6",
    ExpMonth: 80,
    ExpYear: 216,
    Cvv: "cvv8",
    Brand: "brand4",
    Label: "label0",
}

request := models.CreateTokenRequest { 
    Type: "card",
    Card: requestCard,
}

apiResponse, err := tokensController.CreateToken(publicKey, request, idempotencyKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Token

Gets a token from its id

:information_source: **Note** This endpoint does not require authentication.

```go
GetToken(
    id string,
    publicKey string) (
    https.ApiResponse[models.GetTokenResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `string` | Template, Required | Token id |
| `publicKey` | `string` | Template, Required | Public key |

## Response Type

[`models.GetTokenResponse`](../../doc/models/get-token-response.md)

## Example Usage

```go
id := "id0"
publicKey := "public_key6"
apiResponse, err := tokensController.GetToken(id, publicKey)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

