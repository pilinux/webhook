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

// ProcessEventBalance processes the incoming event and binds the raw data to a stripe.Balance struct.
/*
- https://docs.stripe.com/api/balance/balance_object

- `balance.available`
*/
func ProcessEventBalance(event stripe.Event) (balance stripe.Balance, err error) {
	switch event.Type {
	case "balance.available":
		err = json.Unmarshal(event.Data.Raw, &balance)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
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

// ProcessEventCheckoutSession processes the incoming event and binds the raw data to a stripe.CheckoutSession struct.
/*
- https://docs.stripe.com/api/checkout/sessions/object

- `checkout.session.async_payment_failed`

- `checkout.session.async_payment_succeeded`

- `checkout.session.completed`

- `checkout.session.expired`
*/
func ProcessEventCheckoutSession(event stripe.Event) (checkoutSession stripe.CheckoutSession, err error) {
	switch event.Type {
	case
		"checkout.session.async_payment_failed",
		"checkout.session.async_payment_succeeded",
		"checkout.session.completed",
		"checkout.session.expired":
		err = json.Unmarshal(event.Data.Raw, &checkoutSession)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventCoupon processes the incoming event and binds the raw data to a stripe.Coupon struct.
/*
- https://docs.stripe.com/api/coupons/object

- `coupon.created`

- `coupon.deleted`

- `coupon.updated`
*/
func ProcessEventCoupon(event stripe.Event) (coupon stripe.Coupon, err error) {
	switch event.Type {
	case
		"coupon.created",
		"coupon.deleted",
		"coupon.updated":
		err = json.Unmarshal(event.Data.Raw, &coupon)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventCreditNote processes the incoming event and binds the raw data to a stripe.CreditNote struct.
/*
- https://docs.stripe.com/api/credit_notes/object

- `credit_note.created`

- `credit_note.updated`

- `credit_note.voided`
*/
func ProcessEventCreditNote(event stripe.Event) (creditNote stripe.CreditNote, err error) {
	switch event.Type {
	case
		"credit_note.created",
		"credit_note.updated",
		"credit_note.voided":
		err = json.Unmarshal(event.Data.Raw, &creditNote)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventCustomer processes the incoming event and binds the raw data to a stripe.Customer struct.
/*
- https://docs.stripe.com/api/customers/object

- `customer.created`

- `customer.updated`

- `customer.deleted`
*/
func ProcessEventCustomer(event stripe.Event) (customer stripe.Customer, err error) {
	switch event.Type {
	case
		"customer.created",
		"customer.updated",
		"customer.deleted":
		err = json.Unmarshal(event.Data.Raw, &customer)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventCustomerDiscount processes the incoming event and binds the raw data to a stripe.Discount struct.
/*
- https://docs.stripe.com/api/discounts/object

- `customer.discount.created`

- `customer.discount.deleted`

- `customer.discount.updated`
*/
func ProcessEventCustomerDiscount(event stripe.Event) (discount stripe.Discount, err error) {
	switch event.Type {
	case
		"customer.discount.created",
		"customer.discount.deleted",
		"customer.discount.updated":
		err = json.Unmarshal(event.Data.Raw, &discount)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventCustomerSource processes the incoming event and binds the raw data to a stripe.Source struct.
/*
- https://docs.stripe.com/api/sources/object

- `customer.source.created`

- `customer.source.deleted`

- `customer.source.expiring`

- `customer.source.updated`
*/
func ProcessEventCustomerSource(event stripe.Event) (source stripe.Source, err error) {
	switch event.Type {
	case
		"customer.source.created",
		"customer.source.deleted",
		"customer.source.expiring",
		"customer.source.updated":
		err = json.Unmarshal(event.Data.Raw, &source)
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

// ProcessEventCustomerTaxID processes the incoming event and binds the raw data to a stripe.TaxID struct.
/*
- https://docs.stripe.com/api/tax_ids/object

- `customer.tax_id.created`

- `customer.tax_id.deleted`

- `customer.tax_id.updated`
*/
func ProcessEventCustomerTaxID(event stripe.Event) (taxID stripe.TaxID, err error) {
	switch event.Type {
	case
		"customer.tax_id.created",
		"customer.tax_id.deleted",
		"customer.tax_id.updated":
		err = json.Unmarshal(event.Data.Raw, &taxID)
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

// ProcessEventMandate processes the incoming event and binds the raw data to a stripe.Mandate struct.
/*
- https://docs.stripe.com/api/mandates/object

- `mandate.updated`
*/
func ProcessEventMandate(event stripe.Event) (mandate stripe.Mandate, err error) {
	switch event.Type {
	case "mandate.updated":
		err = json.Unmarshal(event.Data.Raw, &mandate)
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

// ProcessEventPaymentLink processes the incoming event and binds the raw data to a stripe.PaymentLink struct.
/*
- https://docs.stripe.com/api/payment-link/object

- `payment_link.created`

- `payment_link.updated`
*/
func ProcessEventPaymentLink(event stripe.Event) (paymentLink stripe.PaymentLink, err error) {
	switch event.Type {
	case
		"payment_link.created",
		"payment_link.updated":
		err = json.Unmarshal(event.Data.Raw, &paymentLink)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventPaymentMethod processes the incoming event and binds the raw data to a stripe.PaymentMethod struct.
/*
- https://docs.stripe.com/api/payment_methods/object

- `payment_method.attached`

- `payment_method.automatically_updated`

- `payment_method.detached`

- `payment_method.updated`
*/
func ProcessEventPaymentMethod(event stripe.Event) (paymentMethod stripe.PaymentMethod, err error) {
	switch event.Type {
	case
		"payment_method.attached",
		"payment_method.automatically_updated",
		"payment_method.detached",
		"payment_method.updated":
		err = json.Unmarshal(event.Data.Raw, &paymentMethod)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventPlan processes the incoming event and binds the raw data to a stripe.Plan struct.
/*
- https://docs.stripe.com/api/plans/object

- `plan.created`

- `plan.deleted`

- `plan.updated`
*/
func ProcessEventPlan(event stripe.Event) (plan stripe.Plan, err error) {
	switch event.Type {
	case
		"plan.created",
		"plan.deleted",
		"plan.updated":
		err = json.Unmarshal(event.Data.Raw, &plan)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventPrice processes the incoming event and binds the raw data to a stripe.Price struct.
/*
- https://docs.stripe.com/api/prices/object

- `price.created`

- `price.deleted`

- `price.updated`
*/
func ProcessEventPrice(event stripe.Event) (price stripe.Price, err error) {
	switch event.Type {
	case
		"price.created",
		"price.deleted",
		"price.updated":
		err = json.Unmarshal(event.Data.Raw, &price)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventProduct processes the incoming event and binds the raw data to a stripe.Product struct.
/*
- https://docs.stripe.com/api/products/object

- `product.created`

- `product.deleted`

- `product.updated`
*/
func ProcessEventProduct(event stripe.Event) (product stripe.Product, err error) {
	switch event.Type {
	case
		"product.created",
		"product.deleted",
		"product.updated":
		err = json.Unmarshal(event.Data.Raw, &product)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventPromotionCode processes the incoming event and binds the raw data to a stripe.PromotionCode struct.
/*
- https://docs.stripe.com/api/promotion_codes/object

- `promotion_code.created`

- `promotion_code.updated`
*/
func ProcessEventPromotionCode(event stripe.Event) (promotionCode stripe.PromotionCode, err error) {
	switch event.Type {
	case
		"promotion_code.created",
		"promotion_code.updated":
		err = json.Unmarshal(event.Data.Raw, &promotionCode)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventQuote processes the incoming event and binds the raw data to a stripe.Quote struct.
/*
- https://docs.stripe.com/api/quotes/object

- `quote.accepted`

- `quote.canceled`

- `quote.created`

- `quote.finalized`

- `quote.will_expire`
*/
func ProcessEventQuote(event stripe.Event) (quote stripe.Quote, err error) {
	switch event.Type {
	case
		"quote.accepted",
		"quote.canceled",
		"quote.created",
		"quote.finalized",
		"quote.will_expire":
		err = json.Unmarshal(event.Data.Raw, &quote)
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

// ProcessEventSubscriptionSchedule processes the incoming event and binds the raw data to a stripe.SubscriptionSchedule struct.
/*
- https://docs.stripe.com/api/subscription_schedules/object

- `subscription_schedule.aborted`

- `subscription_schedule.canceled`

- `subscription_schedule.completed`

- `subscription_schedule.created`

- `subscription_schedule.expiring`

- `subscription_schedule.released`

- `subscription_schedule.updated`
*/
func ProcessEventSubscriptionSchedule(event stripe.Event) (subscriptionSchedule stripe.SubscriptionSchedule, err error) {
	switch event.Type {
	case
		"subscription_schedule.aborted",
		"subscription_schedule.canceled",
		"subscription_schedule.completed",
		"subscription_schedule.created",
		"subscription_schedule.expiring",
		"subscription_schedule.released",
		"subscription_schedule.updated":
		err = json.Unmarshal(event.Data.Raw, &subscriptionSchedule)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventTaxRate processes the incoming event and binds the raw data to a stripe.TaxRate struct.
/*
- https://docs.stripe.com/api/tax_rates/object

- `tax_rate.created`

- `tax_rate.updated`
*/
func ProcessEventTaxRate(event stripe.Event) (taxRate stripe.TaxRate, err error) {
	switch event.Type {
	case
		"tax_rate.created",
		"tax_rate.updated":
		err = json.Unmarshal(event.Data.Raw, &taxRate)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}

// ProcessEventTaxSettings processes the incoming event and binds the raw data to a stripe.TaxSettings struct.
/*
- https://docs.stripe.com/api/tax/settings/object

- `tax.settings.updated`
*/
func ProcessEventTaxSettings(event stripe.Event) (taxSettings stripe.TaxSettings, err error) {
	switch event.Type {
	case "tax.settings.updated":
		err = json.Unmarshal(event.Data.Raw, &taxSettings)
	default:
		err = fmt.Errorf("unhandled event type: %s", event.Type)
	}
	return
}
