package zendesk

import (
	"context"
	"encoding/json"
	"time"
)

type (
	// OrganizationMembership is struct for organization membership payload
	// https://developer.zendesk.com/api-reference/ticketing/organizations/organization_memberships/
	OrganizationMembership struct {
		ID             int64     `json:"id,omitempty"`
		URL            string    `json:"url,omitempty"`
		UserID         int64     `json:"user_id"`
		OrganizationID int64     `json:"organization_id"`
		Default        bool      `json:"default"`
		Name           string    `json:"organization_name"`
		CreatedAt      time.Time `json:"created_at,omitempty"`
		UpdatedAt      time.Time `json:"updated_at,omitempty"`
	}

	// OrganizationMembershipListOptions is a struct for options for organization membership list
	// ref: https://developer.zendesk.com/api-reference/ticketing/organizations/organization_memberships/
	OrganizationMembershipListOptions struct {
		PageOptions
		OrganizationID int64 `json:"organization_id,omitempty" url:"organization_id,omitempty"`
		UserID         int64 `json:"user_id,omitempty" url:"user_id,omitempty"`
	}

	// OrganizationMembershipOptions is a struct for options for organization membership
	// https://developer.zendesk.com/api-reference/ticketing/organizations/organization_memberships/
	OrganizationMembershipOptions struct {
		OrganizationID int64 `json:"organization_id,omitempty"`
		UserID         int64 `json:"user_id,omitempty"`
	}

	// OrganizationMembershipAPI is an interface containing organization membership related methods
	OrganizationMembershipAPI interface {
		GetOrganizationMemberships(context.Context, *OrganizationMembershipListOptions) ([]OrganizationMembership, Page, error)
		CreateOrganizationMembership(context.Context, OrganizationMembershipOptions) (OrganizationMembership, error)
	}
)

// GetOrganizationMemberships gets the memberships of the specified organization
// ref: https://developer.zendesk.com/api-reference/ticketing/organizations/organization_memberships/
func (z *Client) GetOrganizationMemberships(ctx context.Context, opts *OrganizationMembershipListOptions) ([]OrganizationMembership, Page, error) {
	var result struct {
		OrganizationMemberships []OrganizationMembership `json:"organization_memberships"`
		Page
	}

	tmp := opts
	if tmp == nil {
		tmp = new(OrganizationMembershipListOptions)
	}

	u, err := addOptions("/organization_memberships.json", tmp)
	if err != nil {
		return nil, Page{}, err
	}

	body, err := z.get(ctx, u)
	if err != nil {
		return nil, Page{}, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, Page{}, err
	}

	return result.OrganizationMemberships, result.Page, nil
}

// CreateOrganizationMembership creates an organization membership for an existing user and org
// https://developer.zendesk.com/api-reference/ticketing/organizations/organization_memberships/#create-membership
func (z *Client) CreateOrganizationMembership(ctx context.Context, opts OrganizationMembershipOptions) (OrganizationMembership, error) {
	var data, result struct {
		OrganizationMembership OrganizationMembership `json:"organization_membership"`
	}

	data.OrganizationMembership = OrganizationMembership{
		UserID:         opts.UserID,
		OrganizationID: opts.OrganizationID,
	}

	body, err := z.post(ctx, "/organization_memberships.json", data)

	if err != nil {
		return OrganizationMembership{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return OrganizationMembership{}, err
	}

	return result.OrganizationMembership, err
}
