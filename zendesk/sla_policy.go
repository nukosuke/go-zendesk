package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// SlaPolicyCondition zendesk slaPolicy condition
//
// ref: https://developer.zendesk.com/rest_api/docs/core/slas/policies#conditions-reference
type SlaPolicyFilter struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// SlaPolicy is zendesk slaPolicy JSON payload format
//
// ref: https://developer.zendesk.com/rest_api/docs/core/slas/policies#json-format
type SlaPolicy struct {
	ID          int64  `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Position    int64  `json:"position,omitempty"`
	Active      bool   `json:"active,omitempty"`
	Filter      struct {
		All []SlaPolicyFilter `json:"all"`
		Any []SlaPolicyFilter `json:"any"`
	} `json:"filter"`
	PolicyMetric struct {
		Priority      string `json:"priority"`
		Metric        string `json:"metric"`
		Target        int    `json:"target"`
		BusinessHours bool   `json:"business_hours"`
	} `json:"policy_metric"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// SlaPolicyListOptions is options for GetSlaPolicies
//
// ref: https://developer.zendesk.com/rest_api/docs/support/slas/policies#list-slas/policies
type SlaPolicyListOptions struct {
	PageOptions
	Active    bool   `url:"active,omitempty"`
	SortBy    string `url:"sort_by,omitempty"`
	SortOrder string `url:"sort_order,omitempty"`
}

// SlaPolicyAPI an interface containing all slaPolicy related methods
type SlaPolicyAPI interface {
	GetSlaPolicies(ctx context.Context, opts *SlaPolicyListOptions) ([]SlaPolicy, Page, error)
	CreateSlaPolicy(ctx context.Context, slaPolicy SlaPolicy) (SlaPolicy, error)
	GetSlaPolicy(ctx context.Context, id int64) (SlaPolicy, error)
	UpdateSlaPolicy(ctx context.Context, id int64, slaPolicy SlaPolicy) (SlaPolicy, error)
	DeleteSlaPolicy(ctx context.Context, id int64) error
}

// GetSlaPolicies fetch slaPolicy list
//
// ref: https://developer.zendesk.com/rest_api/docs/support/slas/policies#getting-slas/policies
func (z *Client) GetSlaPolicies(ctx context.Context, opts *SlaPolicyListOptions) ([]SlaPolicy, Page, error) {
	var data struct {
		SlaPolicies []SlaPolicy `json:"sla_policies"`
		Page
	}

	if opts == nil {
		return []SlaPolicy{}, Page{}, &OptionsError{opts}
	}

	u, err := addOptions("/slas/policies.json", opts)
	if err != nil {
		return []SlaPolicy{}, Page{}, err
	}

	body, err := z.get(ctx, u)
	if err != nil {
		return []SlaPolicy{}, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return []SlaPolicy{}, Page{}, err
	}
	return data.SlaPolicies, data.Page, nil
}

// CreateSlaPolicy creates new slaPolicy
//
// ref: https://developer.zendesk.com/rest_api/docs/support/slas/policies#create-slaPolicy
func (z *Client) CreateSlaPolicy(ctx context.Context, slaPolicy SlaPolicy) (SlaPolicy, error) {
	var data, result struct {
		SlaPolicy SlaPolicy `json:"sla_policy"`
	}
	data.SlaPolicy = slaPolicy

	body, err := z.post(ctx, "/slas/policies.json", data)
	if err != nil {
		return SlaPolicy{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return SlaPolicy{}, err
	}
	return result.SlaPolicy, nil
}

// GetSlaPolicy returns the specified slaPolicy
//
// ref: https://developer.zendesk.com/rest_api/docs/support/slas/policies#getting-slas/policies
func (z *Client) GetSlaPolicy(ctx context.Context, id int64) (SlaPolicy, error) {
	var result struct {
		SlaPolicy SlaPolicy `json:"sla_policy"`
	}

	body, err := z.get(ctx, fmt.Sprintf("/slas/policies/%d.json", id))
	if err != nil {
		return SlaPolicy{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return SlaPolicy{}, err
	}
	return result.SlaPolicy, nil
}

// UpdateSlaPolicy updates the specified slaPolicy and returns the updated one
//
// ref: https://developer.zendesk.com/rest_api/docs/support/slas/policies#update-slaPolicy
func (z *Client) UpdateSlaPolicy(ctx context.Context, id int64, slaPolicy SlaPolicy) (SlaPolicy, error) {
	var data, result struct {
		SlaPolicy SlaPolicy `json:"sla_policy"`
	}

	data.SlaPolicy = slaPolicy
	body, err := z.put(ctx, fmt.Sprintf("/slas/policies/%d.json", id), data)
	if err != nil {
		return SlaPolicy{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return SlaPolicy{}, err
	}

	return result.SlaPolicy, nil
}

// DeleteSlaPolicy deletes the specified slaPolicy
//
// ref: https://developer.zendesk.com/rest_api/docs/support/slas/policies#delete-slaPolicy
func (z *Client) DeleteSlaPolicy(ctx context.Context, id int64) error {
	err := z.delete(ctx, fmt.Sprintf("/slas/policies/%d.json", id))
	if err != nil {
		return err
	}

	return nil
}
