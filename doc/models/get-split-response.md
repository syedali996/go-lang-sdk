
# Get Split Response

Split response

## Structure

`GetSplitResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mType` | `*string` | Optional | Type |
| `amount` | `*int` | Optional | Amount |
| `recipient` | [`*models.GetRecipientResponse`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `gatewayId` | `*string` | Optional | The split rule gateway id |
| `options` | [`*models.GetSplitOptionsResponse`](../../doc/models/get-split-options-response.md) | Optional | - |
| `id` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "type": null,
  "amount": null,
  "recipient": null,
  "gateway_id": null,
  "options": null,
  "id": null
}
```

