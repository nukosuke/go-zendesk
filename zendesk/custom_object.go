package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type CustomObjectRecord struct {
	Url                string                 `json:"url,omitempty"`
	Name               string                 `json:"name,omitempty"`
	ID                 string                 `json:"id,omitempty"`
	CustomObjectKey    string                 `json:"custom_object_key"`
	CustomObjectFields map[string]interface{} `json:"custom_object_fields" binding:"required"`
	CreatedByUserID    string                 `json:"created_by_user_id,omitempty"`
	UpdatedByUserID    string                 `json:"updated_by_user_id,omitempty"`
	CreatedAt          time.Time              `json:"created_at,omitempty"`
	UpdatedAt          time.Time              `json:"updated_at,omitempty"`
	ExternalID         string                 `json:"external_id,omitempty"`
}

// CustomObjectAPI an interface containing all custom object related methods
type CustomObjectAPI interface {
	CreateCustomObjectRecord(
		ctx context.Context, record CustomObjectRecord, customObjectKey string) (CustomObjectRecord, error)
	AutocompleteSearchCustomObjectRecords(
		ctx context.Context,
		customObjectKey string,
		opts *CustomObjectAutocompleteOptions,
	) ([]CustomObjectRecord, Page, error)
	SearchCustomObjectRecords(
		ctx context.Context, customObjectKey string, opts *SearchCustomObjectRecordsOptions,
	) ([]CustomObjectRecord, Page, error)
	ListCustomObjectRecords(
		ctx context.Context, customObjectKey string, opts *CustomObjectListOptions) ([]CustomObjectRecord, Page, error)
	ShowCustomObjectRecord(
		ctx context.Context, customObjectKey string, customObjectRecordID string,
	) (*CustomObjectRecord, error)
	UpdateCustomObjectRecord(
		ctx context.Context, customObjectKey string, customObjectRecordID string, record CustomObjectRecord,
	) (*CustomObjectRecord, error)
}

// CustomObjectAutocompleteOptions custom object search options
type CustomObjectAutocompleteOptions struct {
	PageOptions
	Name string `url:"name"`
}

// CreateCustomObjectRecord CreateCustomObject create a custom object record
func (z *Client) CreateCustomObjectRecord(
	ctx context.Context, record CustomObjectRecord, customObjectKey string,
) (CustomObjectRecord, error) {

	var data, result struct {
		CustomObjectRecord CustomObjectRecord `json:"custom_object_record"`
	}
	data.CustomObjectRecord = record

	body, err := z.post(ctx, fmt.Sprintf("/custom_objects/%s/records.json", customObjectKey), data)
	if err != nil {
		return CustomObjectRecord{}, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return CustomObjectRecord{}, err
	}
	return result.CustomObjectRecord, nil
}

// CustomObjectListOptions custom object list options
type CustomObjectListOptions struct {
	PageOptions
	ExternalIds []string `url:"external_ids"`
}

// ListCustomObjectRecords list objects
// https://developer.zendesk.com/api-reference/custom-objects/custom_object_records/#list-custom-object-records
func (z *Client) ListCustomObjectRecords(
	ctx context.Context, customObjectKey string, opts *CustomObjectListOptions) ([]CustomObjectRecord, Page, error) {
	var result struct {
		CustomObjectRecords []CustomObjectRecord `json:"custom_object_records"`
		Page
	}
	tmp := opts
	if tmp == nil {
		tmp = &CustomObjectListOptions{}
	}
	url := fmt.Sprintf("/custom_objects/%s/records", customObjectKey)
	urlWithOptions, err := addOptions(url, tmp)
	body, err := z.get(ctx, urlWithOptions)

	if err != nil {
		return nil, Page{}, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, Page{}, err
	}
	return result.CustomObjectRecords, result.Page, nil
}

// AutocompleteSearchCustomObjectRecords search for a custom object record by the name field
// https://developer.zendesk.com/api-reference/custom-objects/custom_object_records/#autocomplete-custom-object-record-search
func (z *Client) AutocompleteSearchCustomObjectRecords(
	ctx context.Context, customObjectKey string, opts *CustomObjectAutocompleteOptions,
) ([]CustomObjectRecord, Page, error) {
	var result struct {
		CustomObjectRecords []CustomObjectRecord `json:"custom_object_records"`
		Page
	}
	tmp := opts
	if tmp == nil {
		tmp = &CustomObjectAutocompleteOptions{}
	}
	url := fmt.Sprintf("/custom_objects/%s/records/autocomplete", customObjectKey)
	urlWithOptions, err := addOptions(url, tmp)
	body, err := z.get(ctx, urlWithOptions)

	if err != nil {
		return nil, Page{}, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, Page{}, err
	}
	return result.CustomObjectRecords, result.Page, nil
}

type SearchCustomObjectRecordsOptions struct {
	PageOptions

	// One of name, created_at, updated_at, -name, -created_at, or -updated_at.
	// The - denotes the sort will be descending. Defaults to sorting by relevance.
	Sort string `url:"sort,omitempty"`

	// Query string
	Query string `url:"query,omitempty"`
}

// SearchCustomObjectRecords search for a custom object record by the name field
// https://developer.zendesk.com/api-reference/custom-objects/custom_object_records/#search-custom-object-records
func (z *Client) SearchCustomObjectRecords(
	ctx context.Context, customObjectKey string, opts *SearchCustomObjectRecordsOptions,
) ([]CustomObjectRecord, Page, error) {
	var result struct {
		CustomObjectRecords []CustomObjectRecord `json:"custom_object_records"`
		Page
	}
	tmp := opts
	if tmp == nil {
		tmp = &SearchCustomObjectRecordsOptions{}
	}
	url := fmt.Sprintf("/custom_objects/%s/records/search", customObjectKey)
	urlWithOptions, err := addOptions(url, tmp)
	body, err := z.get(ctx, urlWithOptions)

	if err != nil {
		return nil, Page{}, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, Page{}, err
	}
	return result.CustomObjectRecords, result.Page, nil
}

// ShowCustomObjectRecord returns a custom record for a specific object using a provided id.
// https://developer.zendesk.com/api-reference/custom-objects/custom_object_records/#show-custom-object-record
func (z *Client) ShowCustomObjectRecord(
	ctx context.Context, customObjectKey string, customObjectRecordID string,
) (*CustomObjectRecord, error) {
	var result struct {
		CustomObjectRecord CustomObjectRecord `json:"custom_object_record"`
	}

	url := fmt.Sprintf("/custom_objects/%s/records/%s", customObjectKey, customObjectRecordID)
	body, err := z.get(ctx, url)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}
	return &result.CustomObjectRecord, nil
}

// UpdateCustomObjectRecord Updates an individual custom object record
// https://developer.zendesk.com/api-reference/custom-objects/custom_object_records/#update-custom-object-record
func (z *Client) UpdateCustomObjectRecord(
	ctx context.Context, customObjectKey string, customObjectRecordID string, record CustomObjectRecord,
) (*CustomObjectRecord, error) {
	var data, result struct {
		CustomObjectRecord CustomObjectRecord `json:"custom_object_record"`
	}
	data.CustomObjectRecord = record

	url := fmt.Sprintf("/custom_objects/%s/records/%s", customObjectKey, customObjectRecordID)
	body, err := z.patch(ctx, url, data)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}
	return &result.CustomObjectRecord, nil
}
