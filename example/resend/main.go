// Package main - example implementation of the resend package
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

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
		// print the request method
		fmt.Println("=====================================")
		fmt.Println("time:", time.Now().Format(time.RFC3339))
		fmt.Println("method:", r.Method)

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// process the incoming request
		payload, err := resend.HandleRequest(r, wh)
		if err != nil {
			fmt.Println("error processing request:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// do something with the payload
		fmt.Println("event type:", payload.Type)
		fmt.Println("created at:", payload.CreatedAt)
		t, err := time.Parse(time.RFC3339, payload.CreatedAt)
		if err == nil {
			fmt.Println("create at (Unix epoch timestamp):", t.Unix())
		}
		fmt.Println("email id:", payload.Data.EmailID)
		fmt.Println("from:", payload.Data.From)
		fmt.Println("to:", payload.Data.To)
		fmt.Println("subject:", payload.Data.Subject)
		if payload.Data.Click != nil {
			fmt.Println("click:")
			fmt.Println("  ip address:", payload.Data.Click.IPAddress)
			fmt.Println("  link:", payload.Data.Click.Link)
			fmt.Println("  timestamp:", payload.Data.Click.Timestamp)
			fmt.Println("  user agent:", payload.Data.Click.UserAgent)
		}
		fmt.Println("=====================================")

		// send a response
		w.WriteHeader(http.StatusOK)
	})

	// start the server
	fmt.Println("starting server at:", time.Now().Format(time.RFC3339))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error starting server:", err)
		return
	}
}
