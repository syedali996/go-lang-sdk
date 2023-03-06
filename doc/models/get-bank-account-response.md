
# Get Bank Account Response

## Structure

`GetBankAccountResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | Id |
| `holderName` | `*string` | Optional | Holder name |
| `holderType` | `*string` | Optional | Holder type |
| `bank` | `*string` | Optional | Bank |
| `branchNumber` | `*string` | Optional | Branch number |
| `branchCheckDigit` | `*string` | Optional | Branch check digit |
| `accountNumber` | `*string` | Optional | Account number |
| `accountCheckDigit` | `*string` | Optional | Account check digit |
| `mType` | `*string` | Optional | Bank account type |
| `status` | `*string` | Optional | Bank account status |
| `createdAt` | `*time.Time` | Optional | Creation date |
| `updatedAt` | `*time.Time` | Optional | Last update date |
| `deletedAt` | `*time.Time` | Optional | Deletion date |
| `recipient` | [`*models.GetRecipientResponse`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `metadata` | `map[string]*string` | Optional | Metadata |
| `pixKey` | `*string` | Optional | Pix Key |

## Example (as JSON)

```json
{
  "id": null,
  "holder_name": null,
  "holder_type": null,
  "bank": null,
  "branch_number": null,
  "branch_check_digit": null,
  "account_number": null,
  "account_check_digit": null,
  "type": null,
  "status": null,
  "created_at": null,
  "updated_at": null,
  "deleted_at": null,
  "recipient": null,
  "metadata": null,
  "pix_key": null
}
```

