
# Get Withdraw Response

## Structure

`GetWithdrawResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `gatewayId` | `*string` | Optional | - |
| `amount` | `*int` | Optional | - |
| `status` | `*string` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `updatedAt` | `*time.Time` | Optional | - |
| `metadata` | `*[]string` | Optional | - |
| `fee` | `*int` | Optional | - |
| `fundingDate` | `*time.Time` | Optional | - |
| `fundingEstimatedDate` | `*time.Time` | Optional | - |
| `mType` | `*string` | Optional | - |
| `source` | [`*models.GetWithdrawSourceResponse`](../../doc/models/get-withdraw-source-response.md) | Optional | - |
| `target` | [`*models.GetWithdrawTargetResponse`](../../doc/models/get-withdraw-target-response.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": null,
  "gateway_id": null,
  "amount": null,
  "status": null,
  "created_at": null,
  "updated_at": null,
  "metadata": null,
  "fee": null,
  "funding_date": null,
  "funding_estimated_date": null,
  "type": null,
  "source": null,
  "target": null
}
```

