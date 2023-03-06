
# Get Recipient Response

Recipient response

## Structure

`GetRecipientResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | Id |
| `name` | `*string` | Optional | Name |
| `email` | `*string` | Optional | Email |
| `document` | `*string` | Optional | Document |
| `description` | `*string` | Optional | Description |
| `mType` | `*string` | Optional | Type |
| `status` | `*string` | Optional | Status |
| `createdAt` | `*time.Time` | Optional | Creation date |
| `updatedAt` | `*time.Time` | Optional | Last update date |
| `deletedAt` | `*time.Time` | Optional | Deletion date |
| `defaultBankAccount` | [`*models.GetBankAccountResponse`](../../doc/models/get-bank-account-response.md) | Optional | Default bank account |
| `gatewayRecipients` | [`*[]models.GetGatewayRecipientResponse`](../../doc/models/get-gateway-recipient-response.md) | Optional | Info about the recipient on the gateway |
| `metadata` | `map[string]*string` | Optional | Metadata |
| `automaticAnticipationSettings` | [`*models.GetAutomaticAnticipationResponse`](../../doc/models/get-automatic-anticipation-response.md) | Optional | - |
| `transferSettings` | [`*models.GetTransferSettingsResponse`](../../doc/models/get-transfer-settings-response.md) | Optional | - |
| `code` | `*string` | Optional | Recipient code |
| `paymentMode` | `*string` | Optional | Payment mode<br>**Default**: `"bank_transfer"` |

## Example (as JSON)

```json
{
  "id": null,
  "name": null,
  "email": null,
  "document": null,
  "description": null,
  "type": null,
  "status": null,
  "created_at": null,
  "updated_at": null,
  "deleted_at": null,
  "default_bank_account": null,
  "gateway_recipients": null,
  "metadata": null,
  "automatic_anticipation_settings": null,
  "transfer_settings": null,
  "code": null,
  "payment_mode": null
}
```

