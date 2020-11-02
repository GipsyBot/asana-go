package asana

// Resource represents a generic Asana Resource,
//  containing a globally unique identifier.
type Resource struct {
	// Read-only. Globally unique ID of the object
	ID           string `json:"gid,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
	Name         string `json:"name"`
}
