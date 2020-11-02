package asana

import (
	"time"
)

// Change represents Information about the type of change that has occurred.
// This field is only present when the value of the property action,
// describing the action taken on the resource, is "changed"
type Change struct {
	Action       string      `json:"action"`
	AddedValue   interface{} `json:"added_value"`
	NewValue     interface{} `json:"new_value"`
	RemovedValue interface{} `json:"removed_value"`
	Field        string      `json:"field"`
}

// Event represents a change to a resource that was observed
// by an event subscription or delivered asynchronously
// to the target location of an active webhook
type Event struct {
	Action    string     `json:"action"`
	Change    Change     `json:"change"`
	Parent    Resource   `json:"parent"`
	Resource  Resource   `json:"resource"`
	User      Resource   `json:"user"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// Events represents the data received in a webhook payload
type Events struct {
	Events []Event `json:"events"`
}
