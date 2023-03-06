
# Get Plan Item Response

Response object for getting a plan item

## Structure

`GetPlanItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `name` | `*string` | Optional | - |
| `status` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `updatedAt` | `*time.Time` | Optional | - |
| `pricingScheme` | [`*models.GetPricingSchemeResponse`](../../doc/models/get-pricing-scheme-response.md) | Optional | - |
| `description` | `*string` | Optional | - |
| `plan` | [`*models.GetPlanResponse`](../../doc/models/get-plan-response.md) | Optional | - |
| `quantity` | `*int` | Optional | - |
| `cycles` | `*int` | Optional | - |
| `deletedAt` | `*time.Time` | Optional | - |

## Example (as JSON)

```json
{
  "id": null,
  "name": null,
  "status": null,
  "created_at": null,
  "updated_at": null,
  "pricing_scheme": null,
  "description": null,
  "plan": null,
  "quantity": null,
  "cycles": null,
  "deleted_at": null
}
```

