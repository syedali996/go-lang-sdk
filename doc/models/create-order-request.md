
# Create Order Request

Request for creating an order

## Structure

`CreateOrderRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `items` | [`[]models.CreateOrderItemRequest`](../../doc/models/create-order-item-request.md) | Required | Items |
| `customer` | [`models.CreateCustomerRequest`](../../doc/models/create-customer-request.md) | Required | Customer |
| `payments` | [`[]models.CreatePaymentRequest`](../../doc/models/create-payment-request.md) | Required | Payment data |
| `code` | `string` | Required | The order code |
| `customerId` | `string` | Required | The customer id |
| `shipping` | [`models.CreateShippingRequest`](../../doc/models/create-shipping-request.md) | Optional | Shipping data |
| `metadata` | `map[string]string` | Required | Metadata |
| `antifraudEnabled` | `bool` | Optional | Defines whether the order will go through anti-fraud |
| `ip` | `string` | Optional | Ip address |
| `sessionId` | `string` | Optional | Session id |
| `location` | [`models.CreateLocationRequest`](../../doc/models/create-location-request.md) | Optional | Request's location |
| `device` | [`models.CreateDeviceRequest`](../../doc/models/create-device-request.md) | Optional | Device's informations |
| `closed` | `bool` | Required | **Default**: `true` |
| `currency` | `string` | Optional | Currency |
| `antifraud` | [`models.CreateAntifraudRequest`](../../doc/models/create-antifraud-request.md) | Optional | - |
| `submerchant` | [`models.CreateSubMerchantRequest`](../../doc/models/create-sub-merchant-request.md) | Optional | SubMerchant |

## Example (as JSON)

```json
{
  "items": null,
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
  "payments": null,
  "code": null,
  "customer_id": null,
  "metadata": null,
  "closed": true
}
```

