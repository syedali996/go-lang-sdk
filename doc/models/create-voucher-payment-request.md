
# Create Voucher Payment Request

The settings for creating a voucher payment

## Structure

`CreateVoucherPaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `statementDescriptor` | `string` | Optional | The text that will be shown on the voucher's statement |
| `cardId` | `string` | Optional | Card id |
| `cardToken` | `string` | Optional | Card token |
| `card` | [`models.CreateCardRequest`](../../doc/models/create-card-request.md) | Optional | Card info |
| `recurrencyCycle` | `string` | Optional | Defines whether the card has been used one or more times. |

## Example (as JSON)

```json
{
  "statement_descriptor": null,
  "card_id": null,
  "card_token": null,
  "Card": null,
  "recurrency_cycle": null
}
```

