
# Get Anticipation Response

Anticipation

## Structure

`GetAnticipationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | Id |
| `requestedAmount` | `*int` | Optional | Requested amount |
| `approvedAmount` | `*int` | Optional | Approved amount |
| `recipient` | [`*models.GetRecipientResponse`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `pgid` | `*string` | Optional | Anticipation id on the gateway |
| `createdAt` | `*time.Time` | Optional | Creation date |
| `updatedAt` | `*time.Time` | Optional | Last update date |
| `paymentDate` | `*time.Time` | Optional | Payment date |
| `status` | `*string` | Optional | Status |
| `timeframe` | `*string` | Optional | Timeframe |

## Example (as JSON)

```json
{
  "id": null,
  "requested_amount": null,
  "approved_amount": null,
  "recipient": null,
  "pgid": null,
  "created_at": null,
  "updated_at": null,
  "payment_date": null,
  "status": null,
  "timeframe": null
}
```

