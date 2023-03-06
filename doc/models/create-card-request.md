
# Create Card Request

Card data

## Structure

`CreateCardRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `number` | `string` | Required | Credit card number |
| `holderName` | `string` | Required | Holder name, as written on the card |
| `expMonth` | `int` | Required | The expiration month |
| `expYear` | `int` | Required | The expiration year, that can be informed with 2 or 4 digits |
| `cvv` | `string` | Required | The card's security code |
| `billingAddress` | [`models.CreateAddressRequest`](../../doc/models/create-address-request.md) | Required | Card's billing address |
| `brand` | `string` | Required | Card brand |
| `billingAddressId` | `string` | Required | The address id for the billing address |
| `metadata` | `map[string]string` | Required | Metadata |
| `mType` | `string` | Required | Card type<br>**Default**: `"credit"` |
| `options` | [`models.CreateCardOptionsRequest`](../../doc/models/create-card-options-request.md) | Required | Options for creating the card |
| `holderDocument` | `string` | Optional | Document number for the card's holder |
| `privateLabel` | `bool` | Required | Indicates whether it is a private label card |
| `label` | `string` | Required | - |
| `id` | `string` | Optional | Identifier |
| `token` | `string` | Optional | token identifier |

## Example (as JSON)

```json
{
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
}
```

