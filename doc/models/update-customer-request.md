
# Update Customer Request

Request for updating a customer

## Structure

`UpdateCustomerRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `name` | `string` | Optional | Name |
| `email` | `string` | Optional | Email |
| `document` | `string` | Optional | Document number |
| `mType` | `string` | Optional | Person type |
| `address` | [`models.CreateAddressRequest`](../../doc/models/create-address-request.md) | Optional | Address |
| `metadata` | `map[string]string` | Optional | Metadata |
| `phones` | [`models.CreatePhonesRequest`](../../doc/models/create-phones-request.md) | Optional | - |
| `code` | `string` | Optional | Código de referência do cliente no sistema da loja. Max: 52 caracteres |
| `gender` | `string` | Optional | Gênero do cliente |
| `documentType` | `string` | Optional | - |

## Example (as JSON)

```json
{
  "name": null,
  "email": null,
  "document": null,
  "type": null,
  "address": null,
  "metadata": null,
  "phones": null,
  "code": null,
  "gender": null,
  "document_type": null
}
```

