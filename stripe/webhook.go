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
