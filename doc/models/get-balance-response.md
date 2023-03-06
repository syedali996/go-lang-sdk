
# Get Balance Response

Balance

## Structure

`GetBalanceResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `currency` | `*string` | Optional | Currency |
| `availableAmount` | `*int64` | Optional | Amount available for transferring |
| `recipient` | [`*models.GetRecipientResponse`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `transferredAmount` | `*int64` | Optional | - |
| `waitingFundsAmount` | `*int64` | Optional | - |

## Example (as JSON)

```json
{
  "currency": null,
  "available_amount": null,
  "recipient": null,
  "transferred_amount": null,
  "waiting_funds_amount": null
}
```

