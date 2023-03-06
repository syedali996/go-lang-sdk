
# Get Invoice Response

Response object for getting an invoice

## Structure

`GetInvoiceResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `code` | `*string` | Optional | - |
| `url` | `*string` | Optional | - |
| `amount` | `*int` | Optional | - |
| `status` | `*string` | Optional | - |
| `paymentMethod` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `items` | [`*[]models.GetInvoiceItemResponse`](../../doc/models/get-invoice-item-response.md) | Optional | - |
| `customer` | [`*models.GetCustomerResponse`](../../doc/models/get-customer-response.md) | Optional | - |
| `charge` | [`*models.GetChargeResponse`](../../doc/models/get-charge-response.md) | Optional | - |
| `installments` | `*int` | Optional | - |
| `billingAddress` | [`*models.GetBillingAddressResponse`](../../doc/models/get-billing-address-response.md) | Optional | - |
| `subscription` | [`*models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md) | Optional | - |
| `cycle` | [`*models.GetPeriodResponse`](../../doc/models/get-period-response.md) | Optional | - |
| `shipping` | [`*models.GetShippingResponse`](../../doc/models/get-shipping-response.md) | Optional | - |
| `metadata` | `map[string]*string` | Optional | - |
| `dueAt` | `*time.Time` | Optional | - |
| `canceledAt` | `*time.Time` | Optional | - |
| `billingAt` | `*time.Time` | Optional | - |
| `seenAt` | `*time.Time` | Optional | - |
| `totalDiscount` | `*int` | Optional | Total discounted value |
| `totalIncrement` | `*int` | Optional | Total discounted value |
| `subscriptionId` | `*string` | Optional | Subscription Id |

## Example (as JSON)

```json
{
  "id": null,
  "code": null,
  "url": null,
  "amount": null,
  "status": null,
  "payment_method": null,
  "created_at": null,
  "items": null,
  "customer": null,
  "charge": null,
  "installments": null,
  "billing_address": null,
  "subscription": null,
  "cycle": null,
  "shipping": null,
  "metadata": null,
  "due_at": null,
  "canceled_at": null,
  "billing_at": null,
  "seen_at": null,
  "total_discount": null,
  "total_increment": null,
  "subscription_id": null
}
```

