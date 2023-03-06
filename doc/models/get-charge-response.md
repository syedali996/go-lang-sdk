
# Get Charge Response

Response object for getting a charge

## Structure

`GetChargeResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*string` | Optional | - |
| `code` | `*string` | Optional | - |
| `gatewayId` | `*string` | Optional | - |
| `amount` | `*int` | Optional | - |
| `status` | `*string` | Optional | - |
| `currency` | `*string` | Optional | - |
| `paymentMethod` | `*string` | Optional | - |
| `dueAt` | `*time.Time` | Optional | - |
| `createdAt` | `*time.Time` | Optional | - |
| `updatedAt` | `*time.Time` | Optional | - |
| `lastTransaction` | [`models.GetTransactionResponseInterface`](../../doc/models/get-transaction-response.md) | Required | - |
| `invoice` | [`*models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md) | Optional | - |
| `order` | [`*models.GetOrderResponse`](../../doc/models/get-order-response.md) | Optional | - |
| `customer` | [`*models.GetCustomerResponse`](../../doc/models/get-customer-response.md) | Optional | - |
| `metadata` | `map[string]*string` | Optional | - |
| `paidAt` | `*time.Time` | Optional | - |
| `canceledAt` | `*time.Time` | Optional | - |
| `canceledAmount` | `*int` | Optional | Canceled Amount |
| `paidAmount` | `*int` | Optional | Paid amount |
| `interestAndFinePaid` | `*int` | Optional | interest and fine paid |
| `recurrencyCycle` | `*string` | Optional | Defines whether the card has been used one or more times. |

## Example (as JSON)

```json
{
  "id": null,
  "code": null,
  "gateway_id": null,
  "amount": null,
  "status": null,
  "currency": null,
  "payment_method": null,
  "due_at": null,
  "created_at": null,
  "updated_at": null,
  "last_transaction": {
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
    "transaction_type": "transaction",
    "id": null,
    "gateway_response": null,
    "antifraud_response": null,
    "metadata": null,
    "split": null,
    "interest": null,
    "fine": null,
    "max_days_to_pay_past_due": null
  },
  "invoice": null,
  "order": null,
  "customer": null,
  "metadata": null,
  "paid_at": null,
  "canceled_at": null,
  "canceled_amount": null,
  "paid_amount": null,
  "interest_and_fine_paid": null,
  "recurrency_cycle": null
}
```

