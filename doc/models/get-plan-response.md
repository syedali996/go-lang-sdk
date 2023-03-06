
# Get Plan Response

Response object for getting a plan

## Structure

`GetPlanResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `name` | `*string` | Optional | - |
| `description` | `*string` | Optional | - |
| `url` | `*string` | Optional | - |
| `statementDescriptor` | `*string` | Optional | - |
| `interval` | `*string` | Optional | - |
| `intervalCount` | `*int` | Optional | - |
| `billingType` | `*string` | Optional | - |
| `paymentMethods` | `*[]string` | Optional | - |
| `installments` | `*[]int` | Optional | - |
| `status` | `*string` | Optional | - |
| `currency` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `updatedAt` | `*time.Time` | Optional | - |
| `items` | [`*[]models.GetPlanItemResponse`](../../doc/models/get-plan-item-response.md) | Optional | - |
| `billingDays` | `*[]int` | Optional | - |
| `shippable` | `*bool` | Optional | - |
| `metadata` | `map[string]*string` | Optional | - |
| `trialPeriodDays` | `*int` | Optional | - |
| `minimumPrice` | `*int` | Optional | - |
| `deletedAt` | `*time.Time` | Optional | - |

## Example (as JSON)

```json
{
  "id": null,
  "name": null,
  "description": null,
  "url": null,
  "statement_descriptor": null,
  "interval": null,
  "interval_count": null,
  "billing_type": null,
  "payment_methods": null,
  "installments": null,
  "status": null,
  "currency": null,
  "created_at": null,
  "updated_at": null,
  "items": null,
  "billing_days": null,
  "shippable": null,
  "metadata": null,
  "trial_period_days": null,
  "minimum_price": null,
  "deleted_at": null
}
```

