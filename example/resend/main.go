// Package main - example implementation of the resend package
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/pilinux/webhook/resend"
	"github.com/pilinux/webhook/svixgo"
)

func main() {
	// load the webhook secret from the environment
	secret := os.Getenv("WEBHOOK_SECRET")
	secret = strings.TrimSpace(secret)
	if secret == "" {
		fmt.Println("missing webhook secret")
		return
	}

	// create a new webhook instance
	wh, err := svixgo.NewWebhook(secret)
	if err != nil {
		fmt.Println("error creating webhook instance:", err)
		return
	}

	// handle incoming request
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		payload, err := resend.HandleRequest(r, wh)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// do something with the payload
		fmt.Println("event type:", payload.Type)
		fmt.Println("created at:", payload.CreatedAt)
		fmt.Println("email id:", payload.Data.EmailID)
		fmt.Println("from:", payload.Data.From)
		fmt.Println("to:", payload.Data.To)
		fmt.Println("subject:", payload.Data.Subject)
		fmt.Println("=====================================")

		w.WriteHeader(http.StatusNoContent)
	})

	// start the server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error starting server:", err)
		return
	}
}
