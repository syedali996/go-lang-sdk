
# Get Checkout Payment Settings Response

Checkout Payment Settings Response

## Structure

`GetCheckoutPaymentSettingsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `successUrl` | `*string` | Optional | Success Url |
| `paymentUrl` | `*string` | Optional | Payment Url |
| `acceptedPaymentMethods` | `*[]string` | Optional | Accepted Payment Methods |
| `status` | `*string` | Optional | Status |
| `customer` | [`*models.GetCustomerResponse`](../../doc/models/get-customer-response.md) | Optional | Customer |
| `amount` | `*int` | Optional | Payment amount |
| `defaultPaymentMethod` | `*string` | Optional | Default Payment Method |
| `gatewayAffiliationId` | `*string` | Optional | Gateway Affiliation Id |

## Example (as JSON)

```json
{
  "success_url": null,
  "payment_url": null,
  "accepted_payment_methods": null,
  "status": null,
  "customer": null,
  "amount": null,
  "default_payment_method": null,
  "gateway_affiliation_id": null
}
```

