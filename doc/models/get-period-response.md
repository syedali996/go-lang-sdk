
# Get Period Response

Response object for getting a period

## Structure

`GetPeriodResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `startAt` | `*time.Time` | Optional | - |
| `endAt` | `*time.Time` | Optional | - |
| `id` | `*string` | Optional | - |
| `billingAt` | `*time.Time` | Optional | - |
| `subscription` | [`*models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md) | Optional | - |
| `status` | `*string` | Optional | - |
| `duration` | `*int` | Optional | - |
| `createdAt` | `*string` | Optional | - |
| `updatedAt` | `*string` | Optional | - |
| `cycle` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "start_at": null,
  "end_at": null,
  "id": null,
  "billing_at": null,
  "subscription": null,
  "status": null,
  "duration": null,
  "created_at": null,
  "updated_at": null,
  "cycle": null
}
```

