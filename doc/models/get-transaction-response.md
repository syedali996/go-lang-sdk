
# Get Transaction Response

Generic response object for getting a transaction.

## Structure

`GetTransactionResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `gatewayId` | `*string` | Optional | Gateway transaction id |
| `amount` | `*int` | Optional | Amount in cents |
| `status` | `*string` | Optional | Transaction status |
| `success` | `*bool` | Optional | Indicates if the transaction ocurred successfuly |
| `createdAt` | `*time.Time` | Optional | Creation date |
| `updatedAt` | `*time.Time` | Optional | Last update date |
| `attemptCount` | `*int` | Optional | Number of attempts tried |
| `maxAttempts` | `*int` | Optional | Max attempts |
| `splits` | [`*[]models.GetSplitResponse`](../../doc/models/get-split-response.md) | Optional | Splits |
| `nextAttempt` | `*time.Time` | Optional | Date and time of the next attempt |
| `transactionType` | `string` | Optional | - |
| `id` | `*string` | Optional | Código da transação |
| `gatewayResponse` | [`*models.GetGatewayResponseResponse`](../../doc/models/get-gateway-response-response.md) | Optional | The Gateway Response |
| `antifraudResponse` | [`*models.GetAntifraudResponse`](../../doc/models/get-antifraud-response.md) | Optional | - |
| `metadata` | `map[string]*string` | Optional | - |
| `split` | [`*[]models.GetSplitResponse`](../../doc/models/get-split-response.md) | Optional | - |
| `interest` | [`*models.GetInterestResponse`](../../doc/models/get-interest-response.md) | Optional | - |
| `fine` | [`*models.GetFineResponse`](../../doc/models/get-fine-response.md) | Optional | - |
| `maxDaysToPayPastDue` | `*int` | Optional | - |

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
  "transaction_type": "transaction",
  "id": null,
  "gateway_response": null,
  "antifraud_response": null,
  "metadata": null,
  "split": null,
  "interest": null,
  "fine": null,
  "max_days_to_pay_past_due": null
}
```

