
# Create Token Request

Token data

## Structure

`CreateTokenRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mType` | `string` | Required | Token type<br>**Default**: `"card"` |
| `card` | [`models.CreateCardTokenRequest`](../../doc/models/create-card-token-request.md) | Required | Card data |

## Example (as JSON)

```json
{
  "type": "card",
  "card": null
}
```

