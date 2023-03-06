
# Get Card Response

Response object for getting a credit card

## Structure

`GetCardResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `lastFourDigits` | `*string` | Optional | - |
| `brand` | `*string` | Optional | - |
| `holderName` | `*string` | Optional | - |
| `expMonth` | `*int` | Optional | - |
| `expYear` | `*int` | Optional | - |
| `status` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `updatedAt` | `*time.Time` | Optional | - |
| `billingAddress` | [`*models.GetBillingAddressResponse`](../../doc/models/get-billing-address-response.md) | Optional | - |
| `customer` | [`*models.GetCustomerResponse`](../../doc/models/get-customer-response.md) | Optional | - |
| `metadata` | `map[string]*string` | Optional | - |
| `mType` | `*string` | Optional | Card type |
| `holderDocument` | `*string` | Optional | Document number for the card's holder |
| `deletedAt` | `*time.Time` | Optional | - |
| `firstSixDigits` | `*string` | Optional | First six digits |
| `label` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "id": null,
  "last_four_digits": null,
  "brand": null,
  "holder_name": null,
  "exp_month": null,
  "exp_year": null,
  "status": null,
  "created_at": null,
  "updated_at": null,
  "billing_address": null,
  "customer": null,
  "metadata": null,
  "type": null,
  "holder_document": null,
  "deleted_at": null,
  "first_six_digits": null,
  "label": null
}
```

