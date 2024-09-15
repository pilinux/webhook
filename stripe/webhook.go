// Package stripe handles incoming requests from stripe.com over webhook.
//
// https://docs.stripe.com/webhooks
package stripe

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/webhook"
)

// HandleRequest validates the incoming payload against the Stripe signature headers
// using the webhook signing secret and binds the raw data to a stripe.Event struct.
func HandleRequest(w http.ResponseWriter, r *http.Request, secret string) (stripe.Event, int, error) {
	// stripe webhook events are always POST requests
	if r.Method != http.MethodPost {
		return stripe.Event{}, http.StatusMethodNotAllowed, fmt.Errorf("invalid request method: %s", r.Method)
	}

	// set the maximum request body size to 64KB to prevent DoS attacks
	const maxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	// get the Stripe-Signature header from the request
	sigHeader := r.Header.Get("Stripe-Signature")

	// read the entire request body into a []byte slice
	body, err := io.ReadAll(r.Body)
	if err != nil {
		if err.Error() == "http: request body too large" {
			return stripe.Event{}, http.StatusRequestEntityTooLarge, err
		}
		return stripe.Event{}, http.StatusBadRequest, err
	}

	// construct the event
	event, err := webhook.ConstructEvent(body, sigHeader, secret)
	if err != nil {
		return stripe.Event{}, http.StatusBadRequest, err
	}

	return event, http.StatusOK, nil
}

// ProcessEventCharge processes the incoming event and binds the raw data to a stripe.Charge struct.
/*
- https://docs.stripe.com/api/charges/object

- `charge.captured`

- `charge.dispute.closed`

- `charge.dispute.created`

- `charge.dispute.funds_reinstated`

- `charge.dispute.funds_withdrawn`

- `charge.dispute.updated`

- `charge.expired`

- `charge.failed`

- `charge.pending`

- `charge.refund.updated`

- `charge.refunded`

- `charge.succeeded`

- `charge.updated`
*/
func ProcessEventCharge(event stripe.Event) (charge stripe.Charge, err error) {
	switch event.Type {
	case
		"charge.captured",
		"charge.dispute.closed",
		"charge.dispute.created",
		"charge.dispute.funds_reinstated",
		"charge.dispute.funds_withdrawn",
		"charge.dispute.updated",
		"charge.expired",
		"charge.failed",
		"charge.pending",
		"charge.refund.updated",
		"charge.refunded",
		"charge.succeeded",
		"charge.updated":
		err = json.Unmarshal(event.Data.Raw, &charge)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventCustomerSubscription processes the incoming event and binds the raw data to a stripe.Subscription struct.
/*
- https://docs.stripe.com/api/subscriptions/object

- `customer.subscription.created`

- `customer.subscription.deleted`

- `customer.subscription.paused`

- `customer.subscription.pending_update_applied`

- `customer.subscription.pending_update_expired`

- `customer.subscription.resumed`

- `customer.subscription.trial_will_end`

- `customer.subscription.updated`
*/
func ProcessEventCustomerSubscription(event stripe.Event) (subscription stripe.Subscription, err error) {
	switch event.Type {
	case
		"customer.subscription.created",
		"customer.subscription.deleted",
		"customer.subscription.paused",
		"customer.subscription.pending_update_applied",
		"customer.subscription.pending_update_expired",
		"customer.subscription.resumed",
		"customer.subscription.trial_will_end",
		"customer.subscription.updated":
		err = json.Unmarshal(event.Data.Raw, &subscription)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventInvoice processes the incoming event and binds the raw data to a stripe.Invoice struct.
/*
- https://docs.stripe.com/api/invoices/object

- `invoice.created`

- `invoice.deleted`

- `invoice.finalization_failed`

- `invoice.finalized`

- `invoice.marked_uncollectible`

- `invoice.overdue`

- `invoice.paid`

- `invoice.payment_action_required`

- `invoice.payment_failed`

- `invoice.payment_succeeded`

- `invoice.sent`

- `invoice.upcoming`

- `invoice.updated`

- `invoice.voided`

- `invoice.will_be_due`
*/
func ProcessEventInvoice(event stripe.Event) (invoice stripe.Invoice, err error) {
	switch event.Type {
	case
		"invoice.created",
		"invoice.deleted",
		"invoice.finalization_failed",
		"invoice.finalized",
		"invoice.marked_uncollectible",
		"invoice.overdue",
		"invoice.paid",
		"invoice.payment_action_required",
		"invoice.payment_failed",
		"invoice.payment_succeeded",
		"invoice.sent",
		"invoice.upcoming",
		"invoice.updated",
		"invoice.voided",
		"invoice.will_be_due":
		err = json.Unmarshal(event.Data.Raw, &invoice)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventInvoiceItem processes the incoming event and binds the raw data to a stripe.InvoiceItem struct.
/*
- https://docs.stripe.com/api/invoiceitems/object

- `invoiceitem.created`

- `invoiceitem.deleted`
*/
func ProcessEventInvoiceItem(event stripe.Event) (invoiceItem stripe.InvoiceItem, err error) {
	switch event.Type {
	case
		"invoiceitem.created",
		"invoiceitem.deleted":
		err = json.Unmarshal(event.Data.Raw, &invoiceItem)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventPaymentIntent processes the incoming event and binds the raw data to a stripe.PaymentIntent struct.
/*
- https://docs.stripe.com/api/payment_intents/object

- `payment_intent.amount_capturable_updated`

- `payment_intent.canceled`

- `payment_intent.created`

- `payment_intent.partially_funded`

- `payment_intent.payment_failed`

- `payment_intent.processing`

- `payment_intent.requires_action`

- `payment_intent.succeeded`
*/
func ProcessEventPaymentIntent(event stripe.Event) (paymentIntent stripe.PaymentIntent, err error) {
	switch event.Type {
	case
		"payment_intent.amount_capturable_updated",
		"payment_intent.canceled",
		"payment_intent.created",
		"payment_intent.partially_funded",
		"payment_intent.payment_failed",
		"payment_intent.processing",
		"payment_intent.requires_action",
		"payment_intent.succeeded":
		err = json.Unmarshal(event.Data.Raw, &paymentIntent)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventSetupIntent processes the incoming event and binds the raw data to a stripe.SetupIntent struct.
/*
- https://docs.stripe.com/api/setup_intents/object

- `setup_intent.canceled`

- `setup_intent.created`

- `setup_intent.requires_action`

- `setup_intent.setup_failed`

- `setup_intent.succeeded`
*/
func ProcessEventSetupIntent(event stripe.Event) (setupIntent stripe.SetupIntent, err error) {
	switch event.Type {
	case
		"setup_intent.canceled",
		"setup_intent.created",
		"setup_intent.requires_action",
		"setup_intent.setup_failed",
		"setup_intent.succeeded":
		err = json.Unmarshal(event.Data.Raw, &setupIntent)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}
