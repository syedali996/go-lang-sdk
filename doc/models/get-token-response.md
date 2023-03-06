
# Get Token Response

Token data

## Structure

`GetTokenResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `mType` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `expiresAt` | `*string` | Optional | - |
| `card` | [`*models.GetCardTokenResponse`](../../doc/models/get-card-token-response.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": null,
  "type": null,
  "created_at": null,
  "expires_at": null,
  "card": null
}
```

