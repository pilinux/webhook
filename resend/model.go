package resend

// EventType represents the type of event
type EventType string

/*
all event types
- email.sent
- email.delivered
- email.delivery_delayed
- email.complained
- email.bounced
- email.opened
- email.clicked
*/
const (
	EmailSent            EventType = "email.sent"
	EmailDelivered       EventType = "email.delivered"
	EmailDeliveryDelayed EventType = "email.delivery_delayed"
	EmailComplained      EventType = "email.complained"
	EmailBounced         EventType = "email.bounced"
	EmailOpened          EventType = "email.opened"
	EmailClicked         EventType = "email.clicked"
)

// Payload struct to process the incoming request over webhook
type Payload struct {
	Type      EventType `json:"type"`
	CreatedAt string    `json:"created_at"`
	Data      Data      `json:"data"`
}

// Data struct
type Data struct {
	CreatedAt string   `json:"created_at"`
	EmailID   string   `json:"email_id"`
	From      string   `json:"from"`
	To        []string `json:"to"`
	Click     *Click   `json:"click,omitempty"`
	Subject   string   `json:"subject"`
}

// Click struct
type Click struct {
	IPAddress string `json:"ipAddress"`
	Link      string `json:"link"`
	Timestamp string `json:"timestamp"`
	UserAgent string `json:"userAgent"`
}
