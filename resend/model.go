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
- contact.created
- contact.updated
- contact.deleted
*/
const (
	EmailSent            EventType = "email.sent"
	EmailDelivered       EventType = "email.delivered"
	EmailDeliveryDelayed EventType = "email.delivery_delayed"
	EmailComplained      EventType = "email.complained"
	EmailBounced         EventType = "email.bounced"
	EmailOpened          EventType = "email.opened"
	EmailClicked         EventType = "email.clicked"
	ContactCreated       EventType = "contact.created"
	ContactUpdated       EventType = "contact.updated"
	ContactDeleted       EventType = "contact.deleted"
)

// Payload struct to process the incoming request over webhook
type Payload struct {
	Type      EventType `json:"type"`
	CreatedAt string    `json:"created_at"`
	Data      Data      `json:"data"`
}

// Data struct
type Data struct {
	// for email events
	CreatedAt string   `json:"created_at,omitempty"`
	EmailID   string   `json:"email_id,omitempty"`
	From      string   `json:"from,omitempty"`
	To        []string `json:"to,omitempty"`
	Click     *Click   `json:"click,omitempty"`
	Subject   string   `json:"subject,omitempty"`

	// for contact events
	ID           string `json:"id,omitempty"`
	AudienceID   string `json:"audience_id,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	Email        string `json:"email,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Unsubscribed bool   `json:"unsubscribed,omitempty"`
}

// Click struct
type Click struct {
	IPAddress string `json:"ipAddress,omitempty"`
	Link      string `json:"link,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	UserAgent string `json:"userAgent,omitempty"`
}
