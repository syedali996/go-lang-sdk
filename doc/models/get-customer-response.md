
# Get Customer Response

Response object for getting a customer

## Structure

`GetCustomerResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `name` | `*string` | Optional | - |
| `email` | `*string` | Optional | - |
| `delinquent` | `*bool` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `updatedAt` | `*time.Time` | Optional | - |
| `document` | `*string` | Optional | - |
| `mType` | `*string` | Optional | - |
| `fbAccessToken` | `*string` | Optional | - |
| `address` | [`*models.GetAddressResponse`](../../doc/models/get-address-response.md) | Optional | - |
| `metadata` | `map[string]*string` | Optional | - |
| `phones` | [`*models.GetPhonesResponse`](../../doc/models/get-phones-response.md) | Optional | - |
| `fbId` | `*int64` | Optional | - |
| `code` | `*string` | Optional | Código de referência do cliente no sistema da loja. Max: 52 caracteres |
| `documentType` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "id": null,
  "name": null,
  "email": null,
  "delinquent": null,
  "created_at": null,
  "updated_at": null,
  "document": null,
  "type": null,
  "fb_access_token": null,
  "address": null,
  "metadata": null,
  "phones": null,
  "fb_id": null,
  "code": null,
  "document_type": null
}
```

