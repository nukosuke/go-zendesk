package zendesk

import (
	"context"
	"encoding/json"
	"time"
)

// Configuration is a dictionary of custom configuration fields
type Configuration map[string]interface{}

// CustomRole is zendesk CustomRole JSON payload format
// https://developer.zendesk.com/api-reference/ticketing/account-configuration/custom_roles/
type CustomRole struct {
	Description     string        `json:"description,omitempty"`
	ID              int64         `json:"id,omitempty"`
	TeamMemberCount int64         `json:"team_member_count"`
	Name            string        `json:"name"`
	Configuration   Configuration `json:"configuration"`
	RoleType        int64         `json:"role_type"`
	CreatedAt       time.Time     `json:"created_at,omitempty"`
	UpdatedAt       time.Time     `json:"updated_at,omitempty"`
}

// CustomRoleAPI an interface containing all CustomRole related methods
type CustomRoleAPI interface {
	GetCustomRoles(ctx context.Context) ([]CustomRole, error)
}

// GetRoles fetch CustomRoles list
func (z *Client) GetCustomRoles(ctx context.Context) ([]CustomRole, error) {
	var data struct {
		CustomRoles []CustomRole `json:"custom_roles"`
		Page
	}

	u := "/custom_roles.json"

	body, err := z.get(ctx, u)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data.CustomRoles, nil
}
