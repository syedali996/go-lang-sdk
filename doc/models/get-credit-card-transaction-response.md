
# Get Credit Card Transaction Response

Response object for getting a credit card transaction

## Structure

`GetCreditCardTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `statementDescriptor` | `*string` | Optional | Text that will appear on the credit card's statement |
| `acquirerName` | `string` | Required | Acquirer name |
| `acquirerAffiliationCode` | `*string` | Optional | Aquirer affiliation code |
| `acquirerTid` | `string` | Required | Acquirer TID |
| `acquirerNsu` | `string` | Required | Acquirer NSU |
| `acquirerAuthCode` | `*string` | Optional | Acquirer authorization code |
| `operationType` | `*string` | Optional | Operation type |
| `card` | [`*models.GetCardResponse`](../../doc/models/get-card-response.md) | Optional | Card data |
| `acquirerMessage` | `*string` | Optional | Acquirer message |
| `acquirerReturnCode` | `*string` | Optional | Acquirer Return Code |
| `installments` | `*int` | Optional | Number of installments |
| `threedAuthenticationUrl` | `*string` | Optional | 3D-S authentication Url |

## Example (as JSON)

```json
{
  "gateway_id": null,
  "amount": null,
  "status": null,
  "success": null,
  "created_at": null,
  "updated_at": null,
  "attempt_count": null,
  "max_attempts": null,
  "splits": null,
  "next_attempt": null,
  "transaction_type": "credit_card",
  "id": null,
  "gateway_response": null,
  "antifraud_response": null,
  "metadata": null,
  "split": null,
  "interest": null,
  "fine": null,
  "max_days_to_pay_past_due": null,
  "statement_descriptor": null,
  "acquirer_name": "acquirer_name4",
  "acquirer_affiliation_code": null,
  "acquirer_tid": "acquirer_tid0",
  "acquirer_nsu": "acquirer_nsu0",
  "acquirer_auth_code": null,
  "operation_type": null,
  "card": null,
  "acquirer_message": null,
  "acquirer_return_code": null,
  "installments": null,
  "threed_authentication_url": null
}
```

