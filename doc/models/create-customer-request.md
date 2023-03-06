
# Create Customer Request

Request for creating a new customer

## Structure

`CreateCustomerRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `name` | `string` | Required | Name |
| `email` | `string` | Required | Email |
| `document` | `string` | Required | Document number. Only numbers, no special characters. |
| `mType` | `string` | Required | Person type. Can be either 'individual' or 'company' |
| `address` | [`models.CreateAddressRequest`](../../doc/models/create-address-request.md) | Required | The customer's address |
| `metadata` | `map[string]string` | Required | Metadata |
| `phones` | [`models.CreatePhonesRequest`](../../doc/models/create-phones-request.md) | Required | - |
| `code` | `string` | Required | Customer code |
| `gender` | `string` | Optional | Customer Gender |
| `documentType` | `string` | Optional | - |

## Example (as JSON)

```json
{
  "name": "{\n    \"name\": \"Tony Stark\"\n}",
  "email": null,
  "document": null,
  "type": null,
  "address": null,
  "metadata": null,
  "phones": null,
  "code": null
}
```

