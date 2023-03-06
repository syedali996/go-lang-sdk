
# Get Increment Response

Response object for getting a increment

## Structure

`GetIncrementResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `value` | `*float64` | Optional | - |
| `incrementType` | `*string` | Optional | - |
| `status` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `cycles` | `*int` | Optional | - |
| `deletedAt` | `*time.Time` | Optional | - |
| `description` | `*string` | Optional | - |
| `subscription` | [`*models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md) | Optional | - |
| `subscriptionItem` | [`*models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md) | Optional | The Subscription Item |

## Example (as JSON)

```json
{
  "id": null,
  "value": null,
  "increment_type": null,
  "status": null,
  "created_at": null,
  "cycles": null,
  "deleted_at": null,
  "description": null,
  "subscription": null,
  "subscription_item": null
}
```

