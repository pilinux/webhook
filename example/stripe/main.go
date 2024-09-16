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

			// event.Type contains the prefix "balance."
			if strings.HasPrefix(string(e.Type), "balance.") {
				balance, err := wh.ProcessEventBalance(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("balance: %+v\n", balance)
			}

			// event.Type contains the prefix "charge."
			if strings.HasPrefix(string(e.Type), "charge.") {
				charge, err := wh.ProcessEventCharge(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("charge: %+v\n", charge)
			}

			// event.Type contains the prefix "checkout.session."
			if strings.HasPrefix(string(e.Type), "checkout.session.") {
				checkoutSession, err := wh.ProcessEventCheckoutSession(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("checkout session: %+v\n", checkoutSession)
			}

			// event.Type contains the prefix "coupon."
			if strings.HasPrefix(string(e.Type), "coupon.") {
				coupon, err := wh.ProcessEventCoupon(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("coupon: %+v\n", coupon)
			}

			// event.Type contains the prefix "credit_note."
			if strings.HasPrefix(string(e.Type), "credit_note.") {
				creditNote, err := wh.ProcessEventCreditNote(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("credit note: %+v\n", creditNote)
			}

			// event.Type contains the prefix "customer.subscription."
			if strings.HasPrefix(string(e.Type), "customer.subscription.") {
				subscription, err := wh.ProcessEventCustomerSubscription(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("subscription: %+v\n", subscription)
			} else if strings.HasPrefix(string(e.Type), "customer.discount.") {
				// event.Type contains the prefix "customer.discount."
				discount, err := wh.ProcessEventCustomerDiscount(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("discount: %+v\n", discount)
			} else if strings.HasPrefix(string(e.Type), "customer.source.") {
				// event.Type contains the prefix "customer.source."
				source, err := wh.ProcessEventCustomerSource(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("source: %+v\n", source)
			} else if strings.HasPrefix(string(e.Type), "customer.tax_id.") {
				// event.Type contains the prefix "customer.tax_id."
				taxID, err := wh.ProcessEventCustomerTaxID(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("tax id: %+v\n", taxID)
			} else if strings.HasPrefix(string(e.Type), "customer.") {
				// event.Type contains the prefix "customer." but not "customer.subscription."
				customer, err := wh.ProcessEventCustomer(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("customer: %+v\n", customer)
			}

			// event.Type contains the prefix "invoice."
			if strings.HasPrefix(string(e.Type), "invoice.") {
				invoice, err := wh.ProcessEventInvoice(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("invoice: %+v\n", invoice)
			}

			// event.Type contains the prefix "invoiceitem."
			if strings.HasPrefix(string(e.Type), "invoiceitem.") {
				invoiceItem, err := wh.ProcessEventInvoiceItem(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("invoice item: %+v\n", invoiceItem)
			}

			// event.Type contains the prefix "mandate."
			if strings.HasPrefix(string(e.Type), "mandate.") {
				mandate, err := wh.ProcessEventMandate(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("mandate: %+v\n", mandate)
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

			// event.Type contains the prefix "payment_link."
			if strings.HasPrefix(string(e.Type), "payment_link.") {
				paymentLink, err := wh.ProcessEventPaymentLink(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("payment link: %+v\n", paymentLink)
			}

			// event.Type contains the prefix "payment_method."
			if strings.HasPrefix(string(e.Type), "payment_method.") {
				paymentMethod, err := wh.ProcessEventPaymentMethod(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("payment method: %+v\n", paymentMethod)
			}

			// event.Type contains the prefix "plan."
			if strings.HasPrefix(string(e.Type), "plan.") {
				plan, err := wh.ProcessEventPlan(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("plan: %+v\n", plan)
			}

			// event.Type contains the prefix "price."
			if strings.HasPrefix(string(e.Type), "price.") {
				price, err := wh.ProcessEventPrice(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("price: %+v\n", price)
			}

			// event.Type contains the prefix "product."
			if strings.HasPrefix(string(e.Type), "product.") {
				product, err := wh.ProcessEventProduct(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("product: %+v\n", product)
			}

			// event.Type contains the prefix "promotion_code."
			if strings.HasPrefix(string(e.Type), "promotion_code.") {
				promotionCode, err := wh.ProcessEventPromotionCode(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("promotion code: %+v\n", promotionCode)
			}

			// event.Type contains the prefix "quote."
			if strings.HasPrefix(string(e.Type), "quote.") {
				quote, err := wh.ProcessEventQuote(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("quote: %+v\n", quote)
			}

			// event.Type contains the prefix "setup_intent."
			if strings.HasPrefix(string(e.Type), "setup_intent.") {
				setupIntent, err := wh.ProcessEventSetupIntent(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("setup intent: %+v\n", setupIntent)
			}

			// event.Type contains the prefix "subscription_schedule."
			if strings.HasPrefix(string(e.Type), "subscription_schedule.") {
				subscriptionSchedule, err := wh.ProcessEventSubscriptionSchedule(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("subscription schedule: %+v\n", subscriptionSchedule)
			}

			// event.Type contains the prefix "tax_rate."
			if strings.HasPrefix(string(e.Type), "tax_rate.") {
				taxRate, err := wh.ProcessEventTaxRate(e)
				if err != nil {
					fmt.Println("error processing event:", err)
					return
				}
				fmt.Printf("tax rate: %+v\n", taxRate)
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
