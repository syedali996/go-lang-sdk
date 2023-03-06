
# Get Checkout Debit Card Payment Response

## Structure

`GetCheckoutDebitCardPaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `statementDescriptor` | `*string` | Optional | Descrição na fatura |
| `authentication` | [`*models.GetPaymentAuthenticationResponse`](../../doc/models/get-payment-authentication-response.md) | Optional | Payment Authentication response object data |

## Example (as JSON)

```json
{
  "statement_descriptor": null,
  "authentication": null
}
```

