// Package svixgo is a wrapper around the github.com/svix/svix-webhooks/go package.
package svixgo

import (
	"net/http"

	svix "github.com/svix/svix-webhooks/go"
)

// NewWebhook creates a new webhook instance with the given secret.
func NewWebhook(secret string) (*svix.Webhook, error) {
	return svix.NewWebhook(secret)
}

// Verify validates the incoming payload against the svix signature headers using the webhooks signing secret.
//
/*
	wh, err := NewWebhook(secret)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = Verify(wh, payload, headers)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Do something with the message...

		w.WriteHeader(http.StatusNoContent)

	})
	http.ListenAndServe(":8080", nil)
*/
func Verify(wh *svix.Webhook, payload []byte, headers http.Header) error {
	return wh.Verify(payload, headers)
}
