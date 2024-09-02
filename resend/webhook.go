// Package resend provides utilities for processing incoming requests from resend.com over webhook.
//
// https://resend.com/docs/dashboard/webhooks/event-types
package resend

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pilinux/webhook/svixgo"
	svix "github.com/svix/svix-webhooks/go"
)

// HandleRequest validates the incoming payload against the svix signature headers
// using the webhook signing secret and binds the raw data to a payload struct.
func HandleRequest(r *http.Request, wh *svix.Webhook) (payload Payload, err error) {
	headers := r.Header
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	// validate the data
	err = svixgo.Verify(wh, body, headers)
	if err != nil {
		return
	}

	// bind raw data to payload struct
	err = json.Unmarshal(body, &payload)
	return
}
