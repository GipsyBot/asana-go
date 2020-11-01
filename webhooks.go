package asana

import (
	"fmt"
	"time"
)

// Resource represents a generic Asana Resource,
//  containing a globally unique identifier.
type Resource struct {
	// Read-only. Globally unique ID of the object
	ID           string `json:"gid,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
	Name         string `json:"name"`
}

// WebhookFilter is used to filter the types of actions
// that trigger delivery of an Event
type WebhookFilter struct {
	Action          string   `json:"action"`
	Fields          []string `json:"fields,omitempty"`
	ResourceType    string   `json:"resource_type,omitempty"`
	ResourceSubtype string   `json:"resource_subtype,omitempty"`
}

// CreateWebhookRequest represents a request to create a new Webhook
type CreateWebhookRequest struct {
	Filters  []WebhookFilter `json:"filters"`
	Target   string          `json:"target"`
	Resource string          `json:"resource"`
}

// Webhook is used to be notified of changes in Asana.
// It is  intended to provide an efficient way for integrations
// which react to changes of state in Asana to take action only when something has actually changed.
type Webhook struct {
	// Read-only. Globally unique ID of the object
	ID     string `json:"gid,omitempty"`
	Active bool   `json:"active"`
	Target string `json:"target"`

	Filters []WebhookFilter `json:"filters"`

	// Read-only. The time at which this object was created.
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	LastFailureAt      *time.Time `json:"last_failure_at,omitempty"`
	LastFailureContent string     `json:"last_failure_content,omitempty"`
	LastSuccessAt      *time.Time `json:"last_success_at,omitempty"`
}

// CreateWebhook creates a new webhook
func (c *Client) CreateWebhook(webhook *CreateWebhookRequest) (*Webhook, error) {
	c.info("Creating webhook")

	result := &Webhook{}

	err := c.post("/webhooks", webhook, result)
	return result, err
}

func (w *Webhook) Delete(client *Client) error {
	client.info("Deleting webhook %s", w.ID)

	return client.delete(fmt.Sprintf("/webhooks/%s", w.ID))
}