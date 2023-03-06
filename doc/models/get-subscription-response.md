
# Get Subscription Response

## Structure

`GetSubscriptionResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `code` | `*string` | Optional | - |
| `startAt` | `*time.Time` | Optional | - |
| `interval` | `*string` | Optional | - |
| `intervalCount` | `*int` | Optional | - |
| `billingType` | `*string` | Optional | - |
| `currentCycle` | [`*models.GetPeriodResponse`](../../doc/models/get-period-response.md) | Optional | - |
| `paymentMethod` | `*string` | Optional | - |
| `currency` | `*string` | Optional | - |
| `installments` | `*int` | Optional | - |
| `status` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `updatedAt` | `*time.Time` | Optional | - |
| `customer` | [`*models.GetCustomerResponse`](../../doc/models/get-customer-response.md) | Optional | - |
| `card` | [`*models.GetCardResponse`](../../doc/models/get-card-response.md) | Optional | - |
| `items` | [`*[]models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md) | Optional | - |
| `statementDescriptor` | `*string` | Optional | - |
| `metadata` | `map[string]*string` | Optional | - |
| `setup` | [`*models.GetSetupResponse`](../../doc/models/get-setup-response.md) | Optional | - |
| `gatewayAffiliationId` | `*string` | Optional | Affiliation Code |
| `nextBillingAt` | `*time.Time` | Optional | - |
| `billingDay` | `*int` | Optional | - |
| `minimumPrice` | `*int` | Optional | - |
| `canceledAt` | `*time.Time` | Optional | - |
| `discounts` | [`*[]models.GetDiscountResponse`](../../doc/models/get-discount-response.md) | Optional | Subscription discounts |
| `increments` | [`*[]models.GetIncrementResponse`](../../doc/models/get-increment-response.md) | Optional | Subscription increments |
| `boletoDueDays` | `*int` | Optional | Days until boleto expires |
| `split` | [`*models.GetSubscriptionSplitResponse`](../../doc/models/get-subscription-split-response.md) | Optional | Subscription's split response |
| `boleto` | [`*models.GetSubscriptionBoletoResponse`](../../doc/models/get-subscription-boleto-response.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": null,
  "code": null,
  "start_at": null,
  "interval": null,
  "interval_count": null,
  "billing_type": null,
  "current_cycle": null,
  "payment_method": null,
  "currency": null,
  "installments": null,
  "status": null,
  "created_at": null,
  "updated_at": null,
  "customer": null,
  "card": null,
  "items": null,
  "statement_descriptor": null,
  "metadata": null,
  "setup": null,
  "gateway_affiliation_id": null,
  "next_billing_at": null,
  "billing_day": null,
  "minimum_price": null,
  "canceled_at": null,
  "discounts": null,
  "increments": null,
  "boleto_due_days": null,
  "split": null,
  "boleto": null
}
```

