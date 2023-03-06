
# Get Voucher Transaction Response

Response for voucher transactions

## Structure

`GetVoucherTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `statementDescriptor` | `*string` | Optional | Text that will appear on the voucher's statement |
| `acquirerName` | `*string` | Optional | Acquirer name |
| `acquirerAffiliationCode` | `*string` | Optional | Acquirer affiliation code |
| `acquirerTid` | `*string` | Optional | Acquirer TID |
| `acquirerNsu` | `*string` | Optional | Acquirer NSU |
| `acquirerAuthCode` | `*string` | Optional | Acquirer authorization code |
| `acquirerMessage` | `*string` | Optional | acquirer_message |
| `acquirerReturnCode` | `*string` | Optional | Acquirer return code |
| `operationType` | `*string` | Optional | Operation type |
| `card` | [`*models.GetCardResponse`](../../doc/models/get-card-response.md) | Optional | Card data |

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
  "transaction_type": "voucher",
  "id": null,
  "gateway_response": null,
  "antifraud_response": null,
  "metadata": null,
  "split": null,
  "interest": null,
  "fine": null,
  "max_days_to_pay_past_due": null,
  "statement_descriptor": null,
  "acquirer_name": null,
  "acquirer_affiliation_code": null,
  "acquirer_tid": null,
  "acquirer_nsu": null,
  "acquirer_auth_code": null,
  "acquirer_message": null,
  "acquirer_return_code": null,
  "operation_type": null,
  "card": null
}
```

