package zendesk

import (
	"encoding/json"
	"time"
)

// User is zendesk user JSON payload format
// https://developer.zendesk.com/rest_api/docs/support/users
type User struct {
	ID                  int64  `json:"id,omitempty"`
	URL                 string `json:"url,omitempty"`
	Email               string `json:"email,omitempty"`
	Name                string `json:"name"`
	Active              bool   `json:"active,omitempty"`
	Alias               string `json:"alias,omitempty"`
	ChatOnly            bool   `json:"chat_only,omitempty"`
	CustomRoleID        int64  `json:"custom_role_id,omitempty"`
	RoleType            int64  `json:"role_type,omitempty"`
	Details             string `json:"details,omitempty"`
	ExternalID          string `json:"external_id,omitempty"`
	Locale              string `json:"locale,omitempty"`
	LocaleID            int64  `json:"locale_id,omitempty"`
	Moderator           bool   `json:"moderator,omitempty"`
	Notes               string `json:"notes,omitempty"`
	OnlyPrivateComments bool   `json:"only_private_comments,omitempty"`
	OrganizationID      int64  `json:"organization_id,omitempty"`
	DefaultGroupID      int64  `json:"default_group_id,omitempty"`
	Phone               string `json:"phone,omitempty"`
	SharedPhoneNumber   bool   `json:"shared_phone_number,omitempty"`
	//TODO: Photo Attachment
	RestrictedAgent      bool                   `json:"restricted_agent,omitempty"`
	Role                 string                 `json:"role,omitempty"`
	Shared               bool                   `json:"shared,omitempty"`
	SharedAgent          bool                   `json:"shared_agent,omitempty"`
	Signature            string                 `json:"signature,omitempty"`
	Suspended            bool                   `json:"suspended,omitempty"`
	Tags                 []string               `json:"tags,omitempty"`
	TicketRestriction    string                 `json:"ticket_restriction,omitempty"`
	Timezone             string                 `json:"time_zone,omitempty"`
	TwoFactorAuthEnabled bool                   `json:"two_factor_auth_enabled,omitempty"`
	UserFields           map[string]interface{} `json:"user_fields,omitempty"`
	Verified             bool                   `json:"verified,omitempty"`
	ReportCSV            bool                   `json:"report_csv,omitempty"`
	LastLoginAt          time.Time              `json:"last_login_at,omitempty"`
	CreatedAt            time.Time              `json:"created_at,omitempty"`
	UpdatedAt            time.Time              `json:"updated_at,omitempty"`
}

const (
	// UserRoleEndUser end-user
	UserRoleEndUser = iota
	// UserRoleAgent agent
	UserRoleAgent
	// UserRoleAdmin admin
	UserRoleAdmin
)

var userRoleText = map[int]string{
	UserRoleEndUser: "end-user",
	UserRoleAgent:   "agent",
	UserRoleAdmin:   "admin",
}

// UserRoleText takes role type and returns role name string
func UserRoleText(role int) string {
	return userRoleText[role]
}

// GetUsers fetch user list
func (z *Client) GetUsers() ([]User, Page, error) {
	var data struct {
		Users []User `json:"users"`
		Page
	}

	body, err := z.Get("/users.json")
	if err != nil {
		return nil, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, Page{}, err
	}
	return data.Users, data.Page, nil
}

//TODO: GetUsersByGroupID, GetUsersByOrganizationID

// CreateUser creates new user
// ref: https://developer.zendesk.com/rest_api/docs/core/triggers#create-trigger
func (z Client) CreateUser(user User) (User, error) {
	var data, result struct {
		User User `json:"user"`
	}

	body, err := z.Post("/users.json", data)
	if err != nil {
		return User{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return User{}, err
	}
	return result.User, nil
}

// TODO: CreateOrUpdateManyUsers(users []User)
