package zendesk

import (
	"context"
	"encoding/json"
	"time"
)

// OrganizationField represents the Organization Custom field structure
type OrganizationField struct {
	ID                     int64               `json:"id,omitempty"`
	URL                    string              `json:"url,omitempty"`
	Title                  string              `json:"title"`
	Type                   string              `json:"type"`
	RelationshipTargetType string              `json:"relationship_target_type"`
	RelationshipFilter     RelationshipFilter  `json:"relationship_filter"`
	Active                 bool                `json:"active,omitempty"`
	CustomFieldOptions     []CustomFieldOption `json:"custom_field_options,omitempty"`
	Description            string              `json:"description,omitempty"`
	Key                    string              `json:"key"`
	Position               int64               `json:"position,omitempty"`
	RawDescription         string              `json:"raw_description,omitempty"`
	RawTitle               string              `json:"raw_title,omitempty"`
	RegexpForValidation    string              `json:"regexp_for_validation,omitempty"`
	System                 bool                `json:"system,omitempty"`
	Tag                    string              `json:"tag,omitempty"`
	CreatedAt              *time.Time          `json:"created_at,omitempty"`
	UpdatedAt              *time.Time          `json:"updated_at,omitempty"`
}

// OrganizationFieldAPI an interface containing all the organization field related zendesk methods
type OrganizationFieldAPI interface {
	GetOrganizationFields(ctx context.Context) ([]OrganizationField, Page, error)
	CreateOrganizationField(ctx context.Context, organizationField OrganizationField) (OrganizationField, error)
}

// GetOrganizationFields fetches organization field list
// ref: https://developer.zendesk.com/api-reference/ticketing/organizations/organization_fields/#list-organization-fields
func (z *Client) GetOrganizationFields(ctx context.Context) ([]OrganizationField, Page, error) {
	var data struct {
		OrganizationFields []OrganizationField `json:"organization_fields"`
		Page
	}

	body, err := z.get(ctx, "/organization_fields.json")
	if err != nil {
		return []OrganizationField{}, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return []OrganizationField{}, Page{}, err
	}
	return data.OrganizationFields, data.Page, nil
}

// CreateOrganizationField creates new organization field
// ref: https://developer.zendesk.com/api-reference/ticketing/organizations/organization_fields/#create-organization-field
func (z *Client) CreateOrganizationField(ctx context.Context, organizationField OrganizationField) (OrganizationField, error) {
	var data, result struct {
		OrganizationField OrganizationField `json:"organization_field"`
	}
	data.OrganizationField = organizationField

	body, err := z.post(ctx, "/organization_fields.json", data)
	if err != nil {
		return OrganizationField{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return OrganizationField{}, err
	}
	return result.OrganizationField, nil
}
