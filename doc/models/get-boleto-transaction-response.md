
# Get Boleto Transaction Response

Response object for getting a boleto transaction

## Structure

`GetBoletoTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `url` | `*string` | Optional | - |
| `barcode` | `*string` | Optional | - |
| `nossoNumero` | `*string` | Optional | - |
| `bank` | `*string` | Optional | - |
| `documentNumber` | `*string` | Optional | - |
| `instructions` | `*string` | Optional | - |
| `billingAddress` | [`*models.GetBillingAddressResponse`](../../doc/models/get-billing-address-response.md) | Optional | - |
| `dueAt` | `*time.Time` | Optional | - |
| `qrCode` | `*string` | Optional | - |
| `line` | `*string` | Optional | - |
| `pdfPassword` | `*string` | Optional | - |
| `pdf` | `*string` | Optional | - |
| `paidAt` | `*time.Time` | Optional | - |
| `paidAmount` | `*string` | Optional | - |
| `mType` | `*string` | Optional | - |
| `creditAt` | `*time.Time` | Optional | - |
| `statementDescriptor` | `*string` | Optional | Soft Descriptor |

## Example (as JSON)

```json
{
  "gateway_id": null,
  "amount": null,
  "status": null,
  "success": null,
  "created_at": null,
  "updated_at": null,
  "attempt_count": null,
  "max_attempts": null,
  "splits": null,
  "next_attempt": null,
  "transaction_type": "boleto",
  "id": null,
  "gateway_response": null,
  "antifraud_response": null,
  "metadata": null,
  "split": null,
  "interest": null,
  "fine": null,
  "max_days_to_pay_past_due": null,
  "url": null,
  "barcode": null,
  "nosso_numero": null,
  "bank": null,
  "document_number": null,
  "instructions": null,
  "billing_address": null,
  "due_at": null,
  "qr_code": null,
  "line": null,
  "pdf_password": null,
  "pdf": null,
  "paid_at": null,
  "paid_amount": null,
  "type": null,
  "credit_at": null,
  "statement_descriptor": null
}
```

