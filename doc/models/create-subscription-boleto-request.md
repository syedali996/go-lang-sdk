
# Create Subscription Boleto Request

Information about fines and interest on the "boleto" used from payment

## Structure

`CreateSubscriptionBoletoRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `interest` | [`models.CreateInterestRequest`](../../doc/models/create-interest-request.md) | Optional | - |
| `fine` | [`models.CreateFineRequest`](../../doc/models/create-fine-request.md) | Optional | - |
| `maxDaysToPayPastDue` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "interest": null,
  "fine": null,
  "max_days_to_pay_past_due": null
}
```

