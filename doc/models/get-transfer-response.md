
# Get Transfer Response

Transfer response

## Structure

`GetTransferResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | Id |
| `amount` | `*int` | Optional | Transfer amount |
| `status` | `*string` | Optional | Transfer status |
| `createdAt` | `*time.Time` | Optional | Transfer creation date |
| `updatedAt` | `*time.Time` | Optional | Transfer last update date |
| `bankAccount` | [`*models.GetBankAccountResponse`](../../doc/models/get-bank-account-response.md) | Optional | Bank account |
| `metadata` | `map[string]*string` | Optional | Metadata |

## Example (as JSON)

```json
{
  "id": null,
  "amount": null,
  "status": null,
  "created_at": null,
  "updated_at": null,
  "bank_account": null,
  "metadata": null
}
```

