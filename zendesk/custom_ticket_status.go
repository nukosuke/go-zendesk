package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type CustomTicketStatus struct {
	Active                bool       `json:"active,omitempty"`
	AgentLabel            string     `json:"agent_label,omitempty"`
	CreatedAt             *time.Time `json:"created_at,omitempty"`
	Default               bool       `json:"default,omitempty"`
	Description           string     `json:"description,omitempty"`
	EndUserDescription    string     `json:"end_user_description,omitempty"`
	EndUserLabel          string     `json:"end_user_label,omitempty"`
	ID                    int64      `json:"id,omitempty"`
	RawAgentLabel         string     `json:"raw_agent_label,omitempty"`
	RawDescription        string     `json:"raw_description,omitempty"`
	RawEndUserDescription string     `json:"raw_end_user_description,omitempty"`
	RawEndUserLabel       string     `json:"raw_end_user_label,omitempty"`
	StatusCategory        string     `json:"status_category,omitempty"`
	UpdatedAt             *time.Time `json:"updated_at,omitempty"`
}

type CustomTicketStatusCreateOption struct {
	// If true, show only active custom ticket statuses. If false, show only inactive custom ticket statuses. If the filter is not used, show all custom ticket statuses
	Active bool `json:"active,omitempty"`

	// The dynamic content placeholder or the label displayed to agents. Maximum length for displayed label is 48 characters
	AgentLabel string `json:"agent_label,omitempty"`

	// The description of when the user should select this custom ticket status
	Description string `json:"description,omitempty"`

	// The description displayed to end users
	EndUserDescription string `json:"end_user_description,omitempty"`

	// The dynamic content placeholder or the label displayed to end users. Maximum length for displayed label is 48 characters
	EndUserLabel string `json:"end_user_label,omitempty"`

	// The status category the custom status belongs to. Allowed values are "new", "open", "pending", "hold", or "solved"
	StatusCategory string `json:"status_category,omitempty"`
}

type CustomTicketStatusUpdateOption struct {
	// If true, show only active custom ticket statuses. If false, show only inactive custom ticket statuses. If the filter is not used, show all custom ticket statuses
	Active bool `json:"active,omitempty"`

	// The dynamic content placeholder or the label displayed to agents. Maximum length for displayed label is 48 characters
	AgentLabel string `json:"agent_label,omitempty"`

	// The description of when the user should select this custom ticket status
	Description string `json:"description,omitempty"`

	// The description displayed to end users
	EndUserDescription string `json:"end_user_description,omitempty"`

	// The dynamic content placeholder or the label displayed to end users. Maximum length for displayed label is 48 characters
	EndUserLabel string `json:"end_user_label,omitempty"`
}

type CustomTicketStatusListOptions struct {
	// If true, show only active custom ticket statuses. If false, show only inactive custom ticket statuses. If the filter is not used, show all custom ticket statuses
	Active bool

	// If true, show only default custom ticket statuses. If false, show only non-default custom ticket statuses. If the filter is not used, show all custom ticket statuses
	Default bool

	// Filter the list of custom ticket statuses by a comma-separated list of status categories
	StatusCategories string
}

type CustomTicketStatusListResult struct {
	CustomTicketStatuses []CustomTicketStatus `json:"custom_statuses"`
}

// CustomTicketStatusAPI an interface containing all custom ticket status related methods
type CustomTicketStatusAPI interface {
	GetCustomTicketStatuses(ctx context.Context, opts *CustomTicketStatusListOptions) ([]CustomTicketStatus, error)
	GetCustomTicketStatus(ctx context.Context, customTicketStatusID int64) (CustomTicketStatus, error)
	CreateCustomTicketStatus(ctx context.Context, customTicketStatusCreateOptions CustomTicketStatusCreateOption) (CustomTicketStatus, error)
	UpdateCustomTicketStatus(ctx context.Context, customTicketStatusID int64, customTicketStatus CustomTicketStatusUpdateOption) (CustomTicketStatus, error)
}

// GetCustomTicketStatuses: Lists all undeleted custom ticket statuses for the account. No pagination is provided.
//
// ref: https://developer.zendesk.com/api-reference/ticketing/tickets/custom_ticket_statuses/#list-custom-ticket-statuses
func (z *Client) GetCustomTicketStatuses(ctx context.Context, opts *CustomTicketStatusListOptions) ([]CustomTicketStatus, error) {
	var data struct {
		CustomTicketStatuses []CustomTicketStatus `json:"custom_statuses"`
	}

	tmp := opts
	if tmp == nil {
		tmp = &CustomTicketStatusListOptions{}
	}

	u, err := addOptions("/custom_statuses.json", tmp)
	if err != nil {
		return nil, err
	}

	body, err := z.get(ctx, u)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data.CustomTicketStatuses, nil
}

// GetCustomTicketStatus: Returns the custom ticket status object.
//
// ref:https://developer.zendesk.com/api-reference/ticketing/tickets/custom_ticket_statuses/#show-custom-ticket-status
func (z *Client) GetCustomTicketStatus(ctx context.Context, customTicketStatusID int64) (CustomTicketStatus, error) {
	var result struct {
		CustomTicketStatus CustomTicketStatus `json:"custom_status"`
	}

	body, err := z.get(ctx, fmt.Sprintf("/custom_statuses/%d.json", customTicketStatusID))
	if err != nil {
		return CustomTicketStatus{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return CustomTicketStatus{}, err
	}

	return result.CustomTicketStatus, err
}

// CreateCustomTicketStatus: Takes a CustomTicketStatusCreateOption object that specifies the custom ticket status properties to create.
//
// ref:https://developer.zendesk.com/api-reference/ticketing/tickets/custom_ticket_statuses/#create-custom-ticket-status
func (z *Client) CreateCustomTicketStatus(ctx context.Context, customTicketStatus CustomTicketStatusCreateOption) (CustomTicketStatus, error) {
	var data struct {
		CustomTicketStatus CustomTicketStatusCreateOption `json:"custom_status"`
	}

	var result struct {
		CustomTicketStatus CustomTicketStatus `json:"custom_status"`
	}

	data.CustomTicketStatus = customTicketStatus

	body, err := z.Post(ctx, "/custom_statuses.json", data)
	if err != nil {
		return CustomTicketStatus{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return CustomTicketStatus{}, err
	}

	return result.CustomTicketStatus, err
}

// UpdateCustomTicketStatus: Takes a CustomTicketStatusUpdateOption object that specifies the custom ticket status properties to update.
//
// ref:https://developer.zendesk.com/api-reference/ticketing/tickets/custom_ticket_statuses/#update-custom-ticket-status
func (z *Client) UpdateCustomTicketStatus(ctx context.Context, customTicketStatusID int64, customTicketStatus CustomTicketStatusUpdateOption) (CustomTicketStatus, error) {
	var data struct {
		CustomTicketStatus CustomTicketStatusUpdateOption `json:"custom_status"`
	}
	var result struct {
		CustomTicketStatus CustomTicketStatus `json:"custom_status"`
	}

	data.CustomTicketStatus = customTicketStatus

	path := fmt.Sprintf("/custom_statuses/%d.json", customTicketStatusID)

	body, err := z.put(ctx, path, data)
	if err != nil {
		return CustomTicketStatus{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return CustomTicketStatus{}, err
	}

	return result.CustomTicketStatus, err
}

// Custom statuses can't be deleted at this time.
// If you deactivate a custom status, it won't be available in the status picker and agents won't be able to use it.
// You can also edit a custom status and change its name, description, etc.
// ref: https://support.zendesk.com/hc/en-us/articles/4412575941402-Managing-ticket-statuses
