// Package main - example implementation of the stripe webhook package.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	wh "github.com/pilinux/webhook/stripe"
	"github.com/stripe/stripe-go/v79"
)

func main() {
	// load the webhook secret from the environment
	secret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	secret = strings.TrimSpace(secret)
	if secret == "" {
		fmt.Println("missing webhook secret")
		return
	}

	// handle incoming request
	http.HandleFunc("/stripe_webhooks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("=====================================")
		fmt.Println("time:", time.Now().Format(time.RFC3339))

		// process the incoming request
		event, statusCode, err := wh.HandleRequest(w, r, secret)
		if err != nil {
			fmt.Println("error processing request:", err)
			w.WriteHeader(statusCode)
			return
		}

		// process the event in a separate goroutine
		go func(e stripe.Event) {
			fmt.Println("event type:", e.Type)
			fmt.Println("event id:", e.ID)

			// event.Type contains the prefix "charge."
			if strings.HasPrefix(string(e.Type), "charge.") {
				charge, err := wh.ProcessEventCharge(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("charge: %+v\n", charge)
			}

			// event.Type contains the prefix "customer.subscription."
			if strings.HasPrefix(string(e.Type), "customer.subscription.") {
				subscription, err := wh.ProcessEventCustomerSubscription(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("subscription: %+v\n", subscription)
			}

			// event.Type contains the prefix "payment_intent."
			if strings.HasPrefix(string(e.Type), "payment_intent.") {
				paymentIntent, err := wh.ProcessEventPaymentIntent(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("payment intent: %+v\n", paymentIntent)
			}
		}(event)

		// immediately respond with a 200 status code to stripe
		w.WriteHeader(http.StatusOK)
	})

	// start the server
	fmt.Println("starting server at:", time.Now().Format(time.RFC3339))
	err := http.ListenAndServe(":4242", nil)
	if err != nil {
		fmt.Println("error starting server:", err)
		return
	}
}
