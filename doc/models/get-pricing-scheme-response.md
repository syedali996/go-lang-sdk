
# Get Pricing Scheme Response

Response object for getting a pricing scheme

## Structure

`GetPricingSchemeResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `price` | `*int` | Optional | - |
| `schemeType` | `*string` | Optional | - |
| `priceBrackets` | [`*[]models.GetPriceBracketResponse`](../../doc/models/get-price-bracket-response.md) | Optional | - |
| `minimumPrice` | `*int` | Optional | - |
| `percentage` | `*float64` | Optional | percentual value used in pricing_scheme Percent |

## Example (as JSON)

```json
{
  "price": null,
  "scheme_type": null,
  "price_brackets": null,
  "minimum_price": null,
  "percentage": null
}
```

