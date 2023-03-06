
# Get Address Response

Response object for getting an Address

## Structure

`GetAddressResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `street` | `*string` | Optional | - |
| `number` | `*string` | Optional | - |
| `complement` | `*string` | Optional | - |
| `zipCode` | `*string` | Optional | - |
| `neighborhood` | `*string` | Optional | - |
| `city` | `*string` | Optional | - |
| `state` | `*string` | Optional | - |
| `country` | `*string` | Optional | - |
| `status` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `updatedAt` | `*time.Time` | Optional | - |
| `customer` | [`*models.GetCustomerResponse`](../../doc/models/get-customer-response.md) | Optional | - |
| `metadata` | `map[string]*string` | Optional | - |
| `line1` | `*string` | Optional | Line 1 for address |
| `line2` | `*string` | Optional | Line 2 for address |
| `deletedAt` | `*time.Time` | Optional | - |

## Example (as JSON)

```json
{
  "id": null,
  "street": null,
  "number": null,
  "complement": null,
  "zip_code": null,
  "neighborhood": null,
  "city": null,
  "state": null,
  "country": null,
  "status": null,
  "created_at": null,
  "updated_at": null,
  "customer": null,
  "metadata": null,
  "line_1": null,
  "line_2": null,
  "deleted_at": null
}
```

