package asana

import (
	"fmt"
)

// UserTaskList represents the tasks assigned to a particular user
type UserTaskList struct {
	// Read-only. Globally unique ID of the object
	ID           string     `json:"gid,omitempty"`
	ResourceType string     `json:"resource_type,omitempty"`
	Name         string     `json:"name"`
	Owner        *User      `json:"owner"`
	Workspace    *Workspace `json:"workspace"`
}

// GetUserTaskList returns the full record for a user's task list.
func (c *Client) GetUserTaskList(id, workspaceID string) (*UserTaskList, error) {
	c.info("Creating webhook")

	result := &UserTaskList{}

	_, err := c.get(fmt.Sprintf("/users/%s/user_task_list?workspace=%s", id, workspaceID), nil, result)
	return result, err
}
