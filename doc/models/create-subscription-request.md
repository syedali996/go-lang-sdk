
# Create Subscription Request

Request for creating a subcription

## Structure

`CreateSubscriptionRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customer` | [`models.CreateCustomerRequest`](../../doc/models/create-customer-request.md) | Required | Customer |
| `card` | [`models.CreateCardRequest`](../../doc/models/create-card-request.md) | Required | Card |
| `code` | `string` | Required | Subscription code |
| `paymentMethod` | `string` | Required | Payment method |
| `billingType` | `string` | Required | Billing type |
| `statementDescriptor` | `string` | Required | Statement descriptor for credit card subscriptions |
| `description` | `string` | Required | Subscription description |
| `currency` | `string` | Required | Currency |
| `interval` | `string` | Required | Interval |
| `intervalCount` | `int` | Required | Interval count |
| `pricingScheme` | [`models.CreatePricingSchemeRequest`](../../doc/models/create-pricing-scheme-request.md) | Required | Subscription pricing scheme |
| `items` | [`[]models.CreateSubscriptionItemRequest`](../../doc/models/create-subscription-item-request.md) | Required | Subscription items |
| `shipping` | [`models.CreateShippingRequest`](../../doc/models/create-shipping-request.md) | Required | Shipping |
| `discounts` | [`[]models.CreateDiscountRequest`](../../doc/models/create-discount-request.md) | Required | Discounts |
| `metadata` | `map[string]string` | Required | Metadata |
| `setup` | [`models.CreateSetupRequest`](../../doc/models/create-setup-request.md) | Optional | Setup data |
| `planId` | `string` | Optional | Plan id |
| `customerId` | `string` | Optional | Customer id |
| `cardId` | `string` | Optional | Card id |
| `billingDay` | `int` | Optional | Billing day |
| `installments` | `int` | Optional | Number of installments |
| `startAt` | `time.Time` | Optional | Subscription start date |
| `minimumPrice` | `int` | Optional | Subscription minimum price |
| `cycles` | `int` | Optional | Number of cycles |
| `cardToken` | `string` | Optional | Card token |
| `gatewayAffiliationId` | `string` | Optional | Gateway Affiliation code |
| `quantity` | `int` | Optional | Quantity |
| `boletoDueDays` | `int` | Optional | Days until boleto expires |
| `increments` | [`[]models.CreateIncrementRequest`](../../doc/models/create-increment-request.md) | Required | Increments |
| `period` | [`models.CreatePeriodRequest`](../../doc/models/create-period-request.md) | Optional | - |
| `submerchant` | [`models.CreateSubMerchantRequest`](../../doc/models/create-sub-merchant-request.md) | Optional | SubMerchant |
| `split` | [`models.CreateSubscriptionSplitRequest`](../../doc/models/create-subscription-split-request.md) | Optional | Subscription's split |
| `boleto` | [`models.CreateSubscriptionBoletoRequest`](../../doc/models/create-subscription-boleto-request.md) | Optional | Information about fines and interest on the "boleto" used from payment |

## Example (as JSON)

```json
{
  "customer": {
    "name": "{\n    \"name\": \"Tony Stark\"\n}",
    "email": null,
    "document": null,
    "type": null,
    "address": null,
    "metadata": null,
    "phones": null,
    "code": null
  },
  "card": {
    "number": null,
    "holder_name": null,
    "exp_month": null,
    "exp_year": null,
    "cvv": null,
    "billing_address": null,
    "brand": null,
    "billing_address_id": null,
    "metadata": null,
    "type": "credit",
    "options": null,
    "private_label": null,
    "label": null
  },
  "code": null,
  "payment_method": null,
  "billing_type": null,
  "statement_descriptor": null,
  "description": null,
  "currency": null,
  "interval": null,
  "interval_count": null,
  "pricing_scheme": null,
  "items": null,
  "shipping": null,
  "discounts": null,
  "metadata": null,
  "increments": null
}
```

