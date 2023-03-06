
# Client Class Documentation

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

## PagarmeApiSDK Client

The gateway for the SDK. This class acts as a factory for the Controllers and also holds the configuration of the SDK.

## Controllers

| Name | Description |
|  --- | --- |
| subscriptions | Gets SubscriptionsController |
| orders | Gets OrdersController |
| plans | Gets PlansController |
| invoices | Gets InvoicesController |
| customers | Gets CustomersController |
| charges | Gets ChargesController |
| recipients | Gets RecipientsController |
| tokens | Gets TokensController |
| transactions | Gets TransactionsController |
| transfers | Gets TransfersController |

